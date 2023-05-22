package main

import (
	"fmt"
	"os"
	"strconv"
)

//Declaring the Task struct
type Task struct {
	ID       int
	Name     string
	Complete bool
}

//Declaring a slice of type struct Task
var tasks []Task

func main() {

	for {
		fmt.Println("Command Options:")
		fmt.Println("1. Add a task")
		fmt.Println("2. Mark task as complete")
		fmt.Println("3. Remove a task")
		fmt.Println("4. Exit")

		fmt.Print("Enter command: ")
		var command int
		fmt.Scan(&command)

		switch command {
		case 1:
			fmt.Print("Enter task name: ")
			var taskName string
			fmt.Scan(&taskName)
			addTask(taskName)
		case 2:
			fmt.Print("Enter task ID: ")
			var taskIDStr string
			fmt.Scan(&taskIDStr)
			taskID, err := strconv.Atoi(taskIDStr)
			if err != nil {
				fmt.Println("Invalid task ID")
				continue
			}
			markTaskComplete(taskID)
		case 3:
			fmt.Print("Enter task ID: ")
			var taskIDStr string
			fmt.Scan(&taskIDStr)
			taskID, err := strconv.Atoi(taskIDStr)
			if err != nil {
				fmt.Println("Invalid task ID")
				continue
			}
			removeTask(taskID)
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
		}

		fmt.Println("Current tasks:")
		printTasks()
		fmt.Println()
	}
}

func addTask(name string) {
	task := Task{
		ID:       len(tasks) + 1,
		Name:     name,
		Complete: false,
	}
	tasks = append(tasks, task)
}

func markTaskComplete(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Complete = true
			return
		}
	}
	fmt.Println("Task not found")
}

func removeTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
	fmt.Println("Task not found")
}

func printTasks() {
	for i := range tasks {
		completeStatus := "Incomplete"
		if tasks[i].Complete {
			completeStatus = "Complete"
		}
		fmt.Printf("ID: %d | Name: %s | Status: %s\n", tasks[i].ID, tasks[i].Name, completeStatus)
	}
}
