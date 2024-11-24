package main

import (
	"todo-cli/cmd"
	"todo-cli/internal/storage"
)

func main() {
	storage.LoadTodos()
	cmd.Execute()
}
