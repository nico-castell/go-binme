package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Parse input flags.
	pOct := flag.Bool("o", false, "Print octal values.")
	pHex := flag.Bool("x", false, "Print hexadecimal values.")
	pDec := flag.Bool("d", false, "Print decimal values.")
	flag.Parse()

	// Prepare bitset comparison
	var oct, hex, dec byte
	if *pOct { oct = 1 }
	if *pHex { hex = 1 }
	if *pDec { dec = 1 }

	// Only one of the flags can be true
	if oct + hex + dec > 1 {
		fmt.Fprintf(os.Stderr, "You can only choose one flag between \"-x\", \"-o\", and \"-d\"\n")
		os.Exit(1)
	}

	// Configure output based on args
	config := "%08b "
	if *pOct { config = "%03o " }
	if *pHex { config = "%02x " }
	if *pDec { config = "%03d " }

	// Find out if we're outputting to terminal or pipe. Based on:
	//   https://rosettacode.org/wiki/Check_output_device_is_a_terminal#Go
	fileInfo, _ := os.Stdout.Stat()
	isTerminal := (fileInfo.Mode() & os.ModeCharDevice) != 0

	// Process Stdin
	reader := bufio.NewReader(os.Stdin)
	var input string
	var err error
	for {
		input, err = reader.ReadString('\n')
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
