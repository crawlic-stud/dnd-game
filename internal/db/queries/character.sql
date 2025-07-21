-- name: CreateCharacter :one 
INSERT INTO characters (name, class, level, avatar, metadata, user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetCharactersByUserID :many
SELECT id, name, class, level, avatar FROM characters WHERE user_id = $1;

-- name: GetCharacterByID :one
SELECT * FROM characters WHERE id = $1;

-- name: UpdateCharacter :one
UPDATE characters SET name = $2, class = $3, level = $4, avatar = $5, metadata = $6 WHERE id = $1 
RETURNING *;

-- name: DeleteCharacter :exec
DELETE FROM characters WHERE id = $1 RETURNING id;