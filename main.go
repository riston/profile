package profile

import (
	"fmt"

	"github.com/riston/profile/facebook"
	"github.com/riston/profile/github"
)

// Common is unified result for profile extraction
type Common struct {
	ProviderID string
	Email      string
	FirstName  string
	LastName   string
	Gender     string
	Name       string
	ProfileURL string
	PictureURL string
}

// Extract method to get the provider Common profile
func Extract(provider, accessToken string) (profile *Common, err error) {

	// TODO: the switch statement here is not good approach
	switch provider {
	case "facebook":
		fbProfile, err := facebook.Extract(accessToken)
		if err != nil {
			return profile, err
		}

		profile = &Common{
			ProviderID: fbProfile.ID,
			Email:      fbProfile.Email,
			FirstName:  fbProfile.FirstName,
			LastName:   fbProfile.LastName,
			Gender:     fbProfile.Gender,
			Name:       fbProfile.Name,
			ProfileURL: fbProfile.Link,
			PictureURL: fbProfile.Picture.Data.URL,
		}

	case "github":
		gitProfile, err := github.Extract(accessToken)
		if err != nil {
			return profile, err
		}

		profile = &Common{
			ProviderID: string(gitProfile.ID),
			Email:      gitProfile.Email,
			Name:       gitProfile.Name,
			ProfileURL: gitProfile.Link,
			PictureURL: gitProfile.Picture,
		}

	default:
		err = fmt.Errorf("There is no such provider %s", provider)
	}

	return profile, err
}
