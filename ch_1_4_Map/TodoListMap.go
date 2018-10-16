package main

import "fmt"

// TodoList  models  common operations on a Todo List
type TodoList struct {
	tasks map[string]bool
}

// NewTodoList creates the TodoList by allocating the memory
func NewTodoList() TodoList {
	return TodoList{
		make(map[string]bool),
	}
}

// Display prints all the tasks of the list
func (t TodoList) Display() {
	fmt.Println("Here is the list")
	for taskName, done := range t.tasks {
		if done {
			fmt.Println("Task ", taskName, " is completed")
		} else {
			fmt.Println("Task ", taskName, " is not completed yet")
		}
	}
}

// Add a string at the end of the slice
func (t *TodoList) Add(taskName string) {
	t.tasks[taskName] = false
}

// MarkAsDone marks a task as done
func (t *TodoList) MarkAsDone(taskName string) {
	t.tasks[taskName] = true
}

// Remove a task
func (t *TodoList) Remove(taskName string) {
	delete(t.tasks, taskName)
}

// Rename a task
func (t *TodoList) Rename(taskName string, newName string) {
	t.tasks[newName] = t.tasks[taskName]
	delete(t.tasks, taskName)
}
