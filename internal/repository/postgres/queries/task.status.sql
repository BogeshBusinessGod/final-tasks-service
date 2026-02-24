-- name: UpdateTaskStatus :one
UPDATE tasks
SET status = $2,
    updated_at = NOW()
WHERE id = $1
    RETURNING id, title, content, status, created_at, updated_at;
