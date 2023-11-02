-- +goose Up
-- +goose StatementBegin
SELECT 'init up migration';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'init down migration';
-- +goose StatementEnd
