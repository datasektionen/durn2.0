package database

import (
	"time"

	// "github.com/google/uuid"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Election struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey"`
	Name       string    `gorm:"not null"`
	published  bool      `gorm:"not null"`
	finalized  bool      `gorm:"not null"`
	openTime   time.Time
	closeTime  time.Time
	Candidates []Candidate `gorm:"many2many:candidate_in_election"`
	Votes      []Vote
}

type Valid_Voter struct {
	gorm.Model
	email string `gorm:"primaryKey"`
}

type Candidate struct {
	gorm.Model
	ID           uuid.UUID `gorm:"primaryKey"`
	Name         string    `gorm:"not null"`
	presentation string    `gorm:"not null"`
}

type Casted_Vote struct {
	gorm.Model
	voter    Valid_Voter `gorm:"primaryKey"`
	election Election    `gorm:"primaryKey"`
}

type Vote struct {
	gorm.Model
	hash     string    `gorm:"primaryKey"`
	election Election  `gorm:"not null"`
	isBlank  bool      `gorm:"not null"`
	voteTime time.Time `gorm:"not null"`
}

// import (
// 	"time"

// 	"github.com/google/uuid"
// 	// _ "go.mongodb.org/mongo-driver/bson"
// )

// type Election struct {
// 	id        uuid.UUID `bson:id`
// 	name      string    `bson:name`
// 	published bool      `bson:published`
// 	openTime  time.Time `bson:`
// 	closeTime time.Time
// }

// type ValidVoter struct {
// 	mail string
// }

// type Candidate struct {
// 	id           uuid.UUID
// 	name         string
// 	presentation string
// }

// type CastedVote struct {
// 	mail       string
// 	electionId uuid.UUID
// }

// type Vote struct {
// 	hash       string
// 	electionId uuid.UUID
// 	isBlank    bool
// }

// type CandidateInElection struct {
// 	electionId  uuid.UUID
// 	CandidateId uuid.UUID
// }

// type VoteEvent struct {
// 	voteId   string
// 	voteTime time.Time
// }
