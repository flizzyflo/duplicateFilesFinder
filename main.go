package main

import (
	"fmt"
	"main/finder"
	"os"
)

// TODO: Add usage instruction
func usage() {
	fmt.Println()
}

// Todo: Add commantline args:
// -l for long list.
// -log for logs of steps,
// -perf for performance measures
// create git repo
func main() {
	f := finder.FileFinder{}
	if len(os.Args) < 1 {
		usage()
		s := fmt.Errorf("[ERROR]: Not provided enough arguments")
		fmt.Println(s)
		return
	}

	f.Initialize("/Users/florianluebke/Desktop/stuff/")
	f.FindDuplicates()
	// f.PrettifyResult()
}
