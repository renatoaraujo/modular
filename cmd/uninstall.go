package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/renatoaraujo/modular/internal/plugin"
	"github.com/spf13/cobra"
)

var name string

func init() {
	pluginCmd.AddCommand(uninstallCmd)
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall [PLUGIN]",
	Short: "Uninstall plugins",
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginName := args[0]
		if args[0] == "" {
			return fmt.Errorf("plugin name is required")
		}
		pluginPath := filepath.Join(pluginsPath, pluginName+".so")

		loader := plugin.NewDefaultPluginLoader()
		installation, err := loader.Load(pluginPath)
		if err != nil {
			return fmt.Errorf("failed to load plugin: %w", err)
		}

		uninstaller := plugin.NewUninstaller(installation)

		if err = uninstaller.Uninstall(); err != nil {
			return fmt.Errorf("failed to load plugin: %w", err)
		}

		return nil
	},
}
