package cmd

import (
	"fmt"
	"os"

	"github.com/marcosgeo/go/task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all fo your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete! Why not take a vacation? :beach:")
			return
		}
		fmt.Println("You have the folowing tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s, Key=%d\n", i+1, task.Value, task.Key)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
