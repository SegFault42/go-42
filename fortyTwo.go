package fortyTwo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

var ApiToken string = ""
var ApiURL string = "https://api.intra.42.fr/v2/"

type sToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	CreatedAt   int    `json:"created_at"`
}

func NewClient(uid, secret string) error {
	reader := strings.NewReader(`grant_type=client_credentials&client_id=` + uid + `&client_secret=` + secret + `&scope=projects%20public%20tig`)
	req, err := http.NewRequest("POST", "https://api.intra.42.fr/oauth/token", reader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var tokenJSON sToken

	if err = json.Unmarshal([]byte(body), &tokenJSON); err != nil {
		return err
	}

	ApiToken = tokenJSON.AccessToken

	return nil
}
