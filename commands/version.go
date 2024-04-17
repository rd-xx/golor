package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of golor",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: 0.1.0")
		fmt.Println("Copyright: (C) 2024 rdx")
		fmt.Println("License: MIT")
	},
}
