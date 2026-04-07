package main

import (
	"errors"
	"fmt"
)

// Define a Task struct with four fields: ID (int), Title (string), Description (string), and IsCompleted (bool).
	type Task struct {
		ID int
		Title string
		Description string
		IsCompleted bool
	}


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
	// Create a map[int]Task, populate it with three tasks, then implement a function that deletes a task by ID safely, ensuring the program doesn't crash if the ID doesn't exist.
	taskMap := make(map[int]Task)
	taskMap[1] = Task{ID: 1, Title: "Go", Description: "Basics of Go", IsCompleted: true}
	taskMap[2] = Task{ID: 2, Title: "python", Description: "Basics of python", IsCompleted: false}
	taskMap[3] = Task{ID: 3, Title: "C++", Description: "Basics of C++", IsCompleted: true}

	DeleteTask(taskMap, 3)
		
	fmt.Println(taskMap)
	// Write a main function that calls Divide with 10 and 2, and another time with 10 and 0. Print the results or the errors accordingly.
	fmt.Println(Divide(10,0))

	// Create a User struct with fields for Username, Email, and Age. Instantiate a new user and print their email address to the console.
	type User struct{
		Username string
		Email string
		Age int
	}
	User1 := User{"Bakinsuy", "blessing.akinsuyi01@gmail.com", 23}
	fmt.Println(User1.Email)
	// Define a Project struct that contains a Name (string) and a slice of Task structs. Initialize a project with two tasks inside it and iterate through them to print their titles.
	type Project struct{
		Name string
		Task Task
	}
     
	 ProjectList := []Project{
          Project{
		"Go-learn", Task{ID: 1, Title: "Go", Description: "Basics of Go", IsCompleted: true},	
	}, 
	Project{
		"Python-learn", Task{ID: 1, Title: "Python", Description: "Basics of python", IsCompleted: false},	
	},
	}
	

	for _,list := range ProjectList{
		fmt.Println(list.Task.Title)
	}

	
}

func DeleteTask(taskMap map[int]Task, ID int){
	value, ok := taskMap[ID]
	if ok {
		delete(taskMap, ID)
		fmt.Println("ID not found", value.ID)
	} else {
		fmt.Println("ID found")
	}
}

// Create a function Divide(a, b float64) (float64, error). If b is 0, return 0 and a descriptive error using errors.New(). Otherwise, return the result of the division and nil.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("A float number should be passed")
	} else {
		return a/b, nil
	}
}

// Create a User struct with fields Username and Email. Define a method Display() that returns a string formatted as "User: [Username] (<[Email]>)".