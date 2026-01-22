package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// --- File Storage Helpers ---

func LoadTasks() []Task {
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

func SaveTasks(tasks []Task) {
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

func PrintUsage() {
	fmt.Println("Task Tracker CLI")
	fmt.Println("Usage:")
	fmt.Println("  add <description>")
	fmt.Println("  update <id> <description>")
	fmt.Println("  delete <id>")
	fmt.Println("  mark-in-progress <id>")
	fmt.Println("  mark-done <id>")
	fmt.Println("  list [done|todo|in-progress]")
}
