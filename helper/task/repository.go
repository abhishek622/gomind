package task

import (
	"abhishek622/gomind/utils"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository provides an interface to interact with the database
type Repository struct {
	Collection *mongo.Collection
}

// initializes a new repository instance
func NewRepository() *Repository {
	return &Repository{
		Collection: utils.GetCollection("tasks"),
	}
}

// GetNextID generates the next unique ID
func (r *Repository) GetNextID() (int64, error) {
	var result struct {
		MaxID int64 `bson:"max_id"`
	}
	err := r.Collection.FindOne(context.TODO(), bson.M{}, nil).
		Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 1, nil // Start from 1 if no tasks exist
		}
		return 0, err
	}

	return result.MaxID + 1, nil
}

func (r *Repository) CreateTask(task *Task) error {
	nextID, err := r.GetNextID()
	if err != nil {
		return err
	}
	task.ID = nextID
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	_, err = r.Collection.InsertOne(context.TODO(), task)
	if err != nil {
		return fmt.Errorf("failed to insert todo: %v", err)
	}

	return nil
}

func (r *Repository) GetTasks() ([]Task, error) {
	var tasks []Task

	cursor, err := r.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks: %v", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task Task
		if err := cursor.Decode(&task); err != nil {
			log.Println("Error decoding task:", err)
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *Repository) MarkAsCompleted(todoID int64) error {
	filter := bson.M{"_id": todoID}
	update := bson.M{"$set": bson.M{"completed": true, "updated_at": time.Now()}}

	_, err := r.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to mark task as completed: %v", err)
	}

	return nil
}
