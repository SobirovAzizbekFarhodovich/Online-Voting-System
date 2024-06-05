package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage struct {
	db         *sql.DB
	Election   *ElectionRepo
	Candidate  *CandidateRepo
	Vote       *VoteRepo
	PublicVote *PublicVote
}

func ConnectDB() (*Storage, error) {
	psql := "user=postgres password=20005 dbname=voting1 sslmode=disable"
	db, err := sql.Open("postgres", psql)
	if err != nil {
		return nil, err
	}
	election := &ElectionRepo{db: db}
	candidate := &CandidateRepo{db: db}
	vote := &VoteRepo{db: db}
	publicVote := &PublicVote{db: db}

	return &Storage{db: db, Election: election, Candidate: candidate, Vote: vote, PublicVote: publicVote}, nil
}
