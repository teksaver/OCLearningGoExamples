package main

func main() {
	var myTodoList = NewTodoList()
	myTodoList.Add("Wake up")
	myTodoList.Add("Shower")
	myTodoList.Add("Have breakfast")
	myTodoList.Add("Go to work")
	myTodoList.Display()
	myTodoList.MarkAsDone("Wake up")
	myTodoList.Remove("Have breakfast")
	myTodoList.Rename("Shower", "Take bath")
	myTodoList.MarkAsDone("Take bath")
	myTodoList.MarkAsDone("Go to work")
}
