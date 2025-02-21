# PROJECT GOAL

- Read/write data file system
- Printing data in tabular format
- CLI App with multiple command
  - add
  - complete
  - delete
  - help
  - list

# Implementation

- Use CSV to store data (encoding/csv)
- To print text into tabular format (text/tabwriter)
- To print time in human readable format (timediff)
- To build command-line interface (cobra)

# Project structure

```
todo-cli/
├── cmd/
│ ├── root.go # Root command for Cobra
│ ├── add.go # Add command for adding todos
│ ├── list.go # List command for listing todos
│ ├── complete.go # Complete command for marking todos as completed
│ └── delete.go # Delete command for deleting todos
├── internal/
│ ├── todo/
│ │ ├── todo.go # Todo-related logic (e.g., struct, CRUD operations)
│ │ └── storage.go # Storage logic (e.g., reading/writing JSON file)
│ └── utils/
│   └── time.go # Time-related utilities (e.g., formatting, timediff)
├── data/
│ └── todos.json # JSON file to store todos
├── go.mod # Go module file
├── go.sum # Go dependencies checksum file
├── main.go # Entry point of the application
└── README.md # Project documentation
```
