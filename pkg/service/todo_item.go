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

func (s *TodoItemService) GetAll(userId int, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetByID(userId int, itemId int) (todo.TodoItem, error) {
	return s.repo.GetByID(userId, itemId)
}

func (s *TodoItemService) Delete(userId int, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId int, itemId int, inputUpdate todo.UpdateItem) error {
	if err := inputUpdate.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, itemId, inputUpdate)
}
