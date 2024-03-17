-- name: AddTodo :one
INSERT INTO todos (
    id, title, description, is_completed
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetTodos :many
SELECT * FROM todos;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1;

-- name: UpdateTodo :exec
UPDATE todos 
    set title = $2,
    description = $2,
    is_completed = $3
WHERE id = $1;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;