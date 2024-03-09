package traduora

import (
	"github.com/chanyeinthaw/traduora-export/auth"
	"net/http"
)

func authenticateRequests(req *http.Request) {
	req.Header.Add("Authorization", auth.BearerToken())
}
