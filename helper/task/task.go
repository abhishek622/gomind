package task

import (
	"time"
)

// Priority represents the priority level of a task
type Priority string

const (
	High   Priority = "High"
	Medium Priority = "Medium"
	Low    Priority = "Low"
)

type Task struct {
	ID           int64     `bson:"_id"`
	Description  string    `bson:"description"`
	Category     string    `bson:"category"`
	Priority     Priority  `bson:"priority"`
	DueDate      time.Time `bson:"due_date,omitempty"`
	Dependencies []string  `bson:"dependencies,omitempty"`
	Completed    bool      `bson:"completed"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}
