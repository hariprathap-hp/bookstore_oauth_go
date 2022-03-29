package db

import (
	"test3/hariprathap-hp/bookstore_oauth_go/src/domain/access_token"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
)

type DBRepository interface {
	GetbyID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbrepo struct{}

func NewRepository() DBRepository {
	return &dbrepo{}
}

func (dr *dbrepo) GetbyID(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("Yet to be implemented")
}
