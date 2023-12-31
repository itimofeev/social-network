-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts
(
    id         uuid NOT NULL PRIMARY KEY,
    user_id    uuid NOT NULL REFERENCES users (id),
    text       text NOT NULL,
    created_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd
