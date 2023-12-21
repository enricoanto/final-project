-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transaction_histories (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products (id),
    user_id INTEGER NOT NULL REFERENCES users (id),
    quantity INTEGER NOT NULL,
    total_price INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transaction_histories;
-- +goose StatementEnd
