-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS feeds (
    feed_id varchar(32) PRIMARY KEY,
    name text NOT NULL,
    url text NOT NULL UNIQUE,
    created_by varchar(32) NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    FOREIGN KEY (created_by) REFERENCES users (user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feeds;
-- +goose StatementEnd
