package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Parse input flags.
	config := "%08b "
	oct := flag.Bool("o", false, "Print octal values.")
	hex := flag.Bool("x", false, "Print hexadecimal values.")
	flag.Parse()

	// Only one of the flags can be true
	if *oct && *hex {
		fmt.Fprintf(os.Stderr, "You can only choose one flag between \"-x\" and \"-o\"\n")
		os.Exit(1)
	}

	// Configure output based on args
	if *oct {
		config = "%03o "
	}
	if *hex {
		config = "%02x "
	}

	// Process Stdin
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		for i := 0; i < len(input); i++ {
			fmt.Printf(config, input[i])
		}

		// If the input stream reaches EOF, break
		if err != nil {
			break
		}
	}
}
