package cmd

import (
	"github.com/NicklasWallgren/go-template/config"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	*cobra.Command
}

func NewRootCommand(assets *config.AssetsConfig) *RootCommand {
	return &RootCommand{&cobra.Command{
		Use:   "go-template",
		Short: "An go template ",
		Long:  assets.Logo,
	}}
}

func (r RootCommand) Add(command Command, boot func(runner CommandRunner)) {
	wrappedCmd := &cobra.Command{
		Use:   command.Use(),
		Short: command.Short(),
		Run: func(c *cobra.Command, args []string) {
			boot(command.Run(c))
		},
	}
	command.Setup(wrappedCmd)
	r.Command.AddCommand(wrappedCmd)
}

type CommandRunner interface{}

// Command interface is used to implement sub-commands in the system.
type Command interface {
	// Use is the one-line usage message.
	// Recommended syntax is as follow:
	//   [ ] identifies an optional argument. Arguments that are not enclosed in brackets are required.
	//   ... indicates that you can specify multiple values for the previous argument.
	//   |   indicates mutually exclusive information. You can use the argument to the left of the separator or the
	//       argument to the right of the separator. You cannot use both arguments in a single use of the command.
	//   { } delimits a set of mutually exclusive arguments when one of the arguments is required. If the arguments are
	//       optional, they are enclosed in brackets ([ ]).
	// Example: add [-F file | -D dir]... [-f format] profile
	Use() string

	// Short returns string about short description of the command
	// the string is shown in help screen of cobra command
	Short() string

	// Setup is used to setup flags or pre-run steps for the command.
	//
	// For example,
	//  cmd.Flags().IntVarP(&r.num, "num", "n", 5, "description")
	//
	Setup(cmd *cobra.Command)

	// Run runs the command runner
	// run returns command runner which is a function with dependency
	// injected arguments.
	//
	// For example,
	//  Command{
	//   Run: func(l lib.Logger) {
	// 	   l.Info("i am working")
	// 	 },
	//  }
	//
	Run(cmd *cobra.Command) CommandRunner
}
