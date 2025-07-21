package models

import (
	"dnd-game/internal/util/validation"
	"time"

	"github.com/google/uuid"
)

type CharacterResponse struct {
	ID        uuid.UUID         `json:"id"`
	Name      string            `json:"name"`
	Class     string            `json:"class"`
	Level     int16             `json:"level"`
	Avatar    *string           `json:"avatar"`
	Metadata  CharacterMetadata `json:"metadata"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

type CharacterCreate struct {
	Name     string            `json:"name"`
	Class    string            `json:"class"`
	Level    int16             `json:"level"`
	Avatar   *string           `json:"avatar"`
	Metadata CharacterMetadata `json:"metadata"`
}

func (c CharacterCreate) Validate() error {
	return validation.NewValidator(c).
		Add(c.Name != "", "name must not be empty").
		Add(c.Class != "", "class must not be empty").
		Add(c.Level > 0 && c.Level <= 20, "level must be between 1 and 20").
		CheckError(c.Metadata.Validate()).
		Validate()
}

// CharacterMetadata represents a D&D character data
type CharacterMetadata struct {
	Race       string `json:"race"`
	Background string `json:"background"`
	Stats      Stats  `json:"stats"`
	Skills     Skills `json:"skills"`
	HitPoints  int    `json:"hitPoints"`
	ArmorClass int    `json:"armorClass"`
}

func (c CharacterMetadata) Validate() error {
	return validation.NewValidator(c).
		Add(c.Race != "", "race must not be empty").
		Add(c.Background != "", "background must not be empty").
		Add(c.HitPoints > 0, "hitPoints must be greater than 0").
		Add(c.ArmorClass >= 0, "armorClass must be greater than or equal to 0").
		CheckError(c.Stats.Validate()).
		CheckError(c.Skills.Validate()).
		Validate()
}

// Ыtats represents the character's ability scores
type Stats struct {
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`
}

func (s Stats) Validate() error {
	return validation.NewValidator(s).
		Add(s.Strength > 0 && s.Strength <= 18, "strength must be between 1 and 18").
		Add(s.Dexterity > 0 && s.Dexterity <= 18, "dexterity must be between 1 and 18").
		Add(s.Constitution > 0 && s.Constitution <= 18, "constitution must be between 1 and 18").
		Add(s.Intelligence > 0 && s.Intelligence <= 18, "intelligence must be between 1 and 18").
		Add(s.Wisdom > 0 && s.Wisdom <= 18, "wisdom must be between 1 and 18").
		Add(s.Charisma > 0 && s.Charisma <= 18, "charisma must be between 1 and 18").
		Validate()
}

// Skills represents the character's Skills
type Skills struct {
	Acrobatics     int `json:"acrobatics"`
	AnimalHandling int `json:"animalHandling"`
	Athletics      int `json:"athletics"`
	Deception      int `json:"deception"`
	History        int `json:"history"`
	Insight        int `json:"insight"`
	Investigation  int `json:"investigation"`
	Medicine       int `json:"medicine"`
	Nature         int `json:"nature"`
	Perception     int `json:"perception"`
	Stealth        int `json:"stealth"`
	Survival       int `json:"survival"`
}

func (s Skills) Validate() error {
	return validation.NewValidator(s).
		Add(s.Acrobatics > 0 && s.Acrobatics <= 18, "acrobatics must be between 1 and 18").
		Add(s.AnimalHandling > 0 && s.AnimalHandling <= 18, "animalHandling must be between 1 and 18").
		Add(s.Athletics > 0 && s.Athletics <= 18, "athletics must be between 1 and 18").
		Add(s.Deception > 0 && s.Deception <= 18, "deception must be between 1 and 18").
		Add(s.History > 0 && s.History <= 18, "history must be between 1 and 18").
		Add(s.Insight > 0 && s.Insight <= 18, "insight must be between 1 and 18").
		Add(s.Investigation > 0 && s.Investigation <= 18, "investigation must be between 1 and 18").
		Add(s.Medicine > 0 && s.Medicine <= 18, "medicine must be between 1 and 18").
		Add(s.Nature > 0 && s.Nature <= 18, "nature must be between 1 and 18").
		Add(s.Perception > 0 && s.Perception <= 18, "perception must be between 1 and 18").
		Add(s.Stealth > 0 && s.Stealth <= 18, "stealth must be between 1 and 18").
		Add(s.Survival > 0 && s.Survival <= 18, "survival must be between 1 and 18").
		Validate()
}
