package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Process Stdin
	for {
		input, err := reader.ReadString('\n')
		for i := 0; i < len(input); i++ {
			fmt.Printf("%08b ", input[i])
		}

		// If the input stream reaches EOF, break
		if err != nil {
			break
		}
	}
}
