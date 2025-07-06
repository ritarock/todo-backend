-- name: CreateTodo :execresult
INSERT INTO todos (
    content, completed
) VALUES (
    ?, ?
);

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = ?;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id;

-- name: UpdateTodo :exec
UPDATE todos
SET content = ?, completed = ?
WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;
