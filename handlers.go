package main

import (
	"fmt"
	"strconv"
	"time"
)

// --- Command Handlers ---

func HandleAdd(args []string) {
	if len(args) < 3 {
		fmt.Println("Error: Missing task description")
		return
	}
	description := args[2]
	tasks := LoadTasks()

	// Determine new ID
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	newTask := Task{
		ID:          maxID + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	SaveTasks(tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
}

func HandleList(args []string) {
	tasks := LoadTasks()

	// Check if a filter is applied (done, todo, in-progress)
	filter := ""
	if len(args) > 2 {
		filter = args[2]
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Printf("%-5s %-20s %-15s %s\n", "ID", "Status", "Created", "Description")
	fmt.Println("----------------------------------------------------------------------")

	for _, t := range tasks {
		if filter != "" && t.Status != filter {
			continue
		}
		// Format date for display
		dateStr := t.CreatedAt.Format("2006-01-02 15:04")
		fmt.Printf("%-5d %-20s %-15s %s\n", t.ID, t.Status, dateStr, t.Description)
	}
}

func HandleUpdate(args []string) {
	if len(args) < 4 {
		fmt.Println("Usage: task-cli update <id> <new_description>")
		return
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error: Invalid ID")
		return
	}
	newDesc := args[3]

	tasks := LoadTasks()
	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Description = newDesc
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}

	SaveTasks(tasks)
	fmt.Printf("Task %d updated successfully\n", id)
}

func HandleDelete(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: task-cli delete <id>")
		return
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error: Invalid ID")
		return
	}

	tasks := LoadTasks()
	newTasks := []Task{}
	found := false

	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue // Skip adding this task to the new slice
		}
		newTasks = append(newTasks, t)
	}

	if !found {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}

	SaveTasks(newTasks)
	fmt.Printf("Task %d deleted successfully\n", id)
}

func HandleMarkStatus(args []string, status string) {
	if len(args) < 3 {
		fmt.Println("Usage: task-cli mark-X <id>")
		return
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error: Invalid ID")
		return
	}

	tasks := LoadTasks()
	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}

	SaveTasks(tasks)
	fmt.Printf("Task %d marked as %s\n", id, status)
}