package app

import "time"

// Todo types
type Todo struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Completed  bool      `json:"completed"`
	CreateTime time.Time `json:"createtime"`
	UpdateTime time.Time `json:"updatetime"`
}

// TodoList types
type TodoList struct {
	Todos []*Todo `json:"todos"`
}
