
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
    id bigint(20) NOT NULL AUTO_INCREMENT,
    name varchar(255)  COLLATE utf8mb4_unicode_ci NOT NULL,
    email varchar(255)  COLLATE utf8mb4_unicode_ci NOT NULL,
    PRIMARY KEY(id),
    UNIQUE INDEX unique_index_on_users_email(email)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
