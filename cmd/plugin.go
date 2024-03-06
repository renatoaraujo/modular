package cmd

import (
	"fmt"
	"github.com/renatoaraujo/modular/internal/plugin"
	"github.com/spf13/cobra"
)

func init() {
	modularCmd.AddCommand(pluginCmd)
}

var pluginCmd = &cobra.Command{
	Use:   "plugin [COMMANDS]",
	Short: "Plugin management",
}

func initialisePlugin(file string) error {
	loader := plugin.NewDefaultPluginLoader()
	installation, err := loader.Load(file)
	if err != nil {
		return fmt.Errorf("error loading plugin from %s: %w", file, err)
	}

	cmd, err := createPluginCommand(installation.Plugin)
	if err != nil {
		return fmt.Errorf("failed to initialise plugin: %w", err)
	}
	modularCmd.AddCommand(cmd)
	return nil
}

func createPluginCommand(plugin plugin.Plugin) (*cobra.Command, error) {
	err := plugin.Initialize()
	if err != nil {
		return nil, fmt.Errorf("failed to create plugin command: %w", err)
	}
	var cmd = &cobra.Command{
		Use:   plugin.GetName(),
		Short: plugin.GetHelp(),
		Run: func(cmd *cobra.Command, args []string) {
			pluginArgs := make(map[string]string)
			for i, argName := range plugin.GetExpectedArgs() {
				if i < len(args) {
					pluginArgs[argName] = args[i]
				}
			}

			err := plugin.Execute(pluginArgs)
			if err != nil {
				fmt.Println("Error executing plugin:", err)
			}
		},
	}

	for _, arg := range plugin.GetExpectedArgs() {
		cmd.Flags().String(arg, "", "Description for "+arg)
	}

	return cmd, nil
}
