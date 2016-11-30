package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/ruxlang/ruxlang/internal/app"
)

type (
	// CommandRunner runs a command
	CommandRunner interface {
		// Run executes the current command and provides the parser
		Run(*arg.Parser) error
	}

	// Commands
	VersionCommand struct{}
	HelpCommand    struct {
		Commands []string `arg:"positional" help:"the commands to get help for"`
	}
	EnvCommand struct{}

	App struct {
		EnvCmd     *EnvCommand     `arg:"subcommand:env" help:"print environment variables and exit"`
		HelpCmd    *HelpCommand    `arg:"subcommand:help" help:"display this help and exit"`
		VersionCmd *VersionCommand `arg:"subcommand:version" help:"display version and exit"`
		EnvArg     bool            `arg:"--env,-e" help:"print environment variables and exit"`
	}
)

var args App

func (*App) Version() string {
	return app.Version()
}

func (c *VersionCommand) Run(*arg.Parser) error {
	fmt.Println(app.Version())
	return nil
}

func (c *HelpCommand) Run(parser *arg.Parser) error {
	parser.WriteHelpForSubcommand(os.Stdout, c.Commands...)
	return nil
}

func (*EnvCommand) Run(*arg.Parser) error {
	pkg, err := app.GetPackagePath()
	if err != nil {
		return err
	}

	fmt.Printf("%s=%s\n", app.PACKAGE_PATH, pkg)
	fmt.Printf("%s=%s\n", app.TEMP_PATH, app.GetTempPath())
	return nil
}

func main() {
	// Copy commands to arguments list with the command matcher appended
	parser := arg.MustParse(&args)
	runner, ok := parser.Subcommand().(CommandRunner)
	if !ok {
		if args.EnvArg {
			runner = &EnvCommand{}
		} else {
			parser.Fail("this is not a valid command")
			return
		}
	}

	err := runner.Run(parser)
	if err != nil {
		parser.Fail(err.Error())
		return
	}
}
