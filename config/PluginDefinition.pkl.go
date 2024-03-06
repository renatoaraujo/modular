// Code generated from Pkl module `modular.gallery.config`. DO NOT EDIT.
package config

type PluginDefinition interface {
	GetName() string

	GetDescription() string
}

var _ PluginDefinition = (*PluginDefinitionImpl)(nil)

type PluginDefinitionImpl struct {
	Name string `pkl:"name"`

	Description string `pkl:"description"`
}

func (rcv *PluginDefinitionImpl) GetName() string {
	return rcv.Name
}

func (rcv *PluginDefinitionImpl) GetDescription() string {
	return rcv.Description
}
