package plugin_test

import (
	"errors"
	"testing"

	"github.com/renatoaraujo/modular/internal/plugin"
	mocks "github.com/renatoaraujo/modular/internal/plugin/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDefaultPluginLoader_Load(t *testing.T) {
	mockPlugin := &mocks.MockPlugin{}
	mockPlugin.On("Execute", mock.Anything).Return(nil)

	testCases := []struct {
		name          string
		path          string
		setupMocks    func(*mocks.MockOpener, *mocks.MockSymbolLoader)
		expectedError bool
	}{
		{
			name: "Successful Load",
			path: "path/to/valid/plugin.so",
			setupMocks: func(opener *mocks.MockOpener, loader *mocks.MockSymbolLoader) {
				opener.On("Open", mock.Anything).Return(loader, nil)
				loader.On("Lookup", "Plugin").Return(mockPlugin, nil)
			},
			expectedError: false,
		},
		{
			name: "Open Error",
			path: "path/to/nonexistent/plugin.so",
			setupMocks: func(opener *mocks.MockOpener, loader *mocks.MockSymbolLoader) {
				opener.On("Open", mock.Anything).Return(nil, errors.New("open error"))
			},
			expectedError: true,
		},
		{
			name: "Lookup Error",
			path: "path/to/valid/plugin.so",
			setupMocks: func(opener *mocks.MockOpener, loader *mocks.MockSymbolLoader) {
				opener.On("Open", mock.Anything).Return(loader, nil)
				loader.On("Lookup", "Plugin").Return(nil, errors.New("lookup error"))
			},
			expectedError: true,
		},
		{
			name: "Invalid Type Assertion",
			path: "path/to/valid/plugin.so",
			setupMocks: func(opener *mocks.MockOpener, loader *mocks.MockSymbolLoader) {
				opener.On("Open", mock.Anything).Return(loader, nil)
				loader.On("Lookup", "Plugin").Return(&struct{}{}, nil)
			},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockOpener := new(mocks.MockOpener)
			mockSymbolLoader := new(mocks.MockSymbolLoader)
			tc.setupMocks(mockOpener, mockSymbolLoader)

			loader := plugin.NewDefaultPluginLoader()
			loader.Opener = mockOpener

			_, err := loader.Load(tc.path)

			if tc.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			mockOpener.AssertExpectations(t)
			mockSymbolLoader.AssertExpectations(t)
		})
	}
}
