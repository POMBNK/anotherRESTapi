package repository

import (
	"fmt"
	todo "github.com/POMBNK/restAPI"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(userId int, listId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int

	createItemQuery := fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1,$2) RETURNING id ", todoItemTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err = row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListsItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id,item_id) VALUES ($1,$2)", listsItemsTable)
	_, err = tx.Exec(createListsItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userId int, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem

	query := fmt.Sprintf(
		`SELECT ti.id,ti.title,ti.description,ti.done FROM %s ti
   			   JOIN %s li ON ti.id = li.item_id
   			   JOIN %s ul ON ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		todoItemTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}
