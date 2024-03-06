package plugin

// Installation represents a plugin that has been installed. It includes the actual
// plugin interface and the path to the compiled plugin file.
type Installation struct {
	Plugin Plugin
	Path   string
}
