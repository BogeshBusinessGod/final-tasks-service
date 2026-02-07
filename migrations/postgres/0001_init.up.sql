
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
CREATE TABLE tasks (
                       id         BIGSERIAL PRIMARY KEY,
                       user_id    BIGINT REFERENCES users(id) ON DELETE CASCADE,
                       title      TEXT NOT NULL,
                       content    TEXT,
                       status     TEXT NOT NULL DEFAULT 'new'
                           CHECK (status IN ('new', 'done', 'deleted')),
                       done       BOOLEAN NOT NULL DEFAULT false,
                       created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
                       updated_at TIMESTAMPTZ,
                       deleted_at TIMESTAMPTZ
);

-- ======================
-- Функция + триггер для done
-- ======================
CREATE OR REPLACE FUNCTION sync_task_done()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status = 'done' THEN
        NEW.done := true;
ELSE
        NEW.done := false;
END IF;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_sync_task_done ON tasks;

CREATE TRIGGER trg_sync_task_done
    BEFORE INSERT OR UPDATE ON tasks
                         FOR EACH ROW
                         EXECUTE FUNCTION sync_task_done();
