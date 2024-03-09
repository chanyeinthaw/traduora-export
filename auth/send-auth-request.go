package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chanyeinthaw/traduora-sync/config"
	"net/http"
	"os"
)

func sendAuthRequest() (resp *http.Response) {
	data := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     config.ClientId(),
		"client_secret": config.ClientSecret(),
	}
	jsonData, _ := json.Marshal(data)

	url := config.ApiURL("/api/v1/auth/token")
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Failed to send authentication request")
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Invalid client credentials")
		os.Exit(1)
	}

	return
}
