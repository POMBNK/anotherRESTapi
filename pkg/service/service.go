package service

import (
	todo "github.com/POMBNK/restAPI"
	"github.com/POMBNK/restAPI/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetByID(userId int, listId int) (todo.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, inputUpdate todo.UpdateList) error
}

type TodoItem interface {
	Create(userId int, listId int, item todo.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]todo.TodoItem, error)
	GetByID(userId int, itemId int) (todo.TodoItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, inputUpdate todo.UpdateItem) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem),
	}
}
