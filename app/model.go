package app

// ------------------stash-------------------
type Performers struct {
	ID           int    `json:"iD" db:"id"`
	Checksum     string `json:"checksum" db:"checksum"`
	Name         string `json:"name" db:"name"`
	Gender       string `json:"gender" db:"gender"`
	URL          string `json:"uRL" db:"url"`
	Twitter      string `json:"twitter" db:"twitter"`
	Instagram    string `json:"instagram" db:"instagram"`
	Birthdate    string `json:"birthdate" db:"birthdate"`
	Ethnicity    string `json:"ethnicity" db:"ethnicity"`
	Country      string `json:"country" db:"country"`
	EyeColor     string `json:"eyeColor" db:"eye_color"`
	Height       string `json:"height" db:"height"`
	Measurements string `json:"measurements" db:"measurements"`
	FakeTits     string `json:"fakeTits" db:"fake_tits"`
	CareerLength string `json:"careerLength" db:"career_length"`
	Tattoos      string `json:"tattoos" db:"tattoos"`
	Piercings    string `json:"piercings" db:"piercings"`
	Aliases      string `json:"aliases" db:"aliases"`
	CreatedAt    string `json:"createdAt" db:"created_at"`
	UpdatedAt    string `json:"updatedAt" db:"updated_at"`
	Details      string `json:"details" db:"details"`
	DeathDate    string `json:"deathDate" db:"death_date"`
	HairColor    string `json:"hairColor" db:"hair_color"`
	Weight       int    `json:"weight" db:"weight"`
	Rating       int    `json:"rating" db:"rating"`
}

type PerformersImage struct {
	PerformerID int    `json:"performerID" db:"performer_id"`
	Image       string `json:"image" db:"image"`
}
