package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var installPath string
var pluginsPath string
var configPath string

var modularCmd = &cobra.Command{
	Use:   "modular",
	Short: "Modular is a light modular CLI tool",
}

func Execute() {
	exitOnFailure(initConfig())
	initialisePlugins()
	exitOnFailure(modularCmd.Execute())
}

func exitOnFailure(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initialisePlugins() {
	files, err := filepath.Glob(filepath.Join(pluginsPath, "*.so"))
	if err != nil {
		fmt.Println("Error scanning for plugins:", err)
		return
	}

	for _, file := range files {
		if err = initialisePlugin(file); err != nil {
			fmt.Println("Failed to initialise plugin:", err)
		}
	}
}

// initConfig initializes Viper to read in configuration variables.
func initConfig() error {
	installPath = viper.GetString("installPath")
	if installPath == "" {
		var err error
		if installPath, err = os.UserHomeDir(); err != nil {
			return fmt.Errorf("failed to get user home directory: %v", err)
		}
	}

	installPath = filepath.Join(installPath, ".modular")
	pluginsPath = filepath.Join(installPath, "plugins")
	configPath = filepath.Join(installPath, "config")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(installPath)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")

	return nil
}
