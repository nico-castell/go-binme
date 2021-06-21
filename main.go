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

	// Find out if we're outputting to terminal or pipe. Based on:
	//   https://rosettacode.org/wiki/Check_output_device_is_a_terminal#Go
	fileInfo, _ := os.Stdout.Stat()
	isTerminal := (fileInfo.Mode() & os.ModeCharDevice) != 0

	// Process Stdin
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		for i := 0; i < len(input); i++ {
			fmt.Printf(config, input[i])
		}

		// If the input stream reaches EOF, break
		if err != nil {
			if isTerminal {
				// Print a new line in the bash prompt
				fmt.Printf("\n")
			}
			break
		}
	}
}
