package todo

import "time"

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
	CreatedAt string `json:"createdAt"`
}

func NewTodo(task string) *Todo {
	return &Todo{
		Task:      task,
		Completed: false,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}
