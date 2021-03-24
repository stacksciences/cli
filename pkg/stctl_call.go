package stctl

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (ctx *Ctx) api_call(api_url, action string, payload interface{}, response interface{}) error {

	if len(ctx.PlatformHostname) == 0 {
		return fmt.Errorf("invalid platform endpoint url")
	}

	url := fmt.Sprintf("https://%s%s/%s", ctx.PlatformHostname, api_url, action)

	postBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("cannot marshal body: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-stack-token", ctx.PlatformToken)

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("error calling platform api:%s: %s", url, resp.Status)
	}

	err = json.Unmarshal(bodyData, response)
	if err != nil {
		return fmt.Errorf("unable to unmarshal response: %v", err)
	}

	return nil
}
