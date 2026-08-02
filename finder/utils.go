package finder

import "fmt"

const START_FOLDER string = "/."

func Usage() {
	fmt.Printf("\n[INFO]\tYou can provide several arguments to this function.\n")
	fmt.Printf("[INFO]\tThe general use of this function is as follows:\n")
	fmt.Printf("[INFO]\tExample usage:\tmain.go [args, optional] [filepath to start from]\n\n")
	fmt.Printf("[ARGS]\t'-help' for this help overview; \t\t\tmain.go -help\n")
	fmt.Printf("[ARGS]\t'-log' for enabling detailed logging; \t\t\tmain.go -log\n")
	fmt.Printf("[ARGS]\t'-perf' for enabling performance measurement; \t\tmain.go -perf\n")
	fmt.Printf("[ARGS]\t'-pretty' for enabling pretty output of results; \tmain.go -pretty\n")
	fmt.Printf("[ARGS]\t'-info' for enabling information about collection; \tmain.go -info\n")

}
