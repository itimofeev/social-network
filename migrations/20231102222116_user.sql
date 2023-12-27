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

INSERT INTO users(id, user_id, password, first_name, second_name, birthdate, biography, interests, city)
VALUES ('a086c063-713e-4497-8a07-0b659a48eb41',
        'a086c063-713e-4497-8a07-0b659a48eb41',
        '2432612431302474705a35717761784d433642445863466b4e45734f2e573256464e55616173665271656379696e4f4376695258377150394b4f2f47',
        'Alex',
        'Ivanov',
        '2000-01-01',
        '',
        '',
        'Moscow'
       ),
       ('d06ff731-f291-4703-8010-a53c62be5d2b',
        'd06ff731-f291-4703-8010-a53c62be5d2b',
        '2432612431302474705a35717761784d433642445863466b4e45734f2e573256464e55616173665271656379696e4f4376695258377150394b4f2f47',
        'Maxim',
        'Tsvetkov',
        '2000-01-01',
        '',
        '',
        'Novosibirsk'
       );

CREATE INDEX users_first_name_second_name_index
    ON users (first_name text_pattern_ops, second_name text_pattern_ops);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users IF EXISTS;
-- +goose StatementEnd
