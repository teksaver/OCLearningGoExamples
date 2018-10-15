package main

import "fmt"

// TodoList  models  common operations on a Todo List
type TodoList struct {
	tasks map[string]bool
}

// Display prints all the tasks of the list
func (t TodoList) Display() {
	fmt.Println("Here is the list")
	for taskName, done := range t.tasks {
		if done {
			fmt.Println("Task ", taskName, " is done")
		} else {
			fmt.Println("Task ", taskName, " is not completed yet")
		}

	}
	fmt.Println()
}

// Add a string at the end of the slice
func (t *TodoList) Add(taskName string) {
	t.tasks[taskName] = false
}

// Remove a string from the slice at the specified index
func (t *TodoList) Remove(taskName string) {
	delete(t.tasks, taskName)
}

// Rename a task
func (t *TodoList) Rename(taskName string, newName string) {
	t.tasks[newName] = t.tasks[taskName]
	delete(t.tasks, taskName)
}
