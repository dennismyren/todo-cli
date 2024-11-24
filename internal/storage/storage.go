package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

const todoFile = "todos.json"

var todoList = []string{}

func GetTodoList() []string {
	return todoList
}

func LoadTodos() {
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

func SaveTodos() {
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

func AddTodo(task string) {
	todoList = append(todoList, task)
	fmt.Printf("Added task: %s\n", task)
	SaveTodos()
}

func RemoveTodo(index int) {
	if index < 1 || index > len(todoList) {
		fmt.Println("Invalid todo index")
		return
	}

	task := todoList[index-1]
	todoList = append(todoList[:index-1], todoList[index:]...)
	fmt.Printf("Removed todo: %s\n", task)
	SaveTodos()
}