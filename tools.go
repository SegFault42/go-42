package fortytwo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func (api APIInfo) LoginToUserID(login string) (string, error) {

	type sUser []struct {
		UserID int `json:"id"`
	}
	var user sUser

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

	fmt.Println(resp.Body)
	if err = json.Unmarshal([]byte(resp.Body), &user); err != nil {
		return "", err
	}

	return strconv.Itoa(user[0].UserID), nil
}
