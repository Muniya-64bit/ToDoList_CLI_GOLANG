package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task          string
	Done          bool
	CreatedDate   time.Time
	CompletedDate time.Time
}

type Todos []item

func (t *Todos) add(task string) {
	todo := item{
		Task:          task,
		Done:          false,
		CreatedDate:   time.Now(),
		CompletedDate: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("index out of range")
	}
	ls[index-1].CompletedDate = time.Now()
	ls[index-1].Done = true
	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("index out of range")
	}
	*t = append(ls[:index-1], ls[index:]...)
	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
	}
	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	fmt.Println("Hello and welcome!")

}
