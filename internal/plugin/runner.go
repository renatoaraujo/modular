package plugin

import "os/exec"

// DefaultRunner implements Runner using the os/exec package.
type DefaultRunner struct{}

// Run executes the named program with the given arguments.
func (cr *DefaultRunner) Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	return cmd.Run()
}
