package todo

import "gorm.io/gorm"

type TodoDB struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoStorage struct {
	db *gorm.DB
}

func NewTodoStorage(db *gorm.DB) *TodoStorage {
	return &TodoStorage{
		db: db,
	}
}
