-- +goose Up
-- +goose StatementBegin
CREATE TABLE feed_follows (
    feed_follow_id varchar(32) PRIMARY KEY,
    feed_id varchar(32) NOT NULL,
    follower_id varchar(32) NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    FOREIGN KEY (feed_id) REFERENCES feeds (feed_id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES users (user_id) ON DELETE CASCADE,
    CONSTRAINT uq_feed_follows UNIQUE (feed_id, follower_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feed_follows;
-- +goose StatementEnd
