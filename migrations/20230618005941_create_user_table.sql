-- +goose Up

-- +goose StatementBegin

SELECT 'up SQL query';

CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        first_name TEXT NOT NULL,
        last_name TEXT NOT NULL,
        email TEXT NOT NULL
    );

-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin

SELECT 'down SQL query';

DROP TABLE IF EXISTS users;

-- +goose StatementEnd