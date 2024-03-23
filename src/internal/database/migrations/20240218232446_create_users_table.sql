-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id           VARCHAR(40) PRIMARY KEY NOT NULL,
    username    VARCHAR(255) NOT NULL,
    email       VARCHAR(255) NOT NULL,
    is_banned  BOOLEAN NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMP(6) NOT NULL,
    updated_at TIMESTAMP(6)
);

CREATE UNIQUE INDEX uidx_users_username ON users (username);
CREATE UNIQUE INDEX uidx_users_email ON users (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX uidx_users_username;
DROP INDEX uidx_users_email;
DROP TABLE users;
-- +goose StatementEnd
