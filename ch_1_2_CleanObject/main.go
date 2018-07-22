package main

func main() {
	var myTodoList TodoList
	myTodoList.Add("Wake up")
	myTodoList.Add("Shower")
	myTodoList.Add("Have breakfast")
	myTodoList.Add("Go to work")
	myTodoList.Display()
	myTodoList.Remove(2)
	myTodoList.Display()
	myTodoList.Rename(1, "Take bath")
	myTodoList.Display()
}
