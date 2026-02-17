
-- ======================
-- Таблица пользователей
-- ======================
CREATE TABLE users (
                       id         BIGSERIAL PRIMARY KEY,
                       username   TEXT UNIQUE NOT NULL,
                       password   BYTEA NOT NULL,
                       created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
                       updated_at TIMESTAMPTZ,
                       deleted_at TIMESTAMPTZ
);

-- ======================
-- Таблица задач
-- ======================
--content - содержание заметки, сами записи.
-- условно title - задача(купить машину)
-- content - последовательность действий для решения задачи(взять налик, поехать в автосалон и т.д)

CREATE TABLE tasks (
                       id         BIGSERIAL PRIMARY KEY,
                       user_id    BIGINT REFERENCES users(id) ON DELETE CASCADE,
                       title      TEXT NOT NULL,
                       content    TEXT NOT NULL DEFAULT '',
                       status     TEXT NOT NULL DEFAULT 'new'
                           CHECK (status IN ('new', 'in_progress', 'done', 'error')),
                       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMPTZ
);

COMMENT ON COLUMN tasks.content IS 'Task description (plain text). Empty string means no description.';
COMMENT ON COLUMN tasks.status  IS 'Task lifecycle status.';
