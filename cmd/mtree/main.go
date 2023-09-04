package main

import (
	"fmt"
	"os"

	"github.com/ssoriche/stools/pkg/mtree"
	"github.com/ssoriche/stools/pkg/version"
)

func main() {
	cmd := mtree.NewCommand()
	cmd.AddCommand(version.NewCommand())

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		os.Exit(1)
	}
}
