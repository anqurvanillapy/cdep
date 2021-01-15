// package main is the Cdep CLI.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/anqurvanillapy/cdep"
)

var (
	version = "0.1.0"

	isHelp    bool
	isVersion bool

	isCmdInit bool
	isCmdGet  string
)

func help() {
	fmt.Print("cdep is a C/C++ package manager that just works.\n\n")
	flag.PrintDefaults()
}

func showVersion() {
	fmt.Printf("cdep v%s\n", version)
}

func parseArgs() {
	flag.Usage = help

	flag.BoolVar(&isHelp, "h", false, "show help")
	flag.BoolVar(&isVersion, "v", false, "show version")

	flag.BoolVar(&isCmdInit, "init", false, "init a cdep project")
	flag.StringVar(&isCmdGet, "get", "", "get a dependency")

	flag.Parse()
}

func run() (rc int) {
	if isHelp {
		flag.Usage()
		return
	}

	if isVersion {
		showVersion()
		return
	}

	c := cdep.New()

	if isCmdInit {
		if err := c.InitProject("."); err != nil {
			return 1
		}
		return
	}

	flag.Usage()
	return 1
}

func main() {
	parseArgs()
	os.Exit(run())
}
