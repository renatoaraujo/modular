package plugin

import (
	"fmt"
	"plugin"
)

// DefaultPluginLoader is the default implementation of Loader,
// using the standard Go plugin package.
type DefaultPluginLoader struct{}

// Load dynamically loads a compiled plugin file, returning an Installation.
func (l *DefaultPluginLoader) Load(path string) (*Installation, error) {
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
