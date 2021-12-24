package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
	"strings"
	"time"
)

type TodoItemRepo struct {
	db *sqlx.DB
}

func NewTodoItemRepo(db *sqlx.DB) *TodoItemRepo {
	return &TodoItemRepo{db: db}
}

func (r *TodoItemRepo) Create(listId int, item domains.TodoItem) (int, error) {
	// применяем транзакцию, объявляем начало транзакции
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf(`INSERT INTO %s 
											(title, description, created_at, updated_at) 
											VALUES ($1, $2, $3, $4) RETURNING id`,
		todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description, item.CreatedAt, item.UpdatedAt)
	if err := row.Scan(&itemId); err != nil {
		// если что-то не так, откатываем транзакцию
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1,$2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		// если что-то не так, откатываем транзакцию
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemRepo) GetAll(userId, listId int) ([]domains.TodoItem, error) {
	var items []domains.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
							INNER JOIN %s li on li.item_id = ti.id
							INNER JOIN %s ul on ul.list_id = li.list_id
						WHERE li.list_id = $1 AND ul.user_id = $2 AND ti.deleted_at IS NULL`,
		todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemRepo) GetById(userId, itemId int) (domains.TodoItem, error) {
	var item domains.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
							INNER JOIN %s li on li.item_id = ti.id
							INNER JOIN %s ul on ul.list_id = li.list_id 
						WHERE ti.id = $1 AND ul.user_id = $2 AND ti.deleted_at IS NULL`,
		todoItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemRepo) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`UPDATE %s ti
								SET deleted_at = $1, updated_at = $1 
								FROM %s li, %s ul
								WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $2 AND ti.id = $3`,
		todoItemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, time.Now(), userId, itemId)
	return err
}

func (r *TodoItemRepo) Update(userId, itemId int, input domains.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}
	setValues = append(setValues, fmt.Sprintf("updated_at=$%d", argId))
	args = append(args, time.Now())
	argId++

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
						WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)
	return err
}
