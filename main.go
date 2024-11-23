package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

const todoFile = "todos.json"

var todoList = []string{}

func main() {
	loadTodos()

	var rootCmd = &cobra.Command{Use: "todo"}

	var addCmd = &cobra.Command{
		Use:   "add [task]",
		Short: "Add new task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			task := args[0]
			addTask(task)
		},
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			listTask()
		},
	}

	var removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove a task by index",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			index, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid index, must be a number")
				os.Exit(1)
			}

			removeTask(index)
			saveTodos()
		},
	}

	rootCmd.AddCommand(addCmd, listCmd, removeCmd)
	rootCmd.Execute()
}

func addTask(task string) {
	todoList = append(todoList, task)
	fmt.Printf("Added task: %s\n", task)
	saveTodos()
}

func listTask() {
	if len(todoList) == 0 {
		fmt.Println("No tasks in list")
		return
	}

	fmt.Println("Your tasks:")
	for i, task := range todoList {
		fmt.Printf("%d: %s\n", i+1, task)
	}
}

func removeTask(index int) {
	if index < 1 || index > len(todoList) {
		fmt.Println("Invalid task index")
		return
	}

	task := todoList[index-1]
	todoList = append(todoList[:index-1], todoList[index:]...)
	fmt.Printf("Removed task: %s\n", task)
	saveTodos()
}

func loadTodos() {
	file, err := os.Open(todoFile)
	if err != nil {
		if os.IsNotExist(err) {
			todoList = []string{}
			return
		}
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&todoList); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		os.Exit(1)
	}
}

func saveTodos() {
	file, err := os.Create(todoFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(todoList); err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}
