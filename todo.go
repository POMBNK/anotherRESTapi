package todo

import "errors"

// TodoList  Todo списки
type TodoList struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

// UserList сущность для связи пользователя и списка задач M to M
type UserList struct {
	ID     int
	UserId int
	ListId int
}

// TodoItem сущность элементов из списка задач
type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// ListItem сущность для связи списков и задач M to M
type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateList struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateItem struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (u *UpdateList) Validate() error {
	if u.Title == nil && u.Description == nil {
		return errors.New("nothing to update")
	}
	return nil
}

func (u *UpdateItem) Validate() error {
	if u.Title == nil && u.Description == nil && u.Done == nil {
		return errors.New("nothing to update")
	}
	return nil
}
