package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hariprathap-hp/bookstore_oauth_go/oauth/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	headerXPublic   = "X-Public"
	headerXClientId = "X-Client-Id"
	headerXCallerId = "X-Caller-Id"

	paramAccessToken = "access_token"
)

type accessToken struct {
	ID       string `json:"id"`
	UserID   int64  `json:"user_id"`
	ClientID int64  `json:"client_id"`
}

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081/",
		Timeout: 100 * time.Millisecond,
	}
)

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}

func GetCallerId(request *http.Request) int64 {
	if request == nil {
		return 0
	}
	callerId, err := strconv.ParseInt(request.Header.Get(headerXCallerId), 10, 64)
	if err != nil {
		return 0
	}
	return callerId
}

func cleanRequest(request *http.Request) {
	if request == nil {
		return
	}
	request.Header.Del(headerXClientId)
	request.Header.Del(headerXCallerId)
}

func GetClientId(request *http.Request) int64 {
	if request == nil {
		return 0
	}
	clientId, err := strconv.ParseInt(request.Header.Get(headerXClientId), 10, 64)
	if err != nil {
		return 0
	}
	return clientId
}

func AuthenticateAccessToken(request *http.Request) *errors.RestErr {
	if request == nil {
		return nil
	}
	cleanRequest(request)
	accessTokenID := strings.TrimSpace(request.URL.Query().Get(paramAccessToken))
	if accessTokenID == "" {
		return errors.NewHTTPBadRequestError("access token can not be null")
	}

	at, err := getAccessToken(accessTokenID)
	if err != nil {
		if err.Status == http.StatusNotFound {
			return nil
		}
		return err
	}

	request.Header.Add(headerXClientId, fmt.Sprintf("%v", at.ClientID))
	request.Header.Add(headerXCallerId, fmt.Sprintf("%v", at.UserID))
	return nil
}

func getAccessToken(accessTokenID string) (*accessToken, *errors.RestErr) {
	fmt.Println(accessTokenID)
	response := oauthRestClient.Get(fmt.Sprintf("/oauth/access_token/%s", accessTokenID))
	if response != nil {
		return nil, errors.NewHTTPInternalServerError("invalid restclient response when trying to get access token")
	}
	if response.StatusCode > 299 {
		return nil, errors.NewHTTPInternalServerError("invalid error interface when trying to get access token")
	}
	var at accessToken
	if err := json.Unmarshal(response.Bytes(), &at); err != nil {
		return nil, errors.NewHTTPInternalServerError("error while unmarshal of the result bytes")
	}
	return &at, nil
}
