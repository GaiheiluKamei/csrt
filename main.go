package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	inputfile := flag.String("f", "", "the name of incorrect src file")
	delimiter := flag.String("d", "n", "delimiter(n or rn)")
	outputfile := flag.String("o", "", "the output file name")

	flag.Parse()

	if *inputfile == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := correctSRT(*inputfile, *outputfile, *delimiter); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
