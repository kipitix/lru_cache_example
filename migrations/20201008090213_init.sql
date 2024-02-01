-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    email text    primary key,
    name  text    not null,
    age   integer not null
);

INSERT INTO users (email, name, age) VALUES
    ('email@example.com', 'ivan', 25),
    ('email2@example.com', 'igor', 33),
    ('email3@example.com', 'oleg', 22)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
