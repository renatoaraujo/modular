package plugin

import (
	"fmt"
	"plugin"
)

// Opener is an interface that abstracts the mechanism of opening a plugin file.
// It returns a SymbolLoader instance for looking up symbols within the opened plugin.
type Opener interface {
	Open(path string) (SymbolLoader, error)
}

// SymbolLoader is an interface that defines the Lookup method.
// This method is used to find a symbol (function or variable) by its name within a plugin.
type SymbolLoader interface {
	Lookup(symName string) (plugin.Symbol, error)
}

// DefaultPluginOpener is a struct that implements the Opener interface.
// It provides the default mechanism for opening a plugin file.
type DefaultPluginOpener struct{}

// Open is a method on DefaultPluginOpener that opens a plugin file from the given path.
// It returns a SymbolLoader for the opened plugin, or an error if the file cannot be opened.
func (DefaultPluginOpener) Open(path string) (SymbolLoader, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}
	return &defaultPluginSymbolLoader{p: p}, nil
}

// defaultPluginSymbolLoader is a struct that wraps a *plugin.Plugin.
// It implements the SymbolLoader interface, allowing symbols to be looked up in the plugin.
type defaultPluginSymbolLoader struct {
	p *plugin.Plugin
}

// Lookup is a method on defaultPluginSymbolLoader that searches for a symbol by name.
// It returns the found symbol or an error if the symbol cannot be found.
func (d *defaultPluginSymbolLoader) Lookup(symName string) (plugin.Symbol, error) {
	return d.p.Lookup(symName)
}

// DefaultPluginLoader is a struct that encapsulates the process of loading a plugin.
// It uses an Opener to open the plugin file and then loads a specific symbol from it.
type DefaultPluginLoader struct {
	Opener Opener
}

// NewDefaultPluginLoader creates and returns a new instance of DefaultPluginLoader.
// It initializes the Opener field with DefaultPluginOpener, the default plugin opener.
func NewDefaultPluginLoader() *DefaultPluginLoader {
	return &DefaultPluginLoader{
		Opener: DefaultPluginOpener{},
	}
}

// Load is a method on DefaultPluginLoader that loads a plugin from a given path.
// It opens the plugin file, looks up a predefined symbol ("Plugin"), and returns an Installation.
// If any step fails, it returns an error.
func (l *DefaultPluginLoader) Load(path string) (*Installation, error) {
	p, err := l.Opener.Open(path)
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
