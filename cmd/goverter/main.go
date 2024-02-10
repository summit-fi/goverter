package main

import (
	"fmt"
	"os"

	"github.com/emp1re/goverter-test/cli"

	goverter "github.com/emp1re/goverter-test"
)

func main() {
	cfg, err := cli.Parse(os.Args)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err = goverter.GenerateConverters(cfg); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
