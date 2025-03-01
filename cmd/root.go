package cmd

import (
	"abhishek622/gomind/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the base command for the CLI
var RootCmd = &cobra.Command{
	Use:   "gomind",
	Short: "A CLI-based todo application",
	Long:  "Go mind allows you to manage tasks with natural language processing.",
}

// preRunCheck ensures config and DB are ready before running commands
func preRunCheck(cmd *cobra.Command, args []string) {
	utils.LoadConfig()
	utils.ConnectDB()
}

// Initialize CLI
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	// Add subcommands to RootCmd
	RootCmd.AddCommand(AddCmd)
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(CompleteCmd)
	RootCmd.AddCommand(DeleteTaskCmd)
	RootCmd.AddCommand(GenCmd)
}
