package plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAndFormatGitHubRepoURL(t *testing.T) {
	testCases := []struct {
		name    string
		repo    string
		want    string
		wantErr bool
	}{
		{"HTTP URL", "http://github.com/example/repo", "http://github.com/example/repo.git", false},
		{"HTTPS URL", "https://github.com/example/repo", "https://github.com/example/repo.git", false},
		{"Missing scheme", "example/repo", "https://github.com/example/repo.git", false},
		{"Invalid URL", "not_a_repo", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := validateAndFormatGitHubRepoURL(tc.repo)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestExtractPluginName(t *testing.T) {
	testCases := []struct {
		repo string
		want string
	}{
		{"https://github.com/example/repo.git", "repo.git"},
		{"example/repo", "repo"},
		{"", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.repo, func(t *testing.T) {
			got := extractPluginName(tc.repo)
			assert.Equal(t, tc.want, got)
		})
	}
}
