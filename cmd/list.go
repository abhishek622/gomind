package cmd

import (
	"abhishek622/gomind/helper/task"
	"abhishek622/gomind/utils"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

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

		// Create a tab writer for formatted output
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)
		// Write the header
		fmt.Fprintln(w, "ID\tTask\tPriority\tCompleted\tCreated At")

		// Write each task
		for _, task := range tasks {
			createdAtFormatted := utils.FormatTime(task.CreatedAt.Format(time.RFC3339))
			fmt.Fprintf(w, "%d\t%s\t%s\t%t\t%s\n", task.ID, task.Description, task.Priority, task.Completed, createdAtFormatted)
		}

		w.Flush()
	},
}
