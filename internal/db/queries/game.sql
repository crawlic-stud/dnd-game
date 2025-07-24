-- name: GetGame :one
SELECT * FROM games WHERE games.id = $1;

-- name: GetGameScenes :many
SELECT * FROM game_scenes WHERE game_id = $1;

-- name: GetGameCharacters :many
SELECT characters.* FROM characters 
JOIN game_characters ON characters.id IN (
    SELECT character_id FROM game_characters WHERE game_characters.game_id = $1
);

-- name: GetGameUsers :many
SELECT id, username FROM users
JOIN game_players ON users.id IN (
    SELECT user_id FROM game_players WHERE game_players.game_id = $1
);

-- name: GetSceneObjects :many
SELECT * FROM scenes_objects WHERE scene_id = $1;

-- name: CreateGame :one
INSERT INTO games (name, max_players) VALUES ($1, $2) returning id;

-- name: CreateGameScene :one
INSERT INTO game_scenes (game_id, name, map_image, width, height) VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: CreateGameObject :exec
INSERT INTO game_objects (object_type, texture) VALUES ($1, $2);

-- name: AddObjectToScene :exec
INSERT INTO scenes_objects (scene_id, object_id, pos_x, pos_y) VALUES ($1, $2, $3, $4);

-- name: AddCharacterToGame :exec
INSERT INTO game_characters (game_id, character_id) VALUES ($1, $2);

-- name: AddUserToGame :exec
INSERT INTO game_players (game_id, user_id, role) VALUES ($1, $2, $3);