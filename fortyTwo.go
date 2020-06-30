package fortyTwo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

var API_TOKEN string = ""
var API_URL string = "https://api.intra.42.fr/v2/"

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

	if err = json.Unmarshal([]byte(body), &tokenJSON); err != nil {}
		return err
	}

	API_TOKEN = tokenJSON.AccessToken

	return nil
}
