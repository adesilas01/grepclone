package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Boolean flags that change how the program behaves
	countPtr := flag.Bool("c", false, "count matches")
	casePtr := flag.Bool("i", false, "case insensitive search")
	numLinePtr := flag.Bool("n", false, "show line numbers")
	wholeMatchPtr := flag.Bool("w", false, "match whole words only")

	flag.Parse()

	// Non-flag arguments
	args := flag.Args()

	// We need at least:
	// 1. the search query
	// 2. the file name
	if len(args) < 2 {
		fmt.Println("Usage: go run . [query] [file] [-i] [-n] [-c] [-w]")
		return
	}

	query := args[0]
	file := args[1]

	// Read the whole file
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Keep the file split by lines.
	// This matters because we want line-based output.
	lines := strings.Split(string(data), "\n")

	// For case-insensitive search,
	// normalize the query once before looping.
	if *casePtr {
		query = strings.ToLower(query)
	}

	matchCount := 0

	for i, line := range lines {
		// Keep the original line for printing,
		// but create a copy for comparison.
		compareLine := line

		if *casePtr {
			compareLine = strings.ToLower(line)
		}

		matched := false

		// Whole-word mode:
		// compare each word separately.
		if *wholeMatchPtr {
			words := strings.Fields(compareLine)

			for _, word := range words {
				// Remove common punctuation so
				// "hello," still matches "hello"
				cleanWord := strings.Trim(word, ".,!?\"'():;")

				if cleanWord == query {
					matched = true
					break
				}
			}
		} else {
			// Normal mode:
			// match if query appears anywhere in the line
			matched = strings.Contains(compareLine, query)
		}

		if matched {
			matchCount++

			// If -c is used, only count.
			// Printing happens after the loop.
			if !*countPtr {
				if *numLinePtr {
					fmt.Printf("%d: %s\n", i+1, line)
				} else {
					fmt.Println(line)
				}
			}
		}
	}

	// Print total number of matches if requested
	if *countPtr {
		fmt.Println(matchCount)
	}
}
