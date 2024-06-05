package service

import (
	"context"
	g "root/genproto"
	"root/storage/postgres"
)

type CandidateService struct {
	Candidate postgres.Storage
	g.UnimplementedCandidateServiceServer
}

func NewCandidateService(candidate *postgres.Storage) *CandidateService {
	return &CandidateService{Candidate: *candidate}
}

func (s *CandidateService) CreateCandidate(ctx context.Context, req *g.CreateCandidateRequest) (*g.Candidate_Void, error) {
	_, err := s.Candidate.Candidate.Create(req)
	if err != nil {
		return nil, err
	}
	return &g.Candidate_Void{}, nil
}

func (s *CandidateService) GetAllCandidates(ctx context.Context, req *g.Candidate_Void) (*g.GetAllCandidatesResponse, error) {
	res, err := s.Candidate.Candidate.GetAll(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CandidateService) GetByIdCandidate(ctx context.Context, req *g.GetByIdCandidateRequest) (*g.GetByIdCandidateResponse, error){
	res, err := s.Candidate.Candidate.GetByID(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (s *CandidateService) UpdateCandidate(ctx context.Context, req *g.UpdateCandidateRequest) (*g.Candidate_Void, error) {
	_, err := s.Candidate.Candidate.Update(req)
	if err != nil {
		return nil, err
	}
	return &g.Candidate_Void{}, nil	
}

func (s *CandidateService) DeleteCandidate(ctx context.Context, req *g.DeleteCandidateRequest) (*g.Candidate_Void, error) {
	_, err := s.Candidate.Candidate.Delete(req)
	if err != nil {
		return nil, err
	}
	return &g.Candidate_Void{}, nil
}