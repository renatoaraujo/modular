package plugin

// Metadata provides metadata for a plugin, such as expected arguments and help text.
type Metadata interface {
	GetName() string
	GetExpectedArgs() []string
	GetHelp() string
}

// Plugin is the interface that all plugins must implement.
type Plugin interface {
	Metadata
	Initialize() error
	Execute(args map[string]string) error
}
