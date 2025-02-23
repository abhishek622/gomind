package cmd

import (
	"abhishek622/gomind/helper/task"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// Flags
var category string
var priority string
var dueDate string
var dependencies string

var AddCmd = &cobra.Command{
	Use:    "add [-c category] [-p priority] [-d YYYY-MM-DD] [-x dependencies]",
	Short:  "Add a new task",
	Args:   cobra.MinimumNArgs(1),
	PreRun: preRunCheck, // Ensure config & DB check only when needed
	Run: func(cmd *cobra.Command, args []string) {
		// get task from the cmd args
		description := args[0]

		if priority == "" {
			priority = string(task.Medium)
		} else {
			switch priority {
			case string(task.High), string(task.Medium), string(task.Low):
				break
			default:
				fmt.Println("Error: Invalid priority.  Choose from High, Medium, or Low.")
				return
			}
		}

		/*
			In Go, the layout string uses a specific reference time: Mon Jan 2 15:04:05 MST 2006.
			Here, "2006-01-02" means the expected format is YYYY-MM-DD (e.g., 2023-10-15).
		*/
		var due time.Time
		if dueDate != "" {
			var err error
			due, err = time.Parse("2006-01-02", dueDate)
			if err != nil {
				fmt.Println("Error: Invalid due date format. Use YYYY-MM-DD.")
				return
			}
		}

		repo := task.NewRepository()

		var deps []int64
		if dependencies != "" {
			depStrs := strings.Split(dependencies, ",")
			for _, dep := range depStrs {
				depID, err := strconv.ParseInt(strings.TrimSpace(dep), 10, 64)
				if err != nil {
					fmt.Printf("Error: Invalid dependency ID '%s'. Must be an integer.\n", dep)
					return
				}

				deps = append(deps, depID)
			}

			// check if parent tasks are present or not
			if !repo.CheckDependenciesExist(deps) {
				fmt.Println("Error: One or more dependencies do not exist in the database.")
				return
			}
		}

		newTask := task.Task{
			Description:  description,
			Category:     category,
			Priority:     task.Priority(priority),
			DueDate:      due,
			Dependencies: deps,
			Completed:    false,
		}

		err := repo.CreateTask(&newTask)
		if err != nil {
			fmt.Println("Failed to add task:", err)
			return
		}

		fmt.Printf("Task added successfully!")
	},
}

func init() {
	AddCmd.Flags().StringVarP(&category, "category", "c", "", "Task category")
	AddCmd.Flags().StringVarP(&priority, "priority", "p", "Medium", "Task priority (High, Medium, Low)")
	AddCmd.Flags().StringVarP(&dueDate, "due", "d", "", "Due date (YYYY-MM-DD)")
	AddCmd.Flags().StringVarP(&dependencies, "dependencies", "x", "", "Comma-separated list of dependency task IDs")
}
