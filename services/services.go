package services

import (
	"context"
	"database/sql"
	"log"
)

type TodoListService interface {
	GetTodoList(ctx context.Context) ([]TodoItem, error)
}

type MyTodoListService struct {
	DB *sql.DB
}

type TodoItem struct {
	ID        int    `json:"id"`
	Task_Name string `json:"task_name"`
	Priority  string `json:"priority"`
}

func (s *MyTodoListService) GetTodoList(ctx context.Context) ([]TodoItem, error) {
	db := s.DB // ponteiro que carrega conexão com banco de dados

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	var todoItems []TodoItem

	for rows.Next() {
		var todoitem TodoItem
		err := rows.Scan(&todoitem.ID, &todoitem.Task_Name, &todoitem.Priority)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		todoItems = append(todoItems, todoitem)
	}

	return todoItems, nil
}
