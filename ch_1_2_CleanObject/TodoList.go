package main

import "fmt"

// TodoList  models  common operations on a Todo List
type TodoList struct {
	tasks []string
}

// Display prints all the tasks of the list
func (t TodoList) Display() {
	fmt.Println("Here is the list")
	for index, task := range t.tasks {
		fmt.Println("Task N°", index, " is ", task)
	}
	fmt.Println()
}

// Add a string at the end of the slice
func (t *TodoList) Add(value string) {
	t.tasks = append(t.tasks, value)
}

// Remove a string from the slice at the specified index
func (t *TodoList) Remove(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

// Rename a task
func (t *TodoList) Rename(index int, newvalue string) {
	t.tasks[index] = newvalue
}
