package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"renatoaraujo/modular/internal/plugin"
)

var repository string

func init() {
	installCmd.Flags().StringVarP(&repository, "repository", "r", "", "Repository of the module")
	if err := installCmd.MarkFlagRequired("repository"); err != nil {
		fmt.Println("Error marking 'repository' flag as required:", err)
	}
	pluginCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install [OPTIONS]",
	Short: "Install plugins from GitHub",
	RunE: func(cmd *cobra.Command, args []string) error {
		installation, err := plugin.Install(repository, pluginsPath)
		if err != nil {
			return err
		}

		return initialisePlugin(installation.Path)
	},
}
