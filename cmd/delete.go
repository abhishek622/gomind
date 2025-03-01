package cmd

import (
	"abhishek622/gomind/helper/task"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var DeleteTaskCmd = &cobra.Command{
	Use:    "del [task_id | all]",
	Short:  "Delete task",
	Long:   "Delete a task by ID or delete all tasks",
	Args:   cobra.MinimumNArgs(1),
	PreRun: preRunCheck,
	Run: func(cmd *cobra.Command, args []string) {
		repo := task.NewRepository()
		if args[0] == "all" {
			err := repo.DeleteAllTask()
			if err != nil {
				fmt.Println("❌ Error deleting all task:", err)
				return
			}

			fmt.Println("✅ All tasks deleted successfully.")
			return
		}

		inputIDs := strings.Split(args[0], ",")
		var taskIDs []int64

		for _, input := range inputIDs {
			id := strings.TrimSpace(input)
			taskID, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				fmt.Println("❌ Invalid task ID:", id)
				return
			}

			taskIDs = append(taskIDs, taskID)
		}

		err := repo.DeleteTasks(taskIDs)
		if err != nil {
			fmt.Println("❌ Error deleting task:", err)
			return
		}

		fmt.Printf("✅ Successfully delete Task!")
	},
}
