package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Election struct {
	ID         uuid.UUID    `gorm:"primaryKey"`
	Name       string       `gorm:"not null"`
	Published  bool         `gorm:"not null"`
	Finalized  bool         `gorm:"not null"`
	OpenTime   sql.NullTime ``
	CloseTime  sql.NullTime ``
	Candidates []Candidate  `gorm:"many2many:candidate_in_election"`
	Votes      []Vote       ``
}

type Valid_Voter struct {
	Email string `gorm:"primaryKey"`
}

type Candidate struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	Name         string    `gorm:"not null"`
	Presentation string    `gorm:"not null"`
}

type Casted_Vote struct {
	VoterID    uuid.UUID   `gorm:"primaryKey"`
	ElectionID uuid.UUID   `gorm:"primaryKey"`
	Voter      Valid_Voter `gorm:"foreignKey:VoterID"`
	Election   Election    `gorm:"foreignKey:ElectionID"`
}

type Vote struct {
	Hash       string    `gorm:"primaryKey"`
	IsBlank    bool      `gorm:"not null"`
	VoteTime   time.Time `gorm:"not null"`
	ElectionID uuid.UUID `gorm:"not null"`
	Election   Election  `gorm:"foreignKey:ElectionID"`
}
