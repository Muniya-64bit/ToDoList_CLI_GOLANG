package main

import (
	"encoding/json"
	"errors"
	"flag"
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

// Add adds a new task to the Todos list
func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedDate: time.Now(),
	}
	*t = append(*t, todo)
	fmt.Println("Task added successfully!")
}

// Complete marks a task as completed by its index
func (t *Todos) Complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("index out of range")
	}
	ls[index-1].CompletedDate = time.Now()
	ls[index-1].Done = true
	fmt.Printf("Task %d marked as completed!\n", index)
	return nil
}

// Delete removes a task by its index
func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("index out of range")
	}
	*t = append(ls[:index-1], ls[index:]...)
	fmt.Printf("Task %d deleted successfully!\n", index)
	return nil
}

// Load reads Todos from a JSON file
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

// Store writes Todos to a JSON file
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

// List displays all tasks with their statuses
func (t *Todos) List() {
	fmt.Println("Todo List:")
	for i, todo := range *t {
		status := "Incomplete"
		if todo.Done {
			status = "Completed"
		}
		fmt.Printf("%d. %s (Status: %s, Created: %s)\n", i+1, todo.Task, status, todo.CreatedDate.Format(time.RFC1123))
	}
}

func main() {
	// Define flags for CLI commands
	add := flag.String("add", "", "Add a task to the list")
	complete := flag.Int("complete", 0, "Mark a task as completed (provide index)")
	delete := flag.Int("delete", 0, "Delete a task (provide index)")
	list := flag.Bool("list", false, "List all tasks")

	// Parse the command-line flags
	flag.Parse()

	// Load existing Todos from file
	todos := &Todos{}
	err := todos.Load("todos.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading todos:", err)
		os.Exit(1)
	}

	// Handle user commands
	switch {
	case *add != "":
		todos.Add(*add)
	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	case *delete > 0:
		err := todos.Delete(*delete)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	case *list:
		todos.List()
	default:
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Save Todos back to file
	err = todos.Store("todos.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error storing todos:", err)
		os.Exit(1)
	}
}
