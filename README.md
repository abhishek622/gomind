# Task Manager CLI

A command-line application built in Go to manage tasks efficiently. This project integrates with MongoDB for data storage, leverages AWANLLM for task generation, and uses the Cobra library to create a robust CLI interface. Tasks are displayed in a clean, tabular format for better readability.

# PROJECT GOAL

- [x] Read/write tasks from mongoDB (mongodb)
- [x] Use AWANLLM to generate tasks
- [x] Use cobra to build CLI (cobra)
- [x] Printing data in tabular format (text/tabwriter)
- [x] Print time in human readable format (timediff)

## DEMO
![GOMind Demo](https://github.com/abhishek622/gomind/blob/main/assets/demo.gif?raw=true)

## Getting Started

### Prerequisites

- Go installed on your machine.
- MongoDB instance running locally or remotely.
- AWANLLM API key for generating tasks.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/abhishek622/gomind.git
   ```
2. Set the environment variables:
   ```bash
   # Set up your MongoDB table name
   echo DB_NAME=your-table-name-here > .env

   # Set up your MongoDB URI
   echo MONGODB_URI=your-mongodb-uri-here > .env

   # Set up your OpenAI key
   echo "AWANLLM_API_KEY=your-key-here" > .env
   ```
3. Build the application:
   ```bash
   go build -o gomind main.go
   ```
4. Use the CLI commands to manage tasks.
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
   ./gomind complete 1 # Mark task as complete (ID: 1)
   ```
   ```bash
   ./gomind del 1 # Delete a task (ID: 1 || all)
   ```
   ```bash
   ./gomind gen "Lunch at 4pm after that study for 2hrs" # Generate a task using AWANLLM
   ```
