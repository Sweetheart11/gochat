-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF EXISTS users (
  
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
