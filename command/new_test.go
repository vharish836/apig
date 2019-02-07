package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

// TestNewCommand_implement ...
func TestNewCommand_implement(t *testing.T) {
	var _ cli.Command = &NewCommand{}
}
