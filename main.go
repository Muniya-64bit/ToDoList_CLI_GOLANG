package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task          string    `json:"task"`
	Done          bool      `json:"done"`
	CreatedDate   time.Time `json:"created_date"`
	CompletedDate time.Time `json:"completed_date,omitempty"`
}

type Todos []item

// Add a new task to the Todos list
func (t *Todos) add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedDate: time.Now(),
	}
	*t = append(*t, todo)
}

// Mark a task as completed by its index
func (t *Todos) complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("index out of range")
	}
	ls[index-1].CompletedDate = time.Now()
	ls[index-1].Done = true
	return nil
}

// Delete a task by its index
func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("index out of range")
	}
	*t = append(ls[:index-1], ls[index:]...)
	return nil
}

// Load Todos from a JSON file
func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// File doesn't exist, return without error
			return nil
		}
		return err
	}

	// Unmarshal file content into Todos
	err = json.Unmarshal(file, t)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %w", err)
	}
	return nil
}

// Store Todos to a JSON file
func (t *Todos) Store(filename string) error {
	// Marshal Todos into JSON
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	// Write JSON to the file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}

func main() {
	todos := &Todos{}

	// Load existing Todos from file
	err := todos.Load("todos.json")
	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	// Add a new task
	todos.add("Learn Go")

	// Mark the first task as complete
	err = todos.complete(1)
	if err != nil {
		fmt.Println("Error completing task:", err)
	}

	// Delete a task
	err = todos.Delete(1)
	if err != nil {
		fmt.Println("Error deleting task:", err)
	}

	// Store updated Todos back to file
	err = todos.Store("todos.json")
	if err != nil {
		fmt.Println("Error storing todos:", err)
	}

	fmt.Println("Todos updated successfully!")
}
