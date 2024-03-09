package traduora

import (
	"github.com/chanyeinthaw/traduora-sync/auth"
	"net/http"
)

func authenticateRequests(req *http.Request) {
	req.Header.Add("Authorization", auth.BearerToken())
}
