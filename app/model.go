package app

import "database/sql"

// ------------------stash-------------------
type Performers struct {
	ID int `json:"iD" db:"id"`
	// Checksum     string `json:"checksum" db:"checksum"`
	Name         string         `json:"name" db:"name"`
	Gender       sql.NullString `json:"gender" db:"gender"`
	URL          sql.NullString `json:"uRL" db:"url"`
	Twitter      sql.NullString `json:"twitter" db:"twitter"`
	Instagram    sql.NullString `json:"instagram" db:"instagram"`
	Birthdate    sql.NullString `json:"birthdate" db:"birthdate"`
	Ethnicity    sql.NullString `json:"ethnicity" db:"ethnicity"`
	Country      sql.NullString `json:"country" db:"country"`
	EyeColor     sql.NullString `json:"eyeColor" db:"eye_color"`
	Height       sql.NullInt32  `json:"height" db:"height"`
	Measurements sql.NullString `json:"measurements" db:"measurements"`
	FakeTits     sql.NullString `json:"fakeTits" db:"fake_tits"`
	CareerLength sql.NullString `json:"careerLength" db:"career_length"`
	Tattoos      sql.NullString `json:"tattoos" db:"tattoos"`
	Piercings    sql.NullString `json:"piercings" db:"piercings"`
	Aliases      sql.NullString `json:"aliases" db:"aliases"`
	CreatedAt    sql.NullString `json:"createdAt" db:"created_at"`
	UpdatedAt    sql.NullString `json:"updatedAt" db:"updated_at"`
	Details      sql.NullString `json:"details" db:"details"`
	DeathDate    sql.NullString `json:"deathDate" db:"death_date"`
	HairColor    sql.NullString `json:"hairColor" db:"hair_color"`
	Weight       sql.NullInt32  `json:"weight" db:"weight"`
	Rating       sql.NullInt32  `json:"rating" db:"rating"`
}

type PerformersImage struct {
	PerformerID int    `json:"performerID" db:"performer_id"`
	Image       string `json:"image" db:"image"`
}

// ---------------------minnano-av.com----------------------
