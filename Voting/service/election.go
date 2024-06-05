package service

import (
	"context"
	g "root/genproto"
	"root/storage/postgres"
)

type ElectionService struct {
	Election postgres.Storage
	g.UnimplementedElectionServiceServer
}

func NewElectionService(election *postgres.Storage) *ElectionService {
	return &ElectionService{
		Election: *election,
	}	
}

func (e *ElectionService) CreateElection(ctx context.Context, req *g.CreateElectionRequest) (*g.Election_Void, error) {
	_, err := e.Election.Election.Create(req)
	if err != nil {
		return nil, err
	}
	return &g.Election_Void{}, nil
}

func (e *ElectionService) GetAllElections(ctx context.Context, req *g.Election_Void) (*g.GetAllElectionsResponse, error){
	res, err := e.Election.Election.GetAll(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *ElectionService) GetByIdElection(ctx context.Context, req *g.GetElectionRequest) (*g.GetElectionResponse, error){
	res, err := e.Election.Election.GetByID(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *ElectionService) UpdateElection(ctx context.Context, req *g.UpdateElectionRequest) (*g.Election_Void, error){
	_, err := e.Election.Election.Update(req)
	if err != nil {
		return nil, err
	}
	return &g.Election_Void{}, nil
}

func (e *ElectionService) DeleteElection(ctx context.Context, req *g.DeleteElectionRequest) (*g.Election_Void, error){
	_, err := e.Election.Election.Delete(req)
	if err != nil {
		return nil, err
	}
	return &g.Election_Void{}, nil
}