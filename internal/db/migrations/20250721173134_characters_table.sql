-- +goose Up
-- +goose StatementBegin
CREATE TABLE characters (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    name TEXT NOT NULL,
    class TEXT NOT NULL,
    level SMALLINT NOT NULL,
    avatar TEXT,
    user_id UUID NOT NULL,
    metadata JSONB NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE characters;
-- +goose StatementEnd
