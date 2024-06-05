package service

import (
	"context"
	g "root/genproto"
	"root/storage/postgres"
)


type VoteService struct{
	Vote postgres.Storage
	g.UnimplementedVoteServiceServer
}

func NewVoteService(Vote *postgres.Storage) *VoteService {
	return &VoteService{Vote: *Vote}
}

func (v *VoteService) CreateVote(ctx context.Context, req *g.Vote) (*g.Vote_Void, error){
	_, err := v.Vote.Vote.Create(req)
	if err != nil {
		return nil, err
	}
	return &g.Vote_Void{}, nil
}

func (v *VoteService) GetAllVote(ctx context.Context, req *g.Vote_Void) (*g.GetAllVotes, error){
	votes, err := v.Vote.Vote.GetAll(req)
	if err != nil {
		return nil, err
	}
	return votes, nil
}


func (v *VoteService) GetByIdVote(ctx context.Context, req *g.GetVoteRequest) (*g.Vote, error){
	vote, err := v.Vote.Vote.GetByID(req)
	if err != nil {
		return nil, err
	}
	return vote, nil
}

func (v *VoteService) DeleteVote(ctx context.Context, req *g.DeleteVoteRequest) (*g.Vote_Void, error){
	_, err := v.Vote.Vote.Delete(req)
	if err != nil {
		return nil, err
	}
	return &g.Vote_Void{}, nil
}

func(v *VoteService) UpdateVote(ctx context.Context, req *g.UpdateVoteRequest) (*g.Vote_Void, error){
	_, err := v.Vote.Vote.Update(req)
	if err != nil {
		return nil, err
	}
	return &g.Vote_Void{}, nil
}
