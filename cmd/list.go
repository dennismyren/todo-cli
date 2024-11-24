package cmd

import (
	"todo-cli/internal/storage"
	"todo-cli/internal/todo"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		todos := storage.GetTodoList()
		todo.ListTodos(todos)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}