package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func readToken(resp *http.Response) {
	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := io.ReadAll(resp.Body)

	err := json.Unmarshal(body, &authData)
	if err != nil {
		fmt.Println("Error reading authentication response")
		os.Exit(1)
	}
}
