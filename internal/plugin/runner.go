package plugin

import "os/exec"

// DefaultRunner implements Runner using the os/exec package.
type DefaultRunner struct {
	Dir string
}

// At defines where the command should be executed.
func (cr *DefaultRunner) At(path string) {
	cr.Dir = path
}

// Run executes the named program with the given arguments.
func (cr *DefaultRunner) Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	if cr.Dir != "" {
		cmd.Dir = cr.Dir
	}

	return cmd.Run()
}
