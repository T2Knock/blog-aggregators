-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    user_id varchar(32) PRIMARY KEY,
    name text NOT NULL UNIQUE,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
