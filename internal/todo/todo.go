package todo

import "fmt"

func ListTodos(todos []string) {
	if len(todos) == 0 {
		fmt.Println("No tasks in list")
		return
	}

	fmt.Println("Your tasks:")
	for i, task := range todos {
		fmt.Printf("%d: %s\n", i+1, task)
	}
}