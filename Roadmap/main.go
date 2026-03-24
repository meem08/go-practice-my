package main

// "errors"

// // Define a Task struct with four fields: ID (int), Title (string), Description (string), and IsCompleted (bool).
// 	type Task struct {
// 		ID int
// 		Title string
// 		Description string
// 		IsCompleted bool
// 	}

func main() {
	// 	var currenttask int = 20
	// 	lasttask := 5.39
	// 	totaltask := currenttask + int(lasttask)
	// 	fmt.Println(totaltask)

	// 	var taskPriority string = "low"
	// 	switch taskPriority {
	// 	case "high":
	// 		fmt.Println("Taskpriority: 1")
	// 	case "medium":
	// 		fmt.Println("Taskpriority: 2")
	// 	case "low":
	// 		fmt.Println("Taskpriority: 3")
	// 	default:
	// 		fmt.Println(0)
	// 	}

	// 	var taskCompleted bool = true
	// 	if taskCompleted == true {
	// 		fmt.Println("Great job!")
	// 	} else {
	// 		fmt.Println("Get to work!")
	// 	}

	// 	// Write a program that defines a slice of five strings representing task names. Use a for loop with range to print each task name preceded by "Task: ".

	// 	doslice := []string{"sweep", "wash", "groceries", "car-repair", "market"}
	// 	for _, do:= range doslice {
	// 		fmt.Printf("Task : %v\n", do)
	// 	}
	// 	// Create a for loop that acts like a counter, printing numbers from 10 down to 1.
	// 	for i := 10; i>= 1; i-- {
	// 		fmt.Println(i)
	// 	}
	// 	// Write a program that defines an integer variable attempts = 0. Use an infinite for loop that increments attempts and prints the value. Add an if statement to break the loop once attempts reaches 5.
	// 	var attempts = 0
	// 	for {
	// 		attempts +=1
	// 		if attempts == 5 {
	// 			break
	// 		}
	// 		fmt.Println(attempts)
	// 	}
	// 	// Create a slice of strings named todoList and add three tasks to it. Print the length and capacity using len() and cap().
	// 	todoList := []string{"sweep", "wash", "groceries"}
	// 	fmt.Println(len(todoList), cap(todoList))
	// 	// Given a slice data := []int{10, 20, 30, 40, 50}, create a new slice that contains only [20, 30, 40].
	// 	data := []int{10,20,30,40,50}
	// 	new := data[1:4]
	// 	fmt.Println(new)

	// 	// Create a slice with a capacity of 5 using make(), add 5 items to it, and then add a 6th item. Observe how the capacity changes (hint: print cap() before and after the 6th addition).
	// 	slinew := make([]int, 5)
	// 	slinew = append(slinew, 4)
	// 	fmt.Println(cap(slinew))

	// 	// Create a slice of these tasks and write a loop to print only the tasks where IsCompleted is false.
	// 	taskList := []Task{
	// 		{ID: 1, Title: "Go", Description: "Basics of Go", IsCompleted: true},
	// 	 	{ID: 2, Title: "python", Description: "Basics of python", IsCompleted: false},
	// 	 	{ID: 3, Title: "C++", Description: "Basics of C++", IsCompleted: true},
	// 	 	{ID: 4, Title: "Java", Description: "Basics of Java", IsCompleted: false},
	// 	}
	// 	for _,task := range taskList{
	// 		if !task.IsCompleted {
	// 			fmt.Println(task.ID, task.Title,task.Description, task.IsCompleted)
	// 		}
	// 	}
	// 	// Create a map[int]Task, populate it with three tasks, then implement a function that deletes a task by ID safely, ensuring the program doesn't crash if the ID doesn't exist.
	// 	taskMap := make(map[int]Task)
	// 	taskMap[1] = Task{ID: 1, Title: "Go", Description: "Basics of Go", IsCompleted: true}
	// 	taskMap[2] = Task{ID: 2, Title: "python", Description: "Basics of python", IsCompleted: false}
	// 	taskMap[3] = Task{ID: 3, Title: "C++", Description: "Basics of C++", IsCompleted: true}

	// 	DeleteTask(taskMap, 3)

	// 	fmt.Println(taskMap)
	// 	// Write a main function that calls Divide with 10 and 2, and another time with 10 and 0. Print the results or the errors accordingly.
	// 	fmt.Println(Divide(10,0))

	// 	// Create a User struct with fields for Username, Email, and Age. Instantiate a new user and print their email address to the console.
	// 	type Userf struct{
	// 		Username string
	// 		Email string
	// 		Age int
	// 	}
	// 	User1 := Userf{"Bakinsuy", "blessing.akinsuyi01@gmail.com", 23}
	// 	fmt.Println(User1.Email)
	// 	// Define a Project struct that contains a Name (string) and a slice of Task structs. Initialize a project with two tasks inside it and iterate through them to print their titles.
	// 	type Project struct{
	// 		Name string
	// 		Task Task
	// 	}

	// 	 ProjectList := []Project{
	//           Project{
	// 		"Go-learn", Task{ID: 1, Title: "Go", Description: "Basics of Go", IsCompleted: true},
	// 	},
	// 	Project{
	// 		"Python-learn", Task{ID: 1, Title: "Python", Description: "Basics of python", IsCompleted: false},
	// 	},
	// 	}

	// 	for _,list := range ProjectList{
	// 		fmt.Println(list.Task.Title)
	// 	}
	// 	user := User{Username: "Aboy", Email: "aboy20@gmail.com"}
	// 	fmt.Println(user.Display())
	// 	u := User{
	// 		Username: "Seun",
	// 		Email: "seungood@gmail.com",
	// 	}
	// 	fmt.Printf("Before: %+v\n", u)
	// 	u.UpdateEmail("goodseun@gmail.com")
	// 	fmt.Printf("After : %+v\n", u)

	// 	var num Myint= 10
	// 	fmt.Println(num.Double())

	// 	// Write a main function that initializes a slice of Task structs, marks one as complete using your method, and prints its status using the interface.
	// 	taskIn := []Task{
	// 		{ID: 1, Title: "Go", Description: "Basics of Go", IsCompleted: true},
	// 	 	{ID: 2, Title: "python", Description: "Basics of python", IsCompleted: false},
	// 	 	{ID: 3, Title: "C++", Description: "Basics of C++", IsCompleted: true},
	// 	 	{ID: 4, Title: "Java", Description: "Basics of Java", IsCompleted: false},
	// 	}
	// 	var compact []Completable
	// 	for i:= range taskIn {
	// 		compact = append(compact, &taskIn[i])
	// 	}
	// 	for _, im := range compact{
	// 		fmt.Printf("Task done: %v\n", im.IsDone())
	// 	}

	// 	FizzBuzz(20)

	// 	// Implement a loop that iterates from 1 to 20. Use continue to skip even numbers and break to stop the loop if the counter hits 15.
	// 	for i:= 1; i <= 20; i++ {
	// 		if i%2 == 0 {
	// 		continue
	// 		}
	// 	}
	// 	res := struct {
	// 	Name string
	// 	Email string
	//      }{
	// 	Name: "Bola",
	// 	Email: "bolala12@gmail.com",
	//     }
	// 	fmt.Println(res)
	// res.Dislay()
	// fmt.Printf("Name: %s\n Email: %s\n", res.Name, res.Email)

}

// func DeleteTask(taskMap map[int]Task, ID int){
// 	value, ok := taskMap[ID]
// 	if ok {
// 		delete(taskMap, ID)
// 		fmt.Println("ID not found", value.ID)
// 	} else {
// 		fmt.Println("ID found")
// 	}
// }

// // Create a function Divide(a, b float64) (float64, error). If b is 0, return 0 and a descriptive error using errors.New(). Otherwise, return the result of the division and nil.
// func Divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0.0, errors.New("A float number should be passed")
// 	} else {
// 		return a/b, nil
// 	}
// }

// // Create a User struct with fields Username and Email. Define a method Display() that returns a string formatted as "User: [Username] (<[Email]>)".
// type User struct {
// 	Username string
// 	Email string
// }
// func(t User)Display() string {
// 	return "User: " + t.Username + "(" + t.Email + ")"
// }
// // Add a ChangeEmail(newEmail string) method to your User struct using a pointer receiver. Write a main function that creates a user, changes their email, and prints the result to verify the update.
// func (t *User) ChangeEmail(newEmail string) {
// 	t.Email= newEmail

// }
// // Attempt to define a method on a basic type (e.g., type MyInt int). Can you define a Double() method on it?
// type Myint int

// func (m Myint) Double()Myint{
// 	return m*2
// }
// // Create a User struct and add a method UpdateEmail(newEmail string) that uses a pointer receiver to change the user's email address.
// func (u *User) UpdateEmail(newEmail string) {
// 	u.Email = newEmail
// }
// // Define an interface named Completable with a method IsDone() bool. Update the Task struct to satisfy this interface by implementing IsDone() to return the value of the Completed field.
// type Completable interface{
// 	IsDone() bool
// }
// func (t Task) IsDone() bool {
// 	return t.IsCompleted
// }
// // Write a function that accepts an integer n and prints "Fizz" if n is divisible by 3, "Buzz" if divisible by 5, and "FizzBuzz" if divisible by both. Use a switch statement for the logic.
// func FizzBuzz(n int) {
// 	if n%3 == 0 && n%5 == 0 {
// 		fmt.Println("FizzBuzz")
// 	} else if n%3 == 0 {
// 		fmt.Println("Fizz")
// 	} else if n%5 == 0 {
// 		fmt.Println("Buzz")
// 	}
// }
// Create a User struct with Name and Email fields. Implement a Display() method that prints the user details.
// type Res struct {
// 	Name string
// 	Email string
// }
// var res = Res{
// 	Name: "Bola",
// 	Email: "bolala12@gmail.com",
// }
// func (w Res) Dislay()string{
// 	return fmt.Sprintf("Name: %v\n Email: %v\n", w.Name, w.Email)
// }
//
// Define an Emailer interface with a Send() method. Create a Newsletter struct and implement the Send() method on it using a pointer receiver to simulate changing a SentCount field.
type Emailer interface {
	Send()
}
type Newsletter struct {
	Title     string
	SendCount int
}
