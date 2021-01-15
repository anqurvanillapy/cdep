// package main is the Cdep CLI.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/anqurvanillapy/cdep"
)

var (
	version   = "0.1.0"
	isHelp    bool
	isVersion bool
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

	cdep.New().Init()
	return
}

func main() {
	parseArgs()
	os.Exit(run())
}
