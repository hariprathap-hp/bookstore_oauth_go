package token_service

import (
	"test3/hariprathap-hp/bookstore_oauth_go/src/domain/access_token"
	"test3/hariprathap-hp/bookstore_oauth_go/src/repository/db"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
)

type Service interface {
	GetbyID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
}

type service struct {
	repository db.DBRepository
}

func NewService(repo db.DBRepository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetbyID(id string) (*access_token.AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetbyID(id)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	err := s.repository.Create(at)
	if err != nil {
		return err
	}
	return nil
}
