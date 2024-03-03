package plugin_test

import (
	"errors"
	"testing"

	"github.com/renatoaraujo/modular/internal/plugin"
	mocks "github.com/renatoaraujo/modular/internal/plugin/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUninstaller_Uninstall(t *testing.T) {
	tests := []struct {
		name          string
		prepareMock   func(*mocks.MockFileSystemHandler)
		installation  *plugin.Installation
		expectedError string
	}{
		{
			name: "Success",
			prepareMock: func(fs *mocks.MockFileSystemHandler) {
				fs.On("Stat", mock.Anything).Return(nil, nil)
				fs.On("Remove", mock.Anything).Return(nil)
			},
			installation: &plugin.Installation{Path: "/path/to/plugin"},
		},
		{
			name: "File_Does_Not_Exist",
			prepareMock: func(fs *mocks.MockFileSystemHandler) {
				fs.On("Stat", mock.Anything).Return(nil, errors.New("file does not exist"))
			},
			installation:  &plugin.Installation{Path: "/path/to/missing/plugin"},
			expectedError: "failed to access plugin file: file does not exist",
		},
		{
			name: "Remove_Fails",
			prepareMock: func(fs *mocks.MockFileSystemHandler) {
				fs.On("Stat", mock.Anything).Return(nil, nil)
				fs.On("Remove", mock.Anything).Return(errors.New("remove failed"))
			},
			installation:  &plugin.Installation{Path: "/path/to/plugin"},
			expectedError: "failed to remove plugin file: remove failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFSHandler := new(mocks.MockFileSystemHandler)
			uninstaller := plugin.NewUninstaller(tt.installation)
			uninstaller.FS = mockFSHandler
			tt.prepareMock(mockFSHandler)

			err := uninstaller.Uninstall()

			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}

			mockFSHandler.AssertExpectations(t)
		})
	}
}
