package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

// TestGenCommand_implement ...
func TestGenCommand_implement(t *testing.T) {
	var _ cli.Command = &GenCommand{}
}
