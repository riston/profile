package profile

import (
	"os"
	"testing"

	"github.com/riston/profile/facebook"
	"github.com/riston/profile/github"
	"github.com/stretchr/testify/assert"
)

// TODO: Create mock HTTP response pages
func TestFacebookExtract(t *testing.T) {

	profile, err := facebook.Extract(os.Getenv("FB_TEST_TOKEN"))
	assert.Nil(t, err)

	assert.Equal(t, profile.Name, "Open Graph Test User")
	assert.Equal(t, profile.FirstName, "Open")

	assert.Equal(t, profile.Picture.Data.Silhouette, true)
}

func TestGithubExtract(t *testing.T) {

	profile, err := github.Extract(os.Getenv("GITHUB_TEST_TOKEN"))
	assert.Nil(t, err)

	assert.Equal(t, profile.Name, "Risto Novik")
}

func TestCommonExtract(t *testing.T) {

	Extract("facebook", os.Getenv("FB_TEST_TOKEN"))
}
