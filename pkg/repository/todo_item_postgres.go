package repository

import (
	"fmt"
	todo "github.com/POMBNK/restAPI"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (r *TodoItemPostgres) GetByID(userId int, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf(
		`SELECT ti.id,ti.title,ti.description,ti.done FROM %s ti
   			   	JOIN %s li ON ti.id = li.item_id
   			   	JOIN %s ul ON ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2 `,
		todoItemTable, listsItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *TodoItemPostgres) Delete(userId int, itemId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s ti USING %s li, %s ul 
       			WHERE ti.id=li.item_id AND ul.list_id = li.list_id AND ul.user_id= $1 AND ti.id=$2`,
		todoItemTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *TodoItemPostgres) Update(userId int, itemId int, inputUpdate todo.UpdateItem) error {

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if inputUpdate.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *inputUpdate.Title)
		argId++
	}

	if inputUpdate.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *inputUpdate.Description)
		argId++
	}

	if inputUpdate.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *inputUpdate.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		"UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id=li.item_id AND li.list_id=ul.list_id AND ul.user_id=$%d AND ti.id=$%d",
		todoItemTable,
		setQuery,
		listsItemsTable,
		usersListsTable,
		argId,
		argId+1,
	)

	args = append(args, userId, itemId)
	_, err := r.db.Exec(query, args...)

	return err
}
