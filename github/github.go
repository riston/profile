package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const baseURL = "https://api.github.com/user"
const fields = "user,user:email"

type Profile struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	Link    string `json:"html_url"`
	Name    string `json:"name"`
	Picture string `json:"avatar_url"`
}

// Extract is method to request and extract the profile
func Extract(accessToken string) (Profile, error) {
	var profile Profile

	url, err := getURL(accessToken, fields)
	if err != nil {
		return profile, fmt.Errorf("Could not construct the profile request url")
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return profile, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &profile); err != nil {
		return profile, err
	}

	return profile, nil
}

func getURL(token, fields string) (*url.URL, error) {

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Add("access_token", token)
	q.Add("scope", fields)
	q.Add("token_type", "bearer")
	u.RawQuery = q.Encode()

	return u, nil
}
