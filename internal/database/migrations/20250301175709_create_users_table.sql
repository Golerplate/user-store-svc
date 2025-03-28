-- +goose Up

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION set_updated_at_column() RETURNS TRIGGER AS $$
  BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
  END;
$$ language 'plpgsql';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE users (
    id              VARCHAR(40) PRIMARY KEY NOT NULL,
    first_name      VARCHAR(255) NOT NULL,
    last_name       VARCHAR(255) NOT NULL,
    nick_name       VARCHAR(255) NOT NULL,
    password        VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL,
    country         VARCHAR(255) NOT NULL,
    created_at      TIMESTAMP(6) NOT NULL,
    updated_at      TIMESTAMP(6) NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX uidx_users_nickname   ON users (nick_name);
CREATE UNIQUE INDEX uidx_users_email      ON users (email);
CREATE INDEX        idx_users_country     ON users (country);
CREATE TRIGGER      tgr_users_updated_at  BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION set_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX    uidx_users_nickname;
DROP INDEX    uidx_users_email;
DROP INDEX    idx_users_country;

DROP TRIGGER  tgr_users_updated_at ON users;
DROP FUNCTION set_updated_at_column;

DROP TABLE    users;
-- +goose StatementEnd
