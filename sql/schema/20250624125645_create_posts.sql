-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
    post_id varchar(32) PRIMARY KEY,
    title text,
    description text,
    url text NOT NULL UNIQUE,
    feed_id varchar(32) NOT NULL,
    published_at timestamp,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    FOREIGN KEY (feed_id) REFERENCES feeds (feed_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
