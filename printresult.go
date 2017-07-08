// Copyright 2014-2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sort"
	"strings"
)

const (
	thousandsSeparator = ' '
	// formatString consists of "filetype", "#files", "#lines", "line%", "size", "size%"
	formatString       = "%-11s %10s %12s %6s %13s %6s\n"
	formatStringLength = 11 + 1 + 10 + 1 + 12 + 1 + 6 + 1 + 13 + 1 + 6
)

// printAnalyticsHeader prints the header
func printAnalyticsHeader() {
	if *showDirectories || *showFiles {
		fmt.Println()
	}

	if *showDirectories {
		if *showOnlyIncluded {
			fmt.Printf("Shows included directories.\n")
		}
		if *showOnlyExcluded {
			fmt.Printf("Shows excluded directories.\n")
		}
	}

	if *showFiles {
		if *showOnlyIncluded {
			fmt.Printf("Shows included files.\n")
		}
		if *showOnlyExcluded {
			fmt.Printf("Shows excluded files.\n")
		}
	}

	if *showDirectories || *showFiles {
		fmt.Printf("---------------------------\n")
	}
}

// printResult
func printResult() {
	//
	// Show result header
	printHeader(root)

	// Sorting keys for presentation
	var keys, binaryKeys = getKeys(countResult.Extensions)

	// Show result for sourcecode, only for extensions found
	for _, ext := range keys {
		printEntry(countResult.Extensions[ext], countResult.TotalNumberOfLines, countResult.TotalSize)
	}

	// For convenience, show result for binaries separately
	for _, ext := range binaryKeys {
		printEntry(countResult.Extensions[ext], 0, countResult.TotalSize)
	}

	// Show footer
	printFooter(countResult)

	if *showBigFiles > 0 {
		sort.Sort(bigFiles)
		fmt.Printf("\n\nThe %3d largest files are:                 #lines\n", *showBigFiles)
		fmt.Printf("-------------------------------------------------\n")
		for i := 0; i < *showBigFiles; i++ {
			if i < len(bigFiles) {
				fmt.Printf("%-42s %6d\n", bigFiles[i].Name, bigFiles[i].Lines)
			}
		}
	}
}

// getKeys
func getKeys(extensions map[string]*ResultEntry) ([]string, []string) {
	var keys []string
	var binaryKeys []string

	for k, v := range extensions {
		if v.NumberOfFiles > 0 {
			if v.IsBinary {
				binaryKeys = append(binaryKeys, k)
			} else {
				keys = append(keys, k)
			}
		}
	}

	sort.Strings(keys)
	sort.Strings(binaryKeys)

	return keys, binaryKeys
}

// printHeader
func printHeader(root string) {
	fmt.Printf("\nDirectory processed:\n")
	fmt.Printf("%v\n", root)
	fmt.Printf("%s\n", strings.Repeat("-", formatStringLength))
	fmt.Printf(formatString, "filetype", "#files", "#lines", "line%", "size", "size%")
	fmt.Printf("%s\n", strings.Repeat("-", formatStringLength))
}

// printEntry
func printEntry(entry *ResultEntry, totalNumberOfLines int, totalSize int64) {
	var numberOfLinesString = ""
	var percentageString = ""

	if !entry.IsBinary {
		numberOfLinesString = intToString(entry.NumberOfLines, thousandsSeparator)

		// Show percentage
		if totalNumberOfLines > 0 {
			var percentage = float64(entry.NumberOfLines) * float64(100) / float64(totalNumberOfLines)
			percentageString = fmt.Sprintf("%.1f", round(percentage, 1))
		}
	}

	var sizePercentage = float64(entry.Filesize) * float64(100) / float64(totalSize)
	var sizePercentageString = fmt.Sprintf("%.1f", round(sizePercentage, 1))

	fmt.Printf(
		formatString,
		entry.ExtensionName,
		intToString(entry.NumberOfFiles, thousandsSeparator),
		numberOfLinesString,
		percentageString,
		int64ToString(entry.Filesize, thousandsSeparator),
		sizePercentageString,
	)
}

// printFooter
func printFooter(result Result) {
	// Show footer
	if result.TotalNumberOfFiles == 0 {
		fmt.Printf("No files found.\n\nCheck given directory,or maybe \ncheck extensions in config file.\n")
		fmt.Printf("%s\n", strings.Repeat("-", formatStringLength))
	} else {
		fmt.Printf("%s\n", strings.Repeat("-", formatStringLength))
		fmt.Printf(
			formatString,
			"Total:",
			intToString(result.TotalNumberOfFiles, thousandsSeparator),
			intToString(result.TotalNumberOfLines, thousandsSeparator),
			"100.0",
			int64ToString(result.TotalSize, thousandsSeparator),
			"100.0",
		)
	}
}
