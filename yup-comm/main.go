package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/comm"
)

const (
	flagSuppressColumn1 = "suppress-column-1"
	flagSuppressColumn2 = "suppress-column-2"
	flagSuppressColumn3 = "suppress-column-3"
	flagCheckOrder      = "check-order"
	flagTotal           = "total"
)

func main() {
	app := &cli.App{
		Name:  "comm",
		Usage: "compare two sorted files line by line",
		UsageText: `comm [OPTIONS] FILE1 FILE2

   Compare sorted files FILE1 and FILE2 line by line.

   With no options, produce three-column output. Column one contains
   lines unique to FILE1, column two contains lines unique to FILE2,
   and column three contains lines common to both files.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagSuppressColumn1,
				Aliases: []string{"1"},
				Usage:   "suppress column 1 (lines unique to FILE1)",
			},
			&cli.BoolFlag{
				Name:    flagSuppressColumn2,
				Aliases: []string{"2"},
				Usage:   "suppress column 2 (lines unique to FILE2)",
			},
			&cli.BoolFlag{
				Name:    flagSuppressColumn3,
				Aliases: []string{"3"},
				Usage:   "suppress column 3 (lines that appear in both files)",
			},
			&cli.BoolFlag{
				Name:  flagCheckOrder,
				Usage: "check that the input is correctly sorted",
			},
			&cli.BoolFlag{
				Name:  flagTotal,
				Usage: "output a summary",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "comm: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (requires exactly 2 files)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, gloo.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.Bool(flagSuppressColumn1) {
		params = append(params, SuppressColumn1)
	}
	if c.Bool(flagSuppressColumn2) {
		params = append(params, SuppressColumn2)
	}
	if c.Bool(flagSuppressColumn3) {
		params = append(params, SuppressColumn3)
	}
	if c.Bool(flagCheckOrder) {
		params = append(params, CheckOrder)
	}
	if c.Bool(flagTotal) {
		params = append(params, Total)
	}

	// Create and execute the comm command
	cmd := Comm(params...)
	return gloo.Run(cmd)
}
