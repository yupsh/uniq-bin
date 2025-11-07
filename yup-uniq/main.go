package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/uniq"
)

const (
	flagCount          = "count"
	flagDuplicatesOnly = "repeated"
	flagUniqueOnly     = "unique"
	flagIgnoreCase     = "ignore-case"
	flagSkipFields     = "skip-fields"
	flagSkipChars      = "skip-chars"
)

func main() {
	app := &cli.App{
		Name:  "uniq",
		Usage: "report or omit repeated lines",
		UsageText: `uniq [OPTIONS] [INPUT [OUTPUT]]

   Filter adjacent matching lines from INPUT (or standard input),
   writing to OUTPUT (or standard output).`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagCount,
				Aliases: []string{"c"},
				Usage:   "prefix lines by the number of occurrences",
			},
			&cli.BoolFlag{
				Name:    flagDuplicatesOnly,
				Aliases: []string{"d"},
				Usage:   "only print duplicate lines, one for each group",
			},
			&cli.BoolFlag{
				Name:    flagUniqueOnly,
				Aliases: []string{"u"},
				Usage:   "only print unique lines",
			},
			&cli.BoolFlag{
				Name:    flagIgnoreCase,
				Aliases: []string{"i"},
				Usage:   "ignore differences in case when comparing",
			},
			&cli.IntFlag{
				Name:    flagSkipFields,
				Aliases: []string{"f"},
				Usage:   "avoid comparing the first N fields",
			},
			&cli.IntFlag{
				Name:    flagSkipChars,
				Aliases: []string{"s"},
				Usage:   "avoid comparing the first N characters",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "uniq: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (or none for stdin)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, yup.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.Bool(flagCount) {
		params = append(params, Count)
	}
	if c.Bool(flagDuplicatesOnly) {
		params = append(params, DuplicatesOnly)
	}
	if c.Bool(flagUniqueOnly) {
		params = append(params, UniqueOnly)
	}
	if c.Bool(flagIgnoreCase) {
		params = append(params, IgnoreCase)
	}
	if c.IsSet(flagSkipFields) {
		params = append(params, SkipFields(c.Int(flagSkipFields)))
	}
	if c.IsSet(flagSkipChars) {
		params = append(params, SkipChars(c.Int(flagSkipChars)))
	}

	// Create and execute the uniq command
	cmd := Uniq(params...)
	return yup.Run(cmd)
}
