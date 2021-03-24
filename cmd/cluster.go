package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	clusterCmd = &cobra.Command{
		Use:   "cluster",
		Short: "manage your clusters",
		Long:  `interact with your cluster analysis and status`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	clusterListCmd = &cobra.Command{
		Use:   "list",
		Short: "list clusters",
		Args:  cobra.MinimumNArgs(0),
		Run:   clusterListExec,
	}

	clusterBulkScanCmd = &cobra.Command{
		Use:   "bulkscan",
		Short: "scan all clusters",
		Args:  cobra.MinimumNArgs(0),
		Run:   clusterBulkScanExec,
	}

	clusterCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "create cluster",
		Args:  cobra.MinimumNArgs(0),
		Run:   clusterCreateExec,
	}
)

func clusterExec(cmd *cobra.Command, args []string) {
	fmt.Printf("Print project: %s\n", args[0])
}

func clusterListExec(cmd *cobra.Command, args []string) {
	clusters, err := stctlCtx.ClusterList()
	if err != nil {
		return
	}

	for _, cluster := range clusters {

		fmt.Printf("Cluster: %s\n", cluster)

	}
}

func clusterBulkScanExec(cmd *cobra.Command, args []string) {
	err := stctlCtx.ClusterBulkScan()
	if err != nil {
		fmt.Printf("unable to bulkscan clusters: %v\n", err)
		return
	}
}

func clusterCreateExec(cmd *cobra.Command, args []string) {

	if len(kubeConfig) > 0 {

		fmt.Printf("Create cluster from kubeconfig file: %s\n", kubeConfig)

		id, err := stctlCtx.ClusterCreate(kubeConfig)
		if err != nil {
			fmt.Printf("cannot create cluster from kubeconfig: %v\n", err)
			os.Exit(-1)
		}

		fmt.Printf("Cluster created with id: %s\n", id)
	} else {
		fmt.Printf("need a name and location")
	}
}
