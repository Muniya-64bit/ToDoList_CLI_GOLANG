package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/alexeyco/simpletable"
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

const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
)

func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
}

func (t *Todos) Print() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			{Align: simpletable.AlignCenter, Text: "CompletedAt"},
		},
	}
	var cells [][]*simpletable.Cell
	for idx, item := range *t {
		idx++
		task := blue(item.Task)
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: fmt.Sprintf("%t", item.Done)},
			{Text: item.CreatedDate.Format(time.RFC1123Z)},
			{Text: item.CompletedDate.Format(time.RFC1123Z)},
		})
	}
	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: "Your todos are here"},
	}}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
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
		todos.Print()
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
