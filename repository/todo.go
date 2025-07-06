package repository

import (
	"context"
	"database/sql"

	"github.com/ritarock/todo-backend/repository/query"
)

type todoRepository struct {
	query query.Queries
}

func NewTodoRepository(db *sql.DB) *todoRepository {
	q := query.New(db)
	return &todoRepository{query: *q}
}

func (t *todoRepository) Create(ctx context.Context, todo *query.Todo) (*query.Todo, error) {
	result, err := t.query.CreateTodo(ctx, query.CreateTodoParams{
		Content:   todo.Content,
		Completed: todo.Completed,
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	todo.ID = id

	return todo, nil
}

func (t *todoRepository) GetByID(ctx context.Context, id int64) (*query.Todo, error) {
	result, err := t.query.GetTodo(ctx, id)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (t *todoRepository) GetAll(ctx context.Context) ([]query.Todo, error) {
	result, err := t.query.ListTodos(ctx)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (t *todoRepository) Update(ctx context.Context, id int64, todo *query.Todo) error {
	return t.query.UpdateTodo(ctx, query.UpdateTodoParams{
		Content:   todo.Content,
		Completed: todo.Completed,
		ID:        id,
	})
}

func (t *todoRepository) Delete(ctx context.Context, id int64) error {
	return t.query.DeleteTodo(ctx, id)
}
