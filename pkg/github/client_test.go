package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGithubTokenIsAbsent(t *testing.T) {
	client, err := NewWithAuth("johndoe", "")
	assert.Nil(t, client)
	assert.Contains(t, err.Error(), "Github Token is required")
}
func TestGithubUsernameIsAbsent(t *testing.T) {
	client, err := NewWithAuth("", "invalid")
	assert.Nil(t, client)
	assert.Contains(t, err.Error(), "Github username is required")
}
func TestGithubTokenValidity(t *testing.T) {
	client, _ := NewWithAuth("johndoe", "invalid")
	ok := client.IsTokenValid()
	assert.False(t, ok)
}
