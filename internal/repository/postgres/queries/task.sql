-- name: CreateTask :one
INSERT INTO tasks (title, content, status, done, created_at)
VALUES ($1, $2, $3, $4, NOW())
    RETURNING *;

-- name: GetTaskByID :one
SELECT * FROM tasks WHERE id = $1 LIMIT 1;

-- name: ListTasks :many
SELECT * FROM tasks ORDER BY created_at DESC;

-- name: DoneTask :exec
UPDATE tasks
SET done = TRUE,
    status = 'done'
WHERE id = $1;

-- name: DeleteTask :exec
UPDATE tasks
SET deleted_at = NOW()
WHERE id = $1;
