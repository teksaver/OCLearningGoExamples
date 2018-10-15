package main

func main() {
	var myTodoList TodoList
	myTodoList.Add("Wake up")
	myTodoList.Add("Shower")
	myTodoList.Add("Have breakfast")
	myTodoList.Add("Go to work")
	myTodoList.Display()
	myTodoList.Remove("Have breakfast")
	myTodoList.Display()
	myTodoList.Rename("Shower", "Take bath")
	myTodoList.Display()
}
