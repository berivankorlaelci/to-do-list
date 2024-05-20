package models

import (
	"time"
)

type TodoList struct {
	ID             uint
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	CompletionRate float32
	Todos          []TodoItem
}

type TodoItem struct {
	ID          uint
	TodoListID  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Content     string
	IsCompleted bool
}
