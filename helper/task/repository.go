package task

import (
	"abhishek622/gomind/utils"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// Define a filter to sort by _id in descending order and limit to 1 result
	opts := options.FindOne().SetSort(bson.M{"_id": -1})

	var result struct {
		ID int64 `bson:"_id"`
	}

	// Find the document with the maximum _id
	err := r.Collection.FindOne(context.TODO(), bson.M{}, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 1, nil // If no documents exist, start from 1
		}
		return 0, err // Return any other error
	}

	// Return the next ID
	return result.ID + 1, nil
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

func (r *Repository) DeleteATask(todoID int64) error {
	filter := bson.M{"_id": todoID}
	_, err := r.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete a task: %v", err)
	}

	return nil
}

func (r *Repository) DeleteAllTask() error {
	_, err := r.Collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return fmt.Errorf("failed to delete all tasks: %v", err)
	}

	return nil
}

func (r *Repository) CheckDependenciesExist(depIDs []int64) bool {
	filter := bson.M{"_id": bson.M{"$in": depIDs}}
	count, err := r.Collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		fmt.Println("Database error while checking dependencies:", err)
		return false
	}

	return count == int64(len(depIDs))
}
