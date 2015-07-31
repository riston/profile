package facebook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://graph.facebook.com/v2.4/me"
	fields  = "id,email,first_name,last_name,gender,picture,link,name"
)

type Profile struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Gender    string  `json:"gender"`
	Link      string  `json:"link"`
	Name      string  `json:"name"`
	Picture   Picture `json:"picture"`
}

// Picture is a composed struct in Profile
type Picture struct {
	Data struct {
		Silhouette bool   `json:"is_silhouette"`
		URL        string `json:"url"`
	}
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
	q.Add("fields", fields)
	u.RawQuery = q.Encode()

	return u, nil
}
