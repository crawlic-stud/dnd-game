-- +goose Up
-- +goose StatementBegin
CREATE TABLE games (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    name TEXT NOT NULL,
    max_players SMALLINT NOT NULL DEFAULT 4
);

-- Scenes that are used during a game
CREATE TABLE game_scenes (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    game_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    name TEXT NOT NULL,
    map_image TEXT NOT NULL,
    width SMALLINT NOT NULL,
    height SMALLINT NOT NULL,
    CONSTRAINT fk_game_scenes_game FOREIGN KEY (game_id) REFERENCES games (id)
);

-- Objects blueprints that are used during a scene
CREATE TABLE game_objects (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    object_type SMALLINT NOT NULL,
    texture TEXT NOT NULL
);

-- Objects that are used during a scene
CREATE TABLE scenes_objects (
    scene_id UUID NOT NULL,
    object_id UUID NOT NULL,
    pos_x SMALLINT NOT NULL,
    pos_y SMALLINT NOT NULL,
    CONSTRAINT fk_scenes_objects_scene FOREIGN KEY (scene_id) REFERENCES game_scenes (id),
    CONSTRAINT fk_scenes_objects_object FOREIGN KEY (object_id) REFERENCES game_objects (id)
);

-- Users that participate in a game
CREATE TABLE game_players (
    user_id UUID NOT NULL,
    game_id UUID NOT NULL,
    role TEXT NOT NULL,
    CONSTRAINT fk_game_players_user FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_game_players_game FOREIGN KEY (game_id) REFERENCES games (id),
    CONSTRAINT game_players_role_constraint CHECK (role = 'HOST' OR role = 'PLAYER')
);

-- Characters that are used during a game
CREATE TABLE game_characters (
    character_id UUID NOT NULL,
    game_id UUID NOT NULL,
    CONSTRAINT fk_game_characters_character FOREIGN KEY (character_id) REFERENCES characters (id),
    CONSTRAINT fk_game_characters_game FOREIGN KEY (game_id) REFERENCES games (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
