package stctl

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

const (
	api_url = "/api/v1/clusters"
)

type listParams struct {
}

type createParams struct {
	Config []byte `json:"config"`
}

type clusterId struct {
	Id string `json:"id"`
}

type bulkScanParams struct {
	List []clusterId `json:"list"`
}

type createResponse struct {
	Id string `json:"id"`
}

type clusterList struct {
	List []cluster `json:"list"`
}

type cluster struct {
	Id   string `json:"id"`
	Meta struct {
		Name string `json:"name"`
	} `json:"meta"`

	LastScan struct {
		Timestamp string `json:"timestamp"`
		Grade     string `json:"grade"`
		Version   string `json:"version"`
	} `json:"last_scan"`
}

func (ctx *Ctx) ClusterList() ([]string, error) {

	var clusters clusterList
	err := ctx.api_call(api_url, "list", listParams{}, &clusters)
	if err != nil {
		fmt.Printf("Can't get cluster list: %v\n", err)
		return nil, err
	}

	for _, cluster := range clusters.List {

		timestamp, err := strconv.ParseInt(cluster.LastScan.Timestamp, 0, 64)
		if err != nil {
			fmt.Printf("Cluster: %s - %s - never scanned\n", cluster.Meta.Name, cluster.LastScan.Version)
			continue
		}

		scanTime := time.Unix(timestamp, 0).UTC()

		if timestamp == 0 {
			fmt.Printf("Cluster: %s - %s - never scanned\n", cluster.Meta.Name, cluster.LastScan.Version)
			continue
		}

		fmt.Printf("Cluster: %s - (Version: %s grade: %s - %s)\n", cluster.Meta.Name, cluster.LastScan.Version, cluster.LastScan.Grade, scanTime)
	}

	return nil, nil
}

func (ctx *Ctx) ClusterCreate(filename string) (string, error) {

	kubeConfig, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("cannot load kubeconfig: %v", err)
	}

	postData := createParams{
		Config: kubeConfig,
	}

	var response createResponse
	err = ctx.api_call(api_url, "create", postData, &response)
	if err != nil {
		return "", err
	}

	return response.Id, nil
}

func (ctx *Ctx) ClusterBulkScan() error {

	scanParams := &bulkScanParams{}

	var clusters clusterList
	err := ctx.api_call(api_url, "list", listParams{}, &clusters)
	if err != nil {
		return fmt.Errorf("cannot retrieve cluster list: %v", err)
	}

	for _, toScan := range clusters.List {
		scanParams.List = append(scanParams.List, clusterId{Id: toScan.Id})
	}

	var response interface{}
	err = ctx.api_call(api_url, "bulkscan", scanParams, &response)
	if err != nil {
		return fmt.Errorf("cannot bulkscan clusters: %v", err)
	}

	return nil
}
