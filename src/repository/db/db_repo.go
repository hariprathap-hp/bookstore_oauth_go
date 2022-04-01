package db

import (
	"test3/hariprathap-hp/bookstore_oauth_go/src/client/cassandra"
	"test3/hariprathap-hp/bookstore_oauth_go/src/domain/access_token"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"

	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken       = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?"
	queryCreateAccessToken    = "INSERT into access_tokens (access_token, user_id, client_id, expires) values (?,?,?,?)"
	queryUpdateExpirationTime = "UPDATE access_tokens SET expires=? WHERE access_token=?"
)

type DBRepository interface {
	GetbyID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbrepo struct{}

func NewRepository() DBRepository {
	return &dbrepo{}
}

func (dr *dbrepo) GetbyID(id string) (*access_token.AccessToken, *errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&result.AccessToken, result.UserID, result.ClientID, result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewInternalServerError("no access token found with given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (dr *dbrepo) Create(at access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserID, at.ClientID, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (dr *dbrepo) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpirationTime, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
