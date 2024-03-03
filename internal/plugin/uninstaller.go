package plugin

import (
	"fmt"
	"os"
)

// Uninstaller facilitates the uninstallation of plugins.
// It requires an instance of Installation to uninstall a specific plugin.
type Uninstaller struct {
	Installation *Installation
	FS           FileSystemHandler // Reuse FileSystemHandler interface for file operations
}

// NewUninstaller creates a new Uninstaller with the given Installation.
func NewUninstaller(installation *Installation, fs FileSystemHandler) *Uninstaller {
	return &Uninstaller{
		Installation: installation,
		FS:           fs,
	}
}

// Uninstall removes the plugin associated with the Installation.
// It first checks if the plugin file exists and returns a specific error if it does not.
func (u *Uninstaller) Uninstall() error {
	_, err := u.FS.Stat(u.Installation.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("plugin isn't installed: %s", u.Installation.Path)
		}
		return fmt.Errorf("failed to access plugin file: %v", err)
	}

	if err := u.FS.Remove(u.Installation.Path); err != nil {
		return fmt.Errorf("failed to remove plugin file: %v", err)
	}
	fmt.Println("Plugin removed successfully:", u.Installation.Path)
	return nil
}
