package fortytwo

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

func (api APIInfo) LoginToUserID(login string) (string, error) {

	type sUser []struct {
		UserID int `json:"id"`
	}
	var user sUser

	login = strings.ToLower(login)
	url := "/v2/users"
	headers := map[string]string{
		"Authorization": "Bearer " + os.Getenv("INTRA_TOKEN"),
		"Content-Type":  "application/json",
	}
	queryParams := map[string]string{
		"filter[login]": login,
	}

	resp, err := api.Get(url, queryParams, headers)
	if err != nil {
		return "", err
	}

	if err = json.Unmarshal([]byte(resp.Body), &user); err != nil {
		return "", err
	}
	if len(user) == 0 {
		return "", errors.New("Login not found")
	}

	return strconv.Itoa(user[0].UserID), nil
}
