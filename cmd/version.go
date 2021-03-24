package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version   string
	BuildDate string
	GitTag    string

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "get cli version",
		Long:  `get cli version, git tag and build date`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s Git tag: %s Build date: %s\n", Version, GitTag, BuildDate)
		},
	}
)
