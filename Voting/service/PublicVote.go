package service

import (
	"context"
	g "root/genproto"
	"root/storage/postgres"
)


type PublicVoteService struct{
	Public postgres.Storage
	g.UnimplementedPublicVoteServiceServer
}

func NewPublicVotoSerive(publicVote *postgres.Storage) *PublicVoteService{
	return &PublicVoteService{Public: *publicVote}
}

func (p *PublicVoteService) CreatePublicVote(ctx context.Context, req *g.CreatePublicVoteReq) (*g.PublicVote_Void, error){
	_, err := p.Public.PublicVote.Create(req)
	if err != nil{
		return nil, err
	}
	return &g.PublicVote_Void{}, nil
}

func (p *PublicVoteService) GetAllPublicVote(ctx context.Context, req *g.PublicVote_Void) (*g.GetPublicVoteResponse, error){
	res, err := p.Public.PublicVote.GetAll(req)
	if err != nil{
		return nil, err
	}
	return res, nil	
}

func (p *PublicVoteService) GetByIdPublicVote(ctx context.Context, req *g.GetPublicVoteRequest) (*g.PublicVote, error){
	res, err := p.Public.PublicVote.GetByPublicVoteId(req)
	if err != nil{
		return nil, err
	}
	return res, nil	
}

func (p *PublicVoteService) UpdatePublicVote(ctx context.Context, req *g.UpdatePublicVoteRequest) (*g.PublicVote_Void, error){
	_, err := p.Public.PublicVote.Update(req)
	if err != nil{
		return nil, err
	}	
	return &g.PublicVote_Void{}, nil
}

func (p *PublicVoteService) DeletePublicVote(ctx context.Context, req *g.DeletePublicVoteRequest) (*g.PublicVote_Void, error){
	_, err := p.Public.PublicVote.Delete(req)
	if err != nil{
		return nil, err
	}
	return &g.PublicVote_Void{}, nil
}

