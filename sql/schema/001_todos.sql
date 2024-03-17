-- +goose Up

CREATE TABLE todos (
    id UUID UNIQUE NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    is_completed BOOLEAN NOT NULL
);

-- +goose Down

DROP TABLE todos;