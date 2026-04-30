package main

import "time"

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time // pointer because it can be null
}

type Todos []Todo //like Todos = Todo[] in py
