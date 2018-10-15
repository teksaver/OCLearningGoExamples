package main

import "fmt"

func displayListOfString(theList []string) {
	fmt.Println("Here is the list")
	// Display one task after the other
	for index, task := range theList {
		fmt.Println("Task N°", index, " is ", task)
	}
	// Write an empty line
	fmt.Println()
}

func removeFromSliceByIndex(slice []string, s int) []string {
	// Append the strings before and after the element to be removed
	return append(slice[:s], slice[s+1:]...)
}

// The main function is where everything begins
func main() {
	// Create an empty list of strings to represent tasks
	var myTodoList []string
	// Add 4 tasks to the list
	myTodoList = append(myTodoList, "Wake up", "Shower", "Have breakfast", "Go to work")
	// Display tasks in the list
	displayListOfString(myTodoList)
	// Remove breakfast
	myTodoList = removeFromSliceByIndex(myTodoList, 2)
	// Display task list after removal
	displayListOfString(myTodoList)
	// change shower to bath
	myTodoList[1] = "Take Bath"

}
