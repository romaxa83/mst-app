package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
	"strings"
	"time"
)

type TodoListRepo struct {
	db *sqlx.DB
}

func NewTodoListRepo(db *sqlx.DB) *TodoListRepo {
	return &TodoListRepo{db: db}
}

func (r *TodoListRepo) Create(ctx context.Context, userId int, list domains.TodoList) (int, error) {
	// применяем транзакцию, объявляем начало транзакции
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf(`INSERT INTO %s 
											(title, description, created_at, updated_at)
										VALUES ($1, $2, $3, $4) RETURNING id`, todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description, list.CreatedAt, list.UpdatedAt)
	if err := row.Scan(&id); err != nil {
		// если что-то не так, откатываем транзакцию
		tx.Rollback()
		return 0, err
	}

	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1,$2)", usersListsTable)
	_, err = tx.Exec(createUserListQuery, userId, id)
	if err != nil {
		// если что-то не так, откатываем транзакцию
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListRepo) GetAll(userId int) ([]domains.TodoList, error) {
	var lists []domains.TodoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description, tl.created_at, tl.updated_at 
									FROM %s tl 
								INNER JOIN %s ul on tl.id = ul.list_id 
									WHERE ul.user_id = $1 AND tl.deleted_at IS NULL`,
		todoListsTable, usersListsTable)

	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListRepo) GetById(userId, listId int) (domains.TodoList, error) {
	var list domains.TodoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
						INNER JOIN %s ul on tl.id = ul.list_id 
						WHERE ul.user_id = $1 AND ul.list_id = $2 AND tl.deleted_at IS NULL`,
		todoListsTable, usersListsTable)

	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *TodoListRepo) Delete(userId, listId int) error {
	query := fmt.Sprintf(`UPDATE %s tl 
								SET deleted_at = $1, updated_at = $1
								FROM %s ul 
								WHERE tl.id = ul.list_id AND ul.user_id = $2 AND ul.list_id = $3`,
		todoListsTable, usersListsTable)

	_, err := r.db.Exec(query, time.Now(), userId, listId)

	return err
}

func (r *TodoListRepo) Update(userId, listId int, input domains.UpdateTodoListInput) error {

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
	setValues = append(setValues, fmt.Sprintf("updated_at=$%d", argId))
	args = append(args, time.Now())
	argId++

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d`,
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	_, err := r.db.Exec(query, args...)

	return err
}
