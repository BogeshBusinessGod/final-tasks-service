-- name: CreateTask :one
INSERT INTO tasks (title, content, status)
VALUES ($1, $2, $3)
    RETURNING id, title, content, status, created_at, updated_at;

-- name: GetTask :one
SELECT id, title, content, status, created_at, updated_at
FROM tasks
WHERE id = $1
    LIMIT 1;

-- name: ListTasks :many
SELECT id, title, content, status, created_at, updated_at
FROM tasks
ORDER BY created_at DESC;



-- name: DoneTask :execrows
UPDATE tasks
SET status = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteTask :execrows
DELETE FROM tasks
WHERE id = $1;
