package plugin

import "os"

// DefaultFileSystemHandler implements FileSystemHandler using the os package.
type DefaultFileSystemHandler struct{}

// MkdirAll creates a directory named path, along with any necessary parents.
func (fs *DefaultFileSystemHandler) MkdirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// Remove removes the named file or directory.
func (fs *DefaultFileSystemHandler) Remove(name string) error {
	return os.Remove(name)
}

// Stat returns the FileInfo structure describing the named file.
func (fs *DefaultFileSystemHandler) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
