package postgres

import (
	"database/sql"
	"fmt"
	g "root/genproto"

	"github.com/google/uuid"
)

type PublicVote struct {
	db *sql.DB
}

func NewPublicVote(db *sql.DB) *PublicVote {
	return &PublicVote{db: db}
}

func (pv *PublicVote) Create(req *g.CreatePublicVoteReq) (*g.PublicVote_Void, error) {
	req.Id = uuid.NewString()
	_, err := pv.db.Exec("insert into public_void(id, election_id, public_id) values($1, $2, $3)",
		req.Id, req.ElectionId, req.PublicId)
	if err != nil {
		return nil, fmt.Errorf("failed to create public vote: %v", err)
	}
	return &g.PublicVote_Void{}, nil

}

func (pv *PublicVote) GetByPublicVoteId(req *g.GetPublicVoteRequest) (*g.PublicVote, error) {
	var publicVote g.PublicVote
	err := pv.db.QueryRow("SELECT id, election_id, public_id FROM public_vote WHERE id = $1", req.Id).
		Scan(&publicVote, &publicVote.ElectionId.Id, &publicVote.PublicId)
	if err != nil {
		return nil, fmt.Errorf("failed to get public vote by public id: %v", err)
	}

	return &publicVote, nil
}

func (pv *PublicVote) Delete(req *g.DeletePublicVoteRequest) (*g.PublicVote_Void, error) {
	_, err := pv.db.Exec("delete from public_vote where id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete public vote: %v", err)
	}
	return &g.PublicVote_Void{}, nil
}

func (pv *PublicVote) Update(req *g.UpdatePublicVoteRequest) (*g.PublicVote_Void, error) {
	_, err := pv.db.Exec("update public_vote set election_id = $1, public_id = $2 where id = $3",
		req.PublicVote.ElectionId.Id, req.PublicVote.PublicId, req.PublicVote.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update public vote: %v", err)
	}
	return &g.PublicVote_Void{}, nil
}

func (pv *PublicVote) GetAll(req *g.PublicVote_Void) (*g.GetPublicVoteResponse, error) {
	rows, err := pv.db.Query("SELECT id, election_id, public_id FROM public_vote")
	if err != nil {
		return nil, fmt.Errorf("failed to get all public votes: %v", err)
	}
	defer rows.Close()

	var publicVotes []*g.PublicVote
	for rows.Next() {
		var publicVote g.PublicVote
		err := rows.Scan(&publicVote.Id, &publicVote.ElectionId.Id, &publicVote.PublicId)
		if err != nil {
			return nil, fmt.Errorf("failed to scan public vote: %v", err)
		}
		publicVotes = append(publicVotes, &publicVote)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return &g.GetPublicVoteResponse{PublicVote: publicVotes}, nil
}
