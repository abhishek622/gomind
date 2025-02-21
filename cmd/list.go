package cmd

import (
	"abhishek622/gomind/helper/todo"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long:  "List all todos",
	Run: func(cmd *cobra.Command, args []string) {

		// Load existing todos
		todos, err := todo.ReadTodos()
		if err != nil {
			fmt.Println("Error reading todos:", err)
			return
		}

		printTodosTable(todos)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func printTodosTable(todos []todo.Todo) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Print the table header
	fmt.Fprintln(w, "ID\tTask\tCompleted\tCreated At")

	// Print each todo
	for _, t := range todos {
		completedStatus := "No"
		if t.Completed {
			completedStatus = "Yes"
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", t.ID, t.Task, completedStatus, t.CreatedAt)
	}

	// Flush the tabwriter to display the table
	w.Flush()
}
