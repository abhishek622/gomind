package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A Simple CLI todo app",
	Long:  "A CLI app to manage your todos.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
