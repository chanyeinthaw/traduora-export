package auth

import (
	"fmt"
)

type authRes struct {
	AccessToken string `json:"access_token"`
}

var authData authRes

func Init() {
	resp := sendAuthRequest()
	readToken(resp)
}

func BearerToken() string {
	return fmt.Sprintf("Bearer %s", authData.AccessToken)
}
