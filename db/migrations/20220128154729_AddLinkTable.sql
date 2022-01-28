
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE link
(
    id           uuid    PRIMARY KEY DEFAULT uuid_generate_v4(),
    short        text NOT NULL,
    retrived     int NOT NULL DEFAULT 0,
    title        text NOT NULL,
    created_at   timestamp with time zone NOT NULL DEFAULT now(),
    updated_at   timestamp with time zone NOT NULL DEFAULT now()
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE link CASCADE;
