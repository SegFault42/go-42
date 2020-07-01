package fortytwo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

var apiURL string = "https://api.intra.42.fr/"

// Struct with api info
type ApiInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	CreatedAt   int    `json:"created_at"`
}

// Get new token
func NewClient(uid, secret, scope string) (ApiInfo, error) {
	var tokenJSON ApiInfo

	reader := strings.NewReader(`grant_type=client_credentials&client_id=` + uid + `&client_secret=` + secret + `&scope=` + scope)
	req, err := http.NewRequest("POST", apiURL+"oauth/token", reader)
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

// Post method
// func (api ApiInfo) Post(url string) (*http.Response, error) {

// }

//Get method
func (api ApiInfo) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", apiURL+url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+api.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
