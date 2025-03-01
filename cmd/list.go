package cmd

import (
	"abhishek622/gomind/helper/task"
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:    "list",
	Short:  "List all tasks",
	PreRun: preRunCheck,
	Run: func(cmd *cobra.Command, args []string) {

		repo := task.NewRepository()
		tasks, err := repo.GetTasks()
		if err != nil {
			fmt.Println("Error fetching tasks:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		// Display tasks in a table
		task.DisplayTasks(tasks)
	},
}
