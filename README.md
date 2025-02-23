# Task Manager CLI

A command-line application built in Go to manage tasks efficiently. This project integrates with MongoDB for data storage, leverages GPT-3 for task generation, and uses the Cobra library to create a robust CLI interface. Tasks are displayed in a clean, tabular format for better readability.

# PROJECT GOAL

- [x] Read/write tasks from mongoDB (mongodb)
- [ ] Use GPT-3 to generate tasks
- [x] Use cobra to build CLI (cobra)
- [x] Printing data in tabular format (text/tabwriter)
- [x] Print time in human readable format (timediff)
- [ ] CLI App with multiple command
  - [x] add
  - [x] complete
  - [ ] delete
  - [x] help
  - [x] list
  - [ ] generate (GPT-3)

# Project structure

```
gomind/
│── cmd/               # CLI command handlers
│   ├── add.go         # Handles adding a task
│   ├── list.go        # Handles listing tasks
│   ├── complete.go    # Handles marking tasks as complete
│   └── root.go        # Root command setup
│── helper/          # Core Application Logic 
│   ├── gpt/           # GPT Integration
│   │   ├── parser.go  # Handles GPT parsing of natural language
│   │   └── client.go  # OpenAI API interaction
│   └── task/          # Task Management Logic
│       ├── task.go    # Task struct & related methods
│       ├── manager.go # Task CRUD operations (Add, Delete, Update)
│       └── storage.go # Handles saving/loading tasks (MongoDB)
│── utils/             # Utility functions (Reusable across project)
│   ├── logger/        # Logging utilities
│   ├── timeparser/    # Parses natural language dates (e.g., "tomorrow 5 PM")
│   └── config/        # Configuration utilities (API keys, env vars)
│── .env               # Environment variables
│── go.mod             # Go module file
│── go.sum             # Go dependencies
│── main.go            # Application entry point
```

## Getting Started

### Prerequisites

- Go installed on your machine.
- MongoDB instance running locally or remotely.
- OpenAI API key for GPT-3 integration.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/abhishek622/gomind.git
   ```
2. Build the application:
   ```bash
   go build -o gomind main.go
   ```
3. Use the CLI commands to manage tasks.
  ```bash
  ./gomind help # List of commands
  ```
  ```bash
  ./gomind add "Buy groceries" # Add a task
  ```
  ```bash
  ./gomind list # List all tasks
  ```
  ```bash
  ./gomind complete 1 # Mark task as complete
  ```
  ```bash
  ./gomind delete 1 # Delete a task
  ```
  ```bash
  ./gomind generate "Lunch at 4pm after that study for 2hrs" # Generate a task using GPT-3
  ```
