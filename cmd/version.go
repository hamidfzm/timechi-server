package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionMajor, versionMinor, versionPatch = "0", "0", "0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information and exit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version %s.%s.%s\n", versionMajor, versionMinor, versionPatch)
	},
}
