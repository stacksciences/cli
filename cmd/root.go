package cmd

import (
	"fmt"
	stctl "stctl/pkg"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Setup struct {
	PlatformHostname string `toml:"platform"`
	PlatformToken    string `toml:token`
}

var (
	stctlCtx  *stctl.Ctx
	fileSetup Setup

	cfgFile    string
	hostname   string
	token      string
	kubeConfig string

	rootCmd = &cobra.Command{
		Use:   "stctl",
		Short: "stctl is a cli tool to interact with the stacksciences platform",
		Long:  `stctl is a cli tool to interact with the stacksciences platform. You can interact with your cluster, exposure and container analysis`,
		Args:  cobra.MinimumNArgs(1),
		Run:   clusterExec,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {

	stctlCtx = &stctl.Ctx{}

	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(clusterCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.stctl/config)")
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "Platform authentication token")
	rootCmd.PersistentFlags().StringVar(&hostname, "platform", "", "Platform hostname")

	clusterCreateCmd.PersistentFlags().StringVar(&kubeConfig, "kubeconfig", "", "Platform hostname")

	err := viper.BindPFlag("platform", rootCmd.PersistentFlags().Lookup("platform"))
	if err != nil {
		fmt.Printf("cannot bind platform flag: %v\n", err)
	}
	err = viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	if err != nil {
		fmt.Printf("cannot bind token flag: %v\n", err)
	}

	viper.SetDefault("platform", "app.stacksciences.com")

	clusterCmd.AddCommand(clusterListCmd)
	clusterCmd.AddCommand(clusterCreateCmd)
	clusterCmd.AddCommand(clusterBulkScanCmd)
}

func initConfig() {

	if len(cfgFile) > 0 {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName(".stctl")
	}

	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	stctlCtx.PlatformHostname = viper.GetString("platform")
	stctlCtx.PlatformToken = viper.GetString("token")

}
