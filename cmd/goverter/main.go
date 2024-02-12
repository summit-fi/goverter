package main

import (
	"fmt"
	"os"

	"github.com/emp1re/goverter-test"
)

func main() {
	cfg, err := cli.Parse(os.Args)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	//cfg := &goverter.GenerateConfig{
	//	PackagePatterns: []string{""},
	//	WorkingDir:      "/home/vsl/go/src/govert/",
	//	Global: config.RawLines{
	//		Lines:    nil,
	//		Location: "global",
	//	},
	//	BuildTags:             "goverter",
	//	OutputBuildConstraint: "!goverter",
	//}
	if err := goverter.GenerateConverters(cfg); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
