package postgres

import (
	"database/sql"
	"fmt"
	g "root/genproto"

	"github.com/google/uuid"
)

type VoteRepo struct{
	db *sql.DB
}

func NewVoteRepo(db *sql.DB) *VoteRepo {
	return &VoteRepo{
		db: db,
	}
}


func (v *VoteRepo) Create(req *g.Vote) (*g.Vote_Void, error){
	req.Id = uuid.NewString()
	_, err := v.db.Exec("insert into votes(id, candidate_id), values($1, $2)",
	req.Id, req.CandidateId)
	if err != nil{
		return nil, fmt.Errorf("failed to create vote: %v", err)
	}
	return &g.Vote_Void{}, nil
}


func (v *VoteRepo) GetByID(req *g.GetVoteRequest) (*g.Vote, error){
	var vote g.Vote
	err := v.db.QueryRow("SELECT id, candidate_id FROM votes WHERE id = $1", req.Id).
	Scan(&vote.Id, &vote.CandidateId)
	if err != nil{
		return nil, fmt.Errorf("failed to get vote by id: %v", err)
	}
	return &vote, nil
}

func (v *VoteRepo) Delete(req *g.DeleteVoteRequest) (*g.Vote_Void, error){
	_, err := v.db.Exec("DELETE FROM votes WHERE id = $1", req.Id)
	if err != nil{
		return nil, fmt.Errorf("failed to delete vote: %v", err)
	}
	return &g.Vote_Void{}, nil
}

func (v *VoteRepo) Update(req *g.UpdateVoteRequest) (*g.Vote_Void, error){
	_, err := v.db.Exec("UPDATE votes SET candidate_id = $1 WHERE id = $2",
	req.Vote.CandidateId, req.Vote.Id)
	if err != nil{
		return nil, fmt.Errorf("failed to update vote: %v", err)
	}
	return &g.Vote_Void{}, nil
}


func (v *VoteRepo) GetAll(req *g.Vote_Void) (*g.GetAllVotes, error){
	rows, err := v.db.Query("SELECT id, candidate_id FROM votes")
	if err != nil{
		return nil, fmt.Errorf("failed to get all votes: %v", err)
	}
	defer rows.Close()
	var votes []*g.GetAll
	for rows.Next(){
		var vote g.GetAll
		err := rows.Scan(&vote.Id, &vote.CandidateId)
		if err != nil{
			return nil, fmt.Errorf("failed to scan vote: %v", err)
		}
		votes = append(votes, &vote)
	}
	return &g.GetAllVotes{Votes: votes}, nil
}
