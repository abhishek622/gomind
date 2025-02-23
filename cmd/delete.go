package cmd

import (
	"abhishek622/gomind/helper/task"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var DeleteTaskCmd = &cobra.Command{
	Use:    "del [task_id | all]",
	Short:  "Delete task",
	Long:   "Delete a task by ID or delete all tasks",
	Args:   cobra.MaximumNArgs(1),
	PreRun: preRunCheck,
	Run: func(cmd *cobra.Command, args []string) {
		repo := task.NewRepository()
		if args[0] == "all" {
			err := repo.DeleteAllTask()
			if err != nil {
				fmt.Println("Error deleting all task:", err)
				return
			}

			fmt.Println("All tasks deleted successfully.")
			return
		}

		taskID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		err = repo.DeleteATask(taskID)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}

		fmt.Printf("Task with ID %d deleted successfully.\n", taskID)
	},
}
