package todo

import (
	"encoding/json"
	"os"
)

const db = "data/todos.json"

func ReadTodos() ([]Todo, error) {
	file, err := os.ReadFile(db)
	if err != nil {
		return nil, err
	}

	var todos []Todo
	if err := json.Unmarshal(file, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func WriteTodos(todos []Todo) error {
	file, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(db, file, 0644)
}
