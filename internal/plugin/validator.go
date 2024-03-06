package plugin

import (
	"fmt"
	"regexp"
)

// DefaultValidator implements Validator to validate and format
// GitHub repository URLs.
type DefaultValidator struct{}

// ValidateAndFormat checks if the provided repository string is a valid GitHub URL
// or shorthand (owner/repo). It returns a formatted GitHub URL or an error if the
// format is invalid.
func (v *DefaultValidator) ValidateAndFormat(repo string) (string, error) {
	githubURLRegex := regexp.MustCompile(`^https?://github\.com/`)
	if githubURLRegex.MatchString(repo) {
		if !regexp.MustCompile(`\.git$`).MatchString(repo) {
			repo += ".git"
		}
		return repo, nil
	}

	ownerRepoRegex := regexp.MustCompile(`^[a-zA-Z0-9_.-]+/[a-zA-Z0-9_.-]+$`)
	if ownerRepoRegex.MatchString(repo) {
		return fmt.Sprintf("https://github.com/%s.git", repo), nil
	}

	return "", fmt.Errorf("invalid or non-GitHub repository format: %s", repo)
}
