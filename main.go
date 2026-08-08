package main

import (
	"fmt"
	"main/finder"
	"os"
	"slices"
)

// Todo: Add commantline args:
// -perf for performance measures
func main() {

	if len(os.Args) < finder.NO_OWN_ARG {
		s := fmt.Errorf("[ERROR]: Not provided enough arguments")
		fmt.Println(s)
		finder.Usage()
		return
	}

	var curArgs []string = os.Args
	var log, perf, pretty, info bool

	if slices.Contains(curArgs, "-help") {
		finder.Usage()
		return
	}

	log = slices.Contains(curArgs, "-log")
	perf = slices.Contains(curArgs, "-perf")
	pretty = slices.Contains(curArgs, "-pretty")
	info = slices.Contains(curArgs, "-info")

	path := curArgs[len(curArgs)-1]

	// position where path is expected is control arg
	if path[0] == '-' {
		s := fmt.Errorf("[ERROR]: Missing path to start search from. Only provided config args: '[%+v]'", curArgs)
		fmt.Println(s)
		finder.Usage()
		return
	}

	f := finder.FileFinder{}

	f.Initialize(path, log, perf, pretty, info)
	f.Run()
}
