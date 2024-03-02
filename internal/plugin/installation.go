package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"regexp"
	"strings"
)

type Installation struct {
	Plugin Plugin
	Path   string
}

// validateAndFormatGitHubRepoURL validates and formats the GitHub URL
func validateAndFormatGitHubRepoURL(repo string) (string, error) {
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

// extractPluginName extracts the plugin name from the repository string
func extractPluginName(repo string) string {
	parts := strings.Split(repo, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1] // Return the last part as the name
	}
	return ""
}

// Install Plugin installation from a repository in GitHub
func Install(repo, outputPath string) (*Installation, error) {
	repoURL, err := validateAndFormatGitHubRepoURL(repo)
	if err != nil {
		return nil, err
	}

	pluginName := extractPluginName(repo)
	if pluginName == "" {
		return nil, fmt.Errorf("failed to extract plugin name from repository: %s", repo)
	}

	fullOutputPath := filepath.Join(outputPath, pluginName)
	if err := exec.Command("mkdir", "-p", fullOutputPath).Run(); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %v", err)
	}

	cmd := exec.Command("git", "clone", "-v", repoURL, fullOutputPath)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to clone repository: %v", err)
	}

	pluginOutputFile := filepath.Join(outputPath, pluginName+".so")
	if _, err := os.Stat(pluginOutputFile); err == nil {
		fmt.Println("Plugin file exists, replacing...")
		if err := os.Remove(pluginOutputFile); err != nil {
			return nil, fmt.Errorf("failed to remove existing plugin file: %v", err)
		}
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to check plugin file: %v", err)
	}

	buildCmd := exec.Command("go", "build", "-buildmode=plugin", "-o", pluginOutputFile, ".")
	buildCmd.Dir = fullOutputPath
	if err := buildCmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to compile plugin: %v", err)
	}

	fmt.Printf("Plugin installed successfully: %s\n", pluginOutputFile)

	return Load(pluginOutputFile)
}

// Load dynamically loads a plugin from a compiled .so file and returns an Installation.
func Load(path string) (*Installation, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	symbol, err := p.Lookup("Plugin")
	if err != nil {
		return nil, err
	}

	plug, ok := symbol.(Plugin)
	if !ok {
		return nil, fmt.Errorf("failed to cast symbol to Plugin")
	}

	return &Installation{
		Plugin: plug,
		Path:   path,
	}, nil
}
