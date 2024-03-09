package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

	authData.ExpiresIn = strings.Replace(authData.ExpiresIn, "s", "", 1)

	now := time.Now().Unix()
	seconds, _ := strconv.ParseInt(authData.ExpiresIn, 10, 64)
	now = now + seconds

	authData.ExpiresAt = time.Unix(now, 0)
}
