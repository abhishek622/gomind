package task

import (
	"abhishek622/gomind/utils"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func DisplayTasks(tasks []Task) {
	// NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(writer, "ID\tTask\tCategory\tPriority\tDue Date\tCompleted\tCreated At")

	for _, task := range tasks {
		var formattedDueDate string
		formattedCompleted := "❌"
		if !task.DueDate.IsZero() { // need to handle empty string, null or undefined
			formattedDueDate = utils.FormatTime(task.DueDate.Format(time.RFC3339))
		} else {
			formattedDueDate = "N/A" // Handle empty due date
		}

		if task.Completed {
			formattedCompleted = "✅"
		}

		fmt.Fprint(writer,
			task.ID, "\t",
			task.Description, "\t",
			task.Category, "\t",
			task.Priority, "\t",
			formattedDueDate, "\t",
			formattedCompleted, "\t",
			utils.FormatTime(task.CreatedAt.Format(time.RFC3339)),
			"\t\n",
		)
	}

	writer.Flush()
}
