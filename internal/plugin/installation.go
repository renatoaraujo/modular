package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Validator defines an interface for validating and formatting GitHub
// repository URLs. This abstraction allows for easier testing and extension.
type Validator interface {
	ValidateAndFormat(repo string) (string, error)
}

// Loader defines an interface for loading plugins.
type Loader interface {
	Load(path string) (*Installation, error)
}

// FileSystemHandler defines an interface for interacting with the file system,
// allowing for easier testing by abstracting actual file system operations.
type FileSystemHandler interface {
	MkdirAll(path string) error
	Remove(name string) error
	Stat(name string) (os.FileInfo, error)
}

// Runner defines an interface for running external commands, facilitating
// testing of functions that require command execution.
type Runner interface {
	Run(name string, args ...string) error
}

// Installation represents a plugin that has been installed. It includes the actual
// plugin interface and the path to the compiled plugin file.
type Installation struct {
	Plugin Plugin
	Path   string
}

// Installer facilitates the installation of plugins from GitHub repositories.
// It uses interfaces for validating URLs, handling file system operations, and
// running external commands to enhance testability.
type Installer struct {
	Validator    Validator
	FS           FileSystemHandler
	Runner       Runner
	PluginLoader Loader
}

// NewInstaller creates a new Installer with default dependencies.
func NewInstaller() *Installer {
	return &Installer{
		Validator:    &DefaultValidator{},
		FS:           &DefaultFileSystemHandler{},
		Runner:       &DefaultRunner{},
		PluginLoader: &DefaultPluginLoader{},
	}
}

// Install downloads and compiles a plugin from a GitHub repository, then returns
// an Installation. It ensures the output directory exists, clones the repository,
// and compiles the plugin.
func (i *Installer) Install(repo, outputPath string) (*Installation, error) {
	repoURL, err := i.Validator.ValidateAndFormat(repo)
	if err != nil {
		return nil, err
	}

	pluginName := extractPluginName(repo)
	if pluginName == "" {
		return nil, fmt.Errorf("failed to extract plugin name from repository: %s", repo)
	}

	fullOutputPath := filepath.Join(outputPath, pluginName)
	if err := i.FS.MkdirAll(fullOutputPath); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %v", err)
	}

	if err := i.cloneRepo(repoURL, fullOutputPath); err != nil {
		return nil, err
	}

	pluginOutputFile := filepath.Join(outputPath, pluginName+".so")
	if err := i.buildPlugin(fullOutputPath, pluginOutputFile); err != nil {
		return nil, err
	}

	fmt.Printf("Plugin installed successfully: %s\n", pluginOutputFile)

	return i.PluginLoader.Load(pluginOutputFile)
}

// cloneRepo clones the repository from repoURL into outputPath.
func (i *Installer) cloneRepo(repoURL, outputPath string) error {
	return i.Runner.Run("git", "clone", "-v", repoURL, outputPath)
}

// buildPlugin compiles the plugin source code into a shared object file at outputFile.
func (i *Installer) buildPlugin(sourcePath, outputFile string) error {
	return i.Runner.Run("go", "build", "-buildmode=plugin", "-o", outputFile, sourcePath)
}

// extractPluginName extracts the plugin name from the repository URL or path.
func extractPluginName(repo string) string {
	parts := strings.Split(repo, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1] // Return the last part as the name
	}
	return ""
}

// Remove deletes the plugin file associated with the Installation.
func (i *Installation) Remove() error {
	if err := os.Remove(i.Path); err != nil {
		return fmt.Errorf("failed to remove plugin file: %v", err)
	}
	fmt.Println("Plugin removed successfully:", i.Path)
	return nil
}
