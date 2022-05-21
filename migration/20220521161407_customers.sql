-- +goose Up
-- +goose StatementBegin
CREATE TABLE customers
(
    id         INT primary key auto_increment,
    username   VARCHAR(100) unique,
    email      VARCHAR(255),
    password   VARCHAR(100),
    gender     VARCHAR(50),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE = InnoDB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS customers;
-- +goose StatementEnd
