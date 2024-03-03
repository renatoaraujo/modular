package plugin_test

import (
	"fmt"
	"testing"

	"github.com/renatoaraujo/modular/internal/plugin"
	mocks "github.com/renatoaraujo/modular/internal/plugin/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestInstaller_Install(t *testing.T) {
	testCases := []struct {
		name            string
		repo            string
		outputPath      string
		mockSetup       func(*mocks.MockValidator, *mocks.MockFileSystemHandler, *mocks.MockRunner, *mocks.MockPluginLoader)
		expectedErr     string
		expectedSuccess bool
	}{
		{
			name:       "Successful Installation",
			repo:       "github.com/example/plugin",
			outputPath: "/plugins",
			mockSetup: func(v *mocks.MockValidator, fs *mocks.MockFileSystemHandler, r *mocks.MockRunner, pl *mocks.MockPluginLoader) {
				v.On("ValidateAndFormat", "github.com/example/plugin").Return("https://github.com/example/plugin", nil)
				fs.On("MkdirAll", mock.AnythingOfType("string")).Return(nil)
				r.On("Run", "git", "clone", "-v", "https://github.com/example/plugin", mock.AnythingOfType("string")).Return(nil)
				r.On("Run", "go", "build", "-buildmode=plugin", "-o", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
				pl.On("Load", mock.AnythingOfType("string")).Return(&plugin.Installation{}, nil)
			},
			expectedErr:     "",
			expectedSuccess: true,
		},
		{
			name:       "Validation Failure",
			repo:       "invalid/repo/format",
			outputPath: "/plugins",
			mockSetup: func(v *mocks.MockValidator, fs *mocks.MockFileSystemHandler, r *mocks.MockRunner, pl *mocks.MockPluginLoader) {
				v.On("ValidateAndFormat", "invalid/repo/format").Return("", fmt.Errorf("invalid repository format"))
			},
			expectedErr:     "invalid repository format",
			expectedSuccess: false,
		},
		{
			name:       "MkdirAll Failure",
			repo:       "github.com/example/plugin",
			outputPath: "/plugins",
			mockSetup: func(v *mocks.MockValidator, fs *mocks.MockFileSystemHandler, r *mocks.MockRunner, pl *mocks.MockPluginLoader) {
				v.On("ValidateAndFormat", "github.com/example/plugin").Return("https://github.com/example/plugin", nil)
				fs.On("MkdirAll", mock.AnythingOfType("string")).Return(fmt.Errorf("failed to create directory"))
			},
			expectedErr:     "failed to create directory",
			expectedSuccess: false,
		},
		{
			name:       "Clone Failure",
			repo:       "github.com/example/plugin",
			outputPath: "/plugins",
			mockSetup: func(v *mocks.MockValidator, fs *mocks.MockFileSystemHandler, r *mocks.MockRunner, pl *mocks.MockPluginLoader) {
				v.On("ValidateAndFormat", "github.com/example/plugin").Return("https://github.com/example/plugin", nil)
				fs.On("MkdirAll", mock.AnythingOfType("string")).Return(nil)
				r.On("Run", "git", "clone", "-v", "https://github.com/example/plugin", mock.AnythingOfType("string")).Return(fmt.Errorf("git clone failed"))
			},
			expectedErr:     "git clone failed",
			expectedSuccess: false,
		},
		{
			name:       "Build Failure",
			repo:       "github.com/example/plugin",
			outputPath: "/plugins",
			mockSetup: func(v *mocks.MockValidator, fs *mocks.MockFileSystemHandler, r *mocks.MockRunner, pl *mocks.MockPluginLoader) {
				v.On("ValidateAndFormat", "github.com/example/plugin").Return("https://github.com/example/plugin", nil)
				fs.On("MkdirAll", mock.AnythingOfType("string")).Return(nil)
				r.On("Run", "git", "clone", "-v", "https://github.com/example/plugin", mock.AnythingOfType("string")).Return(nil)
				r.On("Run", "go", "build", "-buildmode=plugin", "-o", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(fmt.Errorf("go build failed"))
			},
			expectedErr:     "go build failed",
			expectedSuccess: false,
		},
		{
			name:       "Plugin Load Failure",
			repo:       "github.com/example/plugin",
			outputPath: "/plugins",
			mockSetup: func(v *mocks.MockValidator, fs *mocks.MockFileSystemHandler, r *mocks.MockRunner, pl *mocks.MockPluginLoader) {
				v.On("ValidateAndFormat", "github.com/example/plugin").Return("https://github.com/example/plugin", nil)
				fs.On("MkdirAll", mock.AnythingOfType("string")).Return(nil)
				r.On("Run", "git", "clone", "-v", "https://github.com/example/plugin", mock.AnythingOfType("string")).Return(nil)
				r.On("Run", "go", "build", "-buildmode=plugin", "-o", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
				pl.On("Load", mock.AnythingOfType("string")).Return(nil, fmt.Errorf("failed to load plugin"))
			},
			expectedErr:     "failed to load plugin",
			expectedSuccess: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockValidator := new(mocks.MockValidator)
			mockFSHandler := new(mocks.MockFileSystemHandler)
			mockRunner := new(mocks.MockRunner)
			mockPluginLoader := new(mocks.MockPluginLoader)

			installer := plugin.NewInstaller()
			installer.Validator = mockValidator
			installer.FS = mockFSHandler
			installer.Runner = mockRunner
			installer.PluginLoader = mockPluginLoader

			tc.mockSetup(mockValidator, mockFSHandler, mockRunner, mockPluginLoader)

			installation, err := installer.Install(tc.repo, tc.outputPath)

			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErr)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, installation)
				assert.True(t, tc.expectedSuccess)
			}

			mockValidator.AssertExpectations(t)
			mockFSHandler.AssertExpectations(t)
			mockRunner.AssertExpectations(t)
			mockPluginLoader.AssertExpectations(t)
		})
	}
}
