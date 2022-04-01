package access_token

import (
	"strings"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	Expires     int64  `json:"expires"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)

	if at.AccessToken == "" {
		return errors.NewBadRequestError("Access Token can't be empty")
	}

	if at.UserID <= 0 {
		return errors.NewBadRequestError("Invalid user id")
	}

	if at.ClientID <= 0 {
		return errors.NewBadRequestError("Invalid client id")
	}

	if at.Expires <= 0 {
		return errors.NewBadRequestError("Invalid client id")
	}

	return nil
}
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
