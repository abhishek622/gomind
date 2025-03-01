package cmd

import (
	"abhishek622/gomind/helper/task"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var CompleteCmd = &cobra.Command{
	Use:    "complete",
	Short:  "Mark task as completed",
	Args:   cobra.ExactArgs(1),
	PreRun: preRunCheck,
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid task ID:", err)
			return
		}

		repo := task.NewRepository()
		err = repo.MarkAsCompleted(taskId)
		if err != nil {
			fmt.Println("Failed to mark task as completed", err)
			return
		}

		fmt.Println("✅ Task marked as completed")

	},
}
