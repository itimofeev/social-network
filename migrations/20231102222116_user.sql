-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id          uuid PRIMARY KEY,

    user_id     varchar(256) NOT NULL UNIQUE,
    password    varchar(256) NOT NULL,

    first_name  varchar(256) NOT NULL,
    second_name varchar(256) NOT NULL,

    birthdate   timestamptz  NOT NULL,

    biography   text         NOT NULL,

    interests   text         NOT NULL,
    city        varchar(256) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users IF EXISTS;
-- +goose StatementEnd
