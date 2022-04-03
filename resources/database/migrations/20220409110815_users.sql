-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         int NOT NULL AUTO_INCREMENT,
    name       text,
    email      text,
    age        int,
    birthday   datetime,
    created_at datetime,
    updated_at datetime,
    PRIMARY KEY (id),
    UNIQUE KEY `uk_email` (`email`) USING HASH
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
