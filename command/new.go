package command

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/vharish836/apig/apig"
)

const (
	defaultDatabase = "sqlite"
	defaultVCS      = "github.com"
	defaultNamespace = "api"
)

// NewCommand ...
type NewCommand struct {
	Meta

	module    string
	project   string
	namespace string
	database  string
}

// Run ...
func (c *NewCommand) Run(args []string) int {
	if err := c.parseArgs(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return apig.Skeleton(c.module, c.project, c.namespace, c.database)
}

func (c *NewCommand) parseArgs(args []string) error {
	flag := flag.NewFlagSet("apig", flag.ContinueOnError)

	flag.StringVar(&c.module, "m", "", "Module name")
	flag.StringVar(&c.module, "module", "", "Module name")
	flag.StringVar(&c.namespace, "n", defaultNamespace, "Namespace of API")
	flag.StringVar(&c.namespace, "namespace", defaultNamespace, "Namespace of API")
	flag.StringVar(&c.database, "d", defaultDatabase, "Database engine [sqlite,postgres,mysql]")
	flag.StringVar(&c.database, "database", defaultDatabase, "Database engine [sqlite,postgres,mysql]")

	if err := flag.Parse(args); err != nil {
		return err
	}
	if 0 < flag.NArg() {
		c.project = flag.Arg(0)
	}

	if c.project == "" {
		return errors.New("please specify project name")
	}

	if c.module == "" {
		return errors.New("please specify module name")
	}

	return nil
}

// Synopsis ...
func (c *NewCommand) Synopsis() string {
	return "Generate boilerplate"
}

// Help ...
func (c *NewCommand) Help() string {
	helpText := `
Usage: apig new [options] PROJECTNAME

  Generate go project and its boilerplate

Options:
  -database=database, -d     Database engine [sqlite,postgres,mysql] (default: sqlite)
  -namespace=namepace, -n    Namespace of API (default: api)
  -module=name               Module name to use (default: "" (blank string))
`
	return strings.TrimSpace(helpText)
}
