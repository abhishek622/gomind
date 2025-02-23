package cmd

import (
	"abhishek622/gomind/helper/task"
	"fmt"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:    "add",
	Short:  "Add a new task",
	Args:   cobra.ExactArgs(1), // Ensure exactly one argument (the task) is provided
	PreRun: preRunCheck,        // Ensure config & DB check only when needed
	Run: func(cmd *cobra.Command, args []string) {
		// get task from the cmd args
		description := args[0]

		priorityStr, _ := cmd.Flags().GetString("priority")
		priority := task.Priority(priorityStr)

		newTask := task.Task{
			Description: description,
			Priority:    priority,
			Completed:   false,
		}

		repo := task.NewRepository()
		err := repo.CreateTask(&newTask)
		if err != nil {
			fmt.Println("Failed to add task:", err)
			return
		}

		fmt.Printf("Task added successfully!")
	},
}

func init() {
	AddCmd.Flags().StringP("priority", "p", "Medium", "Set task priority: High, Medium, or Low")
}
