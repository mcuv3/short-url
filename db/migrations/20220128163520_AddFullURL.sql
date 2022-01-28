-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE link ADD full_url text NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE link DROP ;
