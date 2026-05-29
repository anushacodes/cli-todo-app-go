package main

import (
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time // pointer because it can be null
}

type Todos []Todo //like Todos = Todo[] in py

// methods
func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	// appending todo to it's struct
	*todos = append(*todos, todo)
}

// helper emthod to find index of todo by title
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {

		// err := errors.New("invalid index")
		err := fmt.Errorf("invalid index") //check diff?
		fmt.Println(err)
	}

	return nil
}

// delete function - we split at del index then join again after removing that index
func (todo *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todo = append(t[:index], t[index+1:]...)
	return nil

}
