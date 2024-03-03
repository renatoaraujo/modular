package plugin_test

import (
	"errors"
	"testing"

	"github.com/renatoaraujo/modular/internal/plugin"
	mocks_test "github.com/renatoaraujo/modular/internal/plugin/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUninstaller_Uninstall(t *testing.T) {
	tests := []struct {
		name          string
		prepareMock   func(*mocks_test.MockFileSystemHandler)
		installation  *plugin.Installation
		expectedError string
	}{
		{
			name: "Success",
			prepareMock: func(fs *mocks_test.MockFileSystemHandler) {
				fs.On("Stat", mock.Anything).Return(nil, nil)
				fs.On("Remove", mock.Anything).Return(nil)
			},
			installation: &plugin.Installation{Path: "/path/to/plugin"},
		},
		{
			name: "FileDoesNotExist",
			prepareMock: func(fs *mocks_test.MockFileSystemHandler) {
				fs.On("Stat", mock.Anything).Return(nil, errors.New("file does not exist"))
			},
			installation:  &plugin.Installation{Path: "/path/to/missing/plugin"},
			expectedError: "failed to access plugin file: file does not exist",
		},
		{
			name: "RemoveFails",
			prepareMock: func(fs *mocks_test.MockFileSystemHandler) {
				fs.On("Stat", mock.Anything).Return(nil, nil)
				fs.On("Remove", mock.Anything).Return(errors.New("remove failed"))
			},
			installation:  &plugin.Installation{Path: "/path/to/plugin"},
			expectedError: "failed to remove plugin file: remove failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFS := &mocks_test.MockFileSystemHandler{}
			tt.prepareMock(mockFS)
			uninstaller := plugin.NewUninstaller(tt.installation, mockFS)

			err := uninstaller.Uninstall()

			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}

			mockFS.AssertExpectations(t)
		})
	}
}
