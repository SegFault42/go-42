package fortytwo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

var APIURL string = "https://api.intra.42.fr/v2/"

type sToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	CreatedAt   int    `json:"created_at"`
}

func NewClient(uid, secret, scope string) (sToken, error) {
	var tokenJSON sToken

	reader := strings.NewReader(`grant_type=client_credentials&client_id=` + uid + `&client_secret=` + secret + `&scope=` + scope)
	req, err := http.NewRequest("POST", "https://api.intra.42.fr/oauth/token", reader)
	if err != nil {
		return tokenJSON, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return tokenJSON, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tokenJSON, err
	}

	if err = json.Unmarshal([]byte(body), &tokenJSON); err != nil {
		return tokenJSON, err
	}

	return tokenJSON, nil
}
