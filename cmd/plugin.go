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
	loader := plugin.NewPluginLoader()
	installation, err := loader.Load(file)
	if err != nil {
		return fmt.Errorf("error loading plugin from %s: %w", file, err)
	}

	cmd := createPluginCommand(installation.Plugin)
	modularCmd.AddCommand(cmd)
	return nil
}

func createPluginCommand(plugin plugin.Plugin) *cobra.Command {
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

	return cmd
}
