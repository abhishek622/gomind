# GOMind CLI 🚀

A simple command-line application built with `Go` for seamless task management. This project integrates with `MongoDB` for data storage, utilizes `AWANLLM` for AI-powered task generation, and leverages the `Cobra` library to create a user-friendly CLI interface. Tasks are displayed in a clean, tabular format for better readability using `text/tabwriter`.


## 📺 Demo
![GOMind Demo](https://github.com/abhishek622/gomind/blob/main/assets/demo2.gif?raw=true)

---

## 🚀 Getting Started

### Prerequisites

Ensure you have the following installed and configured:

- **Go** installed on your machine.
- **MongoDB** running locally or remotely.
- **AWANLLM API key** for task generation (using it because it's free 😅).

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/abhishek622/gomind.git
   cd gomind
   ```
2. **Set up environment variables:**
   Create a `.env` file and configure your MongoDB and AWANLLM API key.
   ```bash
   # MongoDB database name
   echo DB_NAME=your-table-name-here > .env

   # MongoDB URI
   echo MONGODB_URI=your-mongodb-uri-here > .env

   # AWANLLM API Key
   echo AWANLLM_API_KEY=your-key-here > .env
   ```
3. **Build the application:**
   ```bash
   go build -o gomind main.go
   ```
4. **Run the CLI commands:**
   ```bash
   ./gomind help                     # List all available commands

   ./gomind add "Buy groceries"       # Add a new task

   ./gomind list                     # View all tasks

   ./gomind complete 1               # Mark a task as complete (ID: 1)

   ./gomind del 1                    # Delete a task by ID (or use 'all' to clear all tasks)

   ./gomind gen "Lunch at 4pm then study for 2 hours" # Generate tasks using AWANLLM
   ```

---

## 🐛 Bugs or Feature Requests

Spotted a bug? Have a feature idea? Let’s build this together!

- **Report issues:** Open a GitHub issue describing the bug or feature request.
- **Suggest improvements:** If you have ideas to enhance the CLI, create a ticket.
- **Contribute:** Pull requests are welcome! Let’s make task management smarter and smoother. 🚀
