# Task Tracker

## Overview
Task Tracker is a simple command-line application written in Go that helps users manage their tasks efficiently. Users can add, remove, and list tasks from the command line.

## Features
- Add new tasks
- Remove existing tasks
- List all tasks
- Mark tasks as completed

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/ahmed-eid-faried/task_tracker.git
    cd task_tracker
    ```
2. Install dependencies:
    ```bash
    go mod tidy
    ```

## Usage
    To run the application, use the following command:

    ```bash
    go run main.go <command> [arguments]
    ```

   ### Available Commands
    - `add <task>` - Add a new task
    - `list` - Display all tasks
    - `update <task-id> <new-task>` - Update an existing task
    - `delete <task-id>` - Delete a task
    - `mark-in-progress <task-id>` - Mark a task as in progress
    - `mark-done <task-id>` - Mark a task as completed

### Examples
```bash
# Build the application
go build -o task-cli main.go

# Add new tasks
task-cli add "Buy groceries"
task-cli add "Learn Go"

# Mark tasks with status changes
task-cli mark-in-progress 2
task-cli mark-done 1

# Update an existing task
task-cli update 1 "Buy groceries and cook dinner"

# List tasks (all or filtered by status)
task-cli list
task-cli list todo
task-cli list done
task-cli list in-progress

# Delete a task
task-cli delete 1
```

## Project Roadmap
[Task Tracker Project Page](https://roadmap.sh/projects/task-tracker)

## Contributing
Contributions are welcome! Please open an issue or submit a pull request.

## License
This project is licensed under the MIT License.