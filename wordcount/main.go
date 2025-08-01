package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Check if the user provided at least one file argument
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	// Create a map to hold the word counts
	result := make(map[string]int64)
	// Iterate over each file provided in the command line arguments
	for _, filename := range os.Args[1:] {
		// Open the file
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
			continue
		}
		// Create a scanner to read the file line by line
		scanner := bufio.NewScanner(file)
		// Read each line and count occurrences
		for scanner.Scan() {
			line := scanner.Text()
			result[line]++
		}
		// Check for errors during scanning
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning file %s: %v\n", filename, err)
		}
		// Close the file after reading
		file.Close()
	}
	// Print the results, filtering out lines with less than 2 occurrences
	for line, count := range result {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
