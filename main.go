package main

import (
	"fmt"
	"main/finder"
	"os"
	"slices"
)

// TODO: Add usage instruction
func usage() {
	fmt.Printf("\n[INFO]\tYou can provide several arguments to this function.\n")
	fmt.Printf("[INFO]\tThe general use of this function is as follows:\n")
	fmt.Printf("[INFO]\tExample usage:\tmain.go [args, optional] [filepath to start from]\n\n")
	fmt.Printf("[ARGS]\t-help for this help overview; \t\t\t\tmain.go -help\n")
	fmt.Printf("[ARGS]\t-log for enabling detailed logging; \t\t\tmain.go -log\n")
	fmt.Printf("[ARGS]\t-perf for enabling performance measurement; \t\tmain.go -perf\n")
	fmt.Printf("[ARGS]\t-pretty for enabling pretty output of results; \t\tmain.go -pretty\n")

}

// Todo: Add commantline args:
// -perf for performance measures
func main() {

	if len(os.Args) < 1 {
		s := fmt.Errorf("[ERROR]: Not provided enough arguments")
		fmt.Println(s)
		usage()
		return
	}

	var curArgs []string = os.Args
	var log, perf, pretty bool

	if slices.Contains(curArgs, "-help") {
		usage()
		return
	}

	log = slices.Contains(curArgs, "-log")
	perf = slices.Contains(curArgs, "-perf")
	pretty = slices.Contains(curArgs, "-pretty")

	f := finder.FileFinder{}

	f.Initialize("/Users/florianluebke/Desktop/stuff/", log, perf, pretty)
	f.FindDuplicates()
	f.PrintResults()

}
