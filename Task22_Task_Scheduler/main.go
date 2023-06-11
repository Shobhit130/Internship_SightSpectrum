package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID        int
	Time      time.Time
	Name      string
	IsRunning bool
}

var tasks []Task
var taskIDCounter int
var wg sync.WaitGroup
var taskDoneCh = make(chan struct{})
var exitFlag bool // Flag to indicate if the user requested to exit

func main() {
	fmt.Println("Task Scheduler")

	go scheduleTasks()

	for {
		if !showMenu() {
			break
		}
	}

	exitFlag = true // Set the exit flag

	// Check if all tasks have been completed
	for {
		if len(tasks) == 0 {
			break
		}
		time.Sleep(1 * time.Second)
	}

	if exitFlag {
		fmt.Println("Exiting...")
	}
}

func showMenu() bool {
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		addTask()
	case 2:
		viewTasks()
	case 3:
		fmt.Println("Exiting program after completing all scheduled tasks...")
		return false
	default:
		fmt.Println("Invalid choice. Please try again.")
	}

	return true
}

func addTask() {
	taskIDCounter++

	var name string
	fmt.Print("Enter task name: ")
	fmt.Scanln(&name)

	var year, month, day, hour, min int
	fmt.Println("Enter task schedule:")
	fmt.Print("Year: ")
	fmt.Scanf("%d\n", &year)
	fmt.Print("Month: ")
	fmt.Scanf("%d\n", &month)
	fmt.Print("Day: ")
	fmt.Scanf("%d\n", &day)
	fmt.Print("Hour: ")
	fmt.Scanf("%d\n", &hour)
	fmt.Print("Minute: ")
	fmt.Scanf("%d\n", &min)

	t := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	task := Task{
		ID:        taskIDCounter,
		Time:      t,
		Name:      name,
		IsRunning: false,
	}
	tasks = append(tasks, task)

	fmt.Printf("Task '%s' scheduled for %s.\n", task.Name, task.Time.Format(time.RFC1123))
}

func viewTasks() {
	fmt.Println("Task List:")

	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	for _, task := range tasks {
		fmt.Printf("ID: %d, Name: %s, Time: %s\n", task.ID, task.Name, task.Time.Format(time.RFC1123))
	}
}

func scheduleTasks() {
	for {
		currentTime := time.Now()

		for i, task := range tasks {
			if !task.IsRunning && currentTime.After(task.Time) {
				tasks[i].IsRunning = true
				wg.Add(1)
				go executeTask(task)
			}
		}

		time.Sleep(1 * time.Second)

		// Check if all tasks have been completed and the exit flag is set
		if len(tasks) == 0 && exitFlag {
			taskDoneCh <- struct{}{} // Signal that all tasks are completed
			break
		}
	}
}

func executeTask(task Task) {
	defer wg.Done()

	fmt.Printf("\nExecuting task '%s'...\n", task.Name)
	time.Sleep(2 * time.Second) // Simulating task execution time

	fmt.Printf("Task '%s' completed.\n", task.Name)

	// Remove task from the list
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	// Check if all tasks have been completed and the exit flag is set
	if len(tasks) == 0 && exitFlag {
		taskDoneCh <- struct{}{} // Signal that all tasks are completed
	}
}
