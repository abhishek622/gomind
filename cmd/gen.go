package cmd

import (
	"abhishek622/gomind/helper/gpt"
	"abhishek622/gomind/helper/task"
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var GenCmd = &cobra.Command{
	Use:    "gen",
	Short:  "Generate tasks using AI",
	Args:   cobra.MinimumNArgs(1),
	PreRun: preRunCheck,
	Run: func(cmd *cobra.Command, args []string) {
		// Join args into a single user input string
		userInput := args[0]

		fmt.Println("🔄 Using AI to generate structured tasks...")

		// Generate tasks using GPT API
		taskList, err := gpt.GenerateTask(userInput)
		if err != nil {
			fmt.Println("❌ Error generating tasks:", err)
			return
		}

		// Check if tasks were actually generated
		if len(taskList) == 0 {
			fmt.Println("⚠️ No tasks were generated. Try refining your input!")
			return
		}

		// fmt.Println(taskList)

		// Parse taskList from JSON string to []task.Task
		var tasks []task.Task
		err = json.Unmarshal([]byte(taskList), &tasks)
		if err != nil {
			fmt.Println("❌ Error parsing tasks:", err)
			return
		}

		// Set default values for each generated task
		currentTime := time.Now()
		for i := range tasks {
			tasks[i].Completed = false
			tasks[i].CreatedAt = currentTime
			tasks[i].UpdatedAt = currentTime
		}

		fmt.Println("Adding task to db...")
		repo := task.NewRepository()
		err = repo.InsertBulkTask(tasks)
		if err != nil {
			fmt.Println("Error while add task:", err)
			return
		}

		// Display tasks in a table
		task.DisplayTasks(tasks)
	},
}
