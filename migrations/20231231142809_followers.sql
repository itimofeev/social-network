-- +goose Up
-- +goose StatementBegin
CREATE TABLE followers
(
    user_id         uuid NOT NULL REFERENCES users (id),
    follows_user_id uuid NOT NULL REFERENCES users (id),
    PRIMARY KEY (user_id, follows_user_id),
    CHECK ( user_id <> followers.follows_user_id )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS followers;
-- +goose StatementEnd
