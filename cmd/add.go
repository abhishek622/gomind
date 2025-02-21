package cmd

import (
	"abhishek622/gomind/helper/todo"
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add a new todo to the list.`,
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument (the task) is provided
	Run: func(cmd *cobra.Command, args []string) {
		// get task from the cmd args
		task := args[0]

		// Load existing todos
		todos, err := todo.ReadTodos()
		if err != nil {
			fmt.Println("Error reading todos:", err)
			return
		}

		// Create a new todo
		newTodo := todo.NewTodo(task)

		// Add the new todo to the list
		todos = append(todos, *newTodo)

		// Save the updated list back to db
		err = todo.WriteTodos(todos)
		if err != nil {
			fmt.Println("Error saving todos:", err)
			return
		}

		fmt.Printf("Added new todos: %s\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
