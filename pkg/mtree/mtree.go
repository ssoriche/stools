package mtree

import (
	"os"

	"github.com/spf13/cobra"
)

type options struct{}

func NewCommand() *cobra.Command {
	o := options{}
	c := cobra.Command{
		Use:           "mtree",
		Short:         "make a graphical tree",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(_ *cobra.Command, paths []string) error {
			return run(&o, paths)
		},
	}

	return &c
}

func run(o *options, paths []string) error {
	tree := Tree{}

	for _, path := range paths {
		tree.Add(path)
	}

	tree.Fprint(os.Stdout, true, "")
	return nil
}
