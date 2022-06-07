-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         serial primary key,
    name       text,
    email      text,
    age        int,
    birthday   timestamp,
    created_at timestamp,
    updated_at timestamp,
    constraint uk_email unique (email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
