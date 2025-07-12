package api

import (
	"context"

	"github.com/ritarock/todo-backend/api/rest"
	"github.com/ritarock/todo-backend/repository"
	"github.com/ritarock/todo-backend/repository/query"
)

type handler struct {
	repo repository.TodoRepository
}

func NewHandler(repo repository.TodoRepository) *handler {
	return &handler{
		repo: repo,
	}
}

func (h *handler) TodosListTodos(ctx context.Context) ([]rest.Todo, error) {
	todos, err := h.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]rest.Todo, len(todos))
	for i, todo := range todos {
		result[i] = rest.Todo{
			ID:        todo.ID,
			Content:   todo.Content,
			Completed: todo.Completed,
		}
	}

	return result, nil
}

func (h *handler) TodosGetTodo(ctx context.Context, params rest.TodosGetTodoParams) (*rest.Todo, error) {
	todo, err := h.repo.GetByID(ctx, params.TodoID)
	if err != nil {
		return nil, err
	}

	return toRest(todo), nil
}

func (h *handler) TodosCreateTodo(ctx context.Context, req *rest.Todo) (*rest.Todo, error) {
	todo, err := h.repo.Create(ctx, &query.Todo{ID: req.ID, Content: req.Content, Completed: req.Completed})
	if err != nil {
		return nil, err
	}

	return toRest(todo), nil
}

func (h *handler) TodosUpdateTodo(ctx context.Context, req *rest.Todo, params rest.TodosUpdateTodoParams) (*rest.Todo, error) {
	if err := h.repo.Update(ctx, req.ID, &query.Todo{ID: req.ID, Content: req.Content, Completed: req.Completed}); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *handler) TodosDeleteTodo(ctx context.Context, params rest.TodosDeleteTodoParams) error {
	return h.repo.Delete(ctx, params.TodoID)
}

func toRest(todo *query.Todo) *rest.Todo {
	return &rest.Todo{
		ID:        todo.ID,
		Content:   todo.Content,
		Completed: todo.Completed,
	}
}
