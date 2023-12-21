-- +goose Up
-- +goose StatementBegin
INSERT INTO users(full_name, email, password, role, balance) VALUES(
    'super admin', 'admin@mail.com', '$2a$10$jN0is23obwxkbHQhkSBqwuhxga0N5ZYzign6CRly4ozqryE.V/kva', 'admin', 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users
    WHERE email = 'admin@mail.com';
-- +goose StatementEnd
