package service

import (
	todo "github.com/POMBNK/restAPI"
	"github.com/POMBNK/restAPI/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) Create(userId int, listId int, item todo.TodoItem) (int, error) {
	return s.repo.Create(userId, listId, item)
}
