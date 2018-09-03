package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(strings.Join(os.Args[1:], " "))

	// args := os.Args[1:]
	// if len(args) > 0 && args[0] == "-n" {
	// 	fmt.Print(strings.Join(args[1:], " "))
	// } else {
	// 	fmt.Println(strings.Join(args, " "))
	// }

	var (
		noNewLine bool
		separator string
		printHelp bool
	)

	flag.BoolVar(&printHelp, "h", false, "Print this help")
	flag.BoolVar(&noNewLine, "n", false, "Set this flag won't print \\n")
	flag.StringVar(&separator, "s", " ", "Separator char")
	flag.Parse()

	// if -h set, prints usage
	if printHelp {
		flag.PrintDefaults()
		// os.Exit(1)
		return
	}

	args := flag.Args()

	if noNewLine {
		fmt.Print(strings.Join(args, separator))
	} else {
		fmt.Println(strings.Join(args, separator))
	}
}
