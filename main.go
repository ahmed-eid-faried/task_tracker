package main

import (
	"fmt"
	"os"
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
		PrintUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		HandleAdd(os.Args)
	case "list":
		HandleList(os.Args)
	case "update":
		HandleUpdate(os.Args)
	case "delete":
		HandleDelete(os.Args)
	case "mark-in-progress":
		HandleMarkStatus(os.Args, "in-progress")
	case "mark-done":
		HandleMarkStatus(os.Args, "done")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		PrintUsage()
	}
}
