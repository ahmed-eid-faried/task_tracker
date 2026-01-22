package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Constants
const fileName = "tasks.json"

// Task struct matches the requirements
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // todo, in-progress, done
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func main() {
	// Ensure we have at least a command argument
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAdd(os.Args)
	case "list":
		handleList(os.Args)
	case "update":
		handleUpdate(os.Args)
	case "delete":
		handleDelete(os.Args)
	case "mark-in-progress":
		handleMarkStatus(os.Args, "in-progress")
	case "mark-done":
		handleMarkStatus(os.Args, "done")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

// --- Command Handlers ---

func handleAdd(args []string) {
	if len(args) < 3 {
		fmt.Println("Error: Missing task description")
		return
	}
	description := args[2]
	tasks := loadTasks()

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
	saveTasks(tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
}

func handleList(args []string) {
	tasks := loadTasks()

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

func handleUpdate(args []string) {
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

	tasks := loadTasks()
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

	saveTasks(tasks)
	fmt.Printf("Task %d updated successfully\n", id)
}

func handleDelete(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: task-cli delete <id>")
		return
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error: Invalid ID")
		return
	}

	tasks := loadTasks()
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

	saveTasks(newTasks)
	fmt.Printf("Task %d deleted successfully\n", id)
}

func handleMarkStatus(args []string, status string) {
	if len(args) < 3 {
		fmt.Println("Usage: task-cli mark-X <id>")
		return
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error: Invalid ID")
		return
	}

	tasks := loadTasks()
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

	saveTasks(tasks)
	fmt.Printf("Task %d marked as %s\n", id, status)
}

// --- File Storage Helpers ---

func loadTasks() []Task {
	file, err := os.ReadFile(fileName)
	if err != nil {
		// If file doesn't exist, return empty list
		if os.IsNotExist(err) {
			return []Task{}
		}
		panic(err)
	}

	// Handle empty file case
	if len(file) == 0 {
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println("Error parsing JSON file. It might be corrupt.")
		return []Task{}
	}
	return tasks
}

func saveTasks(tasks []Task) {
	// Marshal with indentation for readability
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}

	// 0644 is standard read/write permission
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		panic(err)
	}
}

func printUsage() {
	fmt.Println("Task Tracker CLI")
	fmt.Println("Usage:")
	fmt.Println("  add <description>")
	fmt.Println("  update <id> <description>")
	fmt.Println("  delete <id>")
	fmt.Println("  mark-in-progress <id>")
	fmt.Println("  mark-done <id>")
	fmt.Println("  list [done|todo|in-progress]")
}
