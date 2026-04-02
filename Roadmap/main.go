package main

import "fmt"

const Maxtask = 100
func main() {
	var currenttask int = 20
	lasttask := 5.39
	totaltask := currenttask + int(lasttask)
	fmt.Println(totaltask)

	var taskPriority string = "low"
	switch taskPriority {
	case "high":
		fmt.Println("Taskpriority: 1")
	case "medium":
		fmt.Println("Taskpriority: 2")
	case "low":
		fmt.Println("Taskpriority: 3")
	default:
		fmt.Println(0)
	}

	var taskCompleted bool = true
	if taskCompleted == true {
		fmt.Println("Great job!")
	} else {
		fmt.Println("Get to work!")
	}

	// Write a program that defines a slice of five strings representing task names. Use a for loop with range to print each task name preceded by "Task: ".

	doslice := []string{"sweep", "wash", "groceries", "car-repair", "market"}
	for _, do:= range doslice {
		fmt.Printf("Task : %v\n", do)	
	}
	// Create a for loop that acts like a counter, printing numbers from 10 down to 1.
	for i := 10; i>= 1; i-- {
		fmt.Println(i)
	}
	// Write a program that defines an integer variable attempts = 0. Use an infinite for loop that increments attempts and prints the value. Add an if statement to break the loop once attempts reaches 5.
	var attempts = 0
	for {
		attempts +=1
		if attempts == 5 {
			break
		}
		fmt.Println(attempts)	
	}
	// Create a slice of strings named todoList and add three tasks to it. Print the length and capacity using len() and cap().
	todoList := []string{"sweep", "wash", "groceries"}
	fmt.Println(len(todoList), cap(todoList))
	// Given a slice data := []int{10, 20, 30, 40, 50}, create a new slice that contains only [20, 30, 40].
	data := []int{10,20,30,40,50}
	new := data[1:4]
	fmt.Println(new)
	// Create a slice with a capacity of 5 using make(), add 5 items to it, and then add a 6th item. Observe how the capacity changes (hint: print cap() before and after the 6th addition).
	slinew := make([]int, 5)
	slinew = append(slinew, 4)
	fmt.Println(cap(slinew))

	// Define a Task struct with four fields: ID (int), Title (string), Description (string), and IsCompleted (bool).
	type Task struct {
		ID int
		Title string
		Description string
		IsCompleted bool
	}
	// Create a slice of these tasks and write a loop to print only the tasks where IsCompleted is false.
	taskList := []Task{
		{ID: 1, Title: "Go", Description: "Basics of Go", IsCompleted: true},
	 	{ID: 2, Title: "python", Description: "Basics of python", IsCompleted: false},
	 	{ID: 3, Title: "C++", Description: "Basics of C++", IsCompleted: true},
	 	{ID: 4, Title: "Java", Description: "Basics of Java", IsCompleted: false},
	}
	for _,task := range taskList{
		if !task.IsCompleted {
			fmt.Println(task.IsCompleted)
		}
	}
	fmt.Println()
}
