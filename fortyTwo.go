package fortytwo

import (
	"encoding/json"
	"os"
	"time"

	"github.com/sendgrid/rest"
)

var apiURL string = "https://api.intra.42.fr"

// APIInfo Struct with api info
type APIInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	CreatedAt   int    `json:"created_at"`
}

// NewClient Get new token
func (api APIInfo) NewClient(uid, secret, scope string) (APIInfo, error) {
	var tokenJSON APIInfo
	url := "/oauth/token"

	param := `grant_type=client_credentials&client_id=` +
		uid +
		`&client_secret=` +
		secret +
		`&scope=` +
		scope

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := api.Post(url, []byte(param), headers)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic("NewClient() : " + resp.Body)
	}

	if err = json.Unmarshal([]byte(resp.Body), &tokenJSON); err != nil {
		return tokenJSON, err
	}

	// store it in env
	err = os.Setenv("INTRA_TOKEN", tokenJSON.AccessToken)
	if err != nil {
		panic(err)
	}

	return tokenJSON, nil
}

// CheckToken Renew token if expired
func (api APIInfo) CheckToken() APIInfo {

	var client APIInfo
	var err error
	now := time.Now()
	// get new token if is expired
	if int(now.Unix()) >= api.CreatedAt+api.ExpiresIn {
		client, err = api.NewClient(os.Getenv("INTRA_CLIENT_ID"), os.Getenv("INTRA_CLIENT_SECRET"), os.Getenv("INTRA_SCOPE"))
		if err != nil {
			panic(err)
		}
	}

	return client
}

// Get req
func (api APIInfo) Get(url string, queryParams map[string]string, headers map[string]string) (resp *rest.Response, err error) {

	request := rest.Request{
		Method:      rest.Get,
		BaseURL:     apiURL + url,
		Headers:     headers,
		QueryParams: queryParams,
	}

	response, err := rest.API(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Post req
func (api APIInfo) Post(url string, body []byte, headers map[string]string) (resp *rest.Response, err error) {

	request := rest.Request{
		Method:  rest.Post,
		BaseURL: apiURL + url,
		Body:    body,
		Headers: headers,
	}

	response, err := rest.API(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Patch req
func (api APIInfo) Patch(url string, body []byte, headers map[string]string) (resp *rest.Response, err error) {

	request := rest.Request{
		Method:  rest.Patch,
		BaseURL: apiURL + url,
		Body:    body,
		Headers: headers,
	}

	response, err := rest.API(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
