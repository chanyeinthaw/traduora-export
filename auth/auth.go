package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type authRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	ExpiresAt   time.Time
}

var authData authRes

const authInfoFile = ".traduora-token.json"

func Init() {
	authed := loadAuthInfo()
	if authed {
		return
	}

	defer func() {
		if !authed {
			saveAuthInfo()
		}
	}()

	fmt.Println("Authenticating")
	resp := sendAuthRequest()
	readToken(resp)
}

func loadAuthInfo() bool {
	bytes, err := os.ReadFile(authInfoFile)
	if err != nil {
		return false
	}

	err = json.Unmarshal(bytes, &authData)
	if err != nil {
		return false
	}

	if time.Now().Equal(authData.ExpiresAt) || time.Now().After(authData.ExpiresAt) {
		return false
	}

	return true
}

func saveAuthInfo() {
	bytes, err := json.Marshal(authData)
	if err != nil {
		return
	}

	outFile, err := os.OpenFile(authInfoFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer func() {
		_ = outFile.Close()
	}()

	if err != nil {
		return
	}
	_, _ = outFile.WriteString(string(bytes))
}

func BearerToken() string {
	return fmt.Sprintf("Bearer %s", authData.AccessToken)
}
