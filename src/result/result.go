// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

import (
	"fmt"
	"sort"
	"strings"

	"utils"
)

const (
	thousandsSeparator = ' '
	formatString       = "%-11s %10s %12s %13s\n"
	formatStringLength = 11 + 1 + 10 + 1 + 12 + 1 + 13
)

// Entry
type Entry struct {
	ExtensionName string
	IsBinary      bool
	NumberOfFiles int
	NumberOfLines int
	Filesize      int64
}

// Result
type Result struct {
	Directory          string
	Extensions         map[string]*Entry
	TotalNumberOfFiles int
	TotalNumberOfLines int
	TotalSize          int64
}

// ------------------------------------------
// PrintResult
// ------------------------------------------
func PrintResult(root string, result Result) {
	// Show result header
	printHeader(root)

	// Sorting keys for presentation
	var keys, binaryKeys = getKeys(result.Extensions)

	// Show result for sourcecode, only for extensions found
	for _, ext := range keys {
		printEntry(result.Extensions[ext])
	}

	// For convenience, show result for binaries separately
	for _, ext := range binaryKeys {
		printEntry(result.Extensions[ext])
	}

	// Show footer
	printFooter(result)
}

// ------------------------------------------
// getKeys
// ------------------------------------------
func getKeys(extensions map[string]*Entry) ([]string, []string) {
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

// ------------------------------------------
// printHeader
// ------------------------------------------
func printHeader(root string) {
	fmt.Printf("\nDirectory processed:\n")
	fmt.Printf("%v\n", root)
	fmt.Printf("%s\n", strings.Repeat("-", formatStringLength))
	fmt.Printf(formatString, "filetype", "#files", "#lines", "size")
	fmt.Printf("%s\n", strings.Repeat("-", formatStringLength))
}

// ------------------------------------------
// printEntry
// ------------------------------------------
func printEntry(entry *Entry) {
	var numberOfLinesString = ""

	if !entry.IsBinary {
		numberOfLinesString = utils.IntToString(entry.NumberOfLines, thousandsSeparator)
	}

	fmt.Printf(
		formatString,
		entry.ExtensionName,
		utils.IntToString(entry.NumberOfFiles, thousandsSeparator),
		numberOfLinesString,
		utils.Int64ToString(entry.Filesize, thousandsSeparator),
	)
}

// ------------------------------------------
// printFooter
// ------------------------------------------
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
			utils.IntToString(result.TotalNumberOfFiles, thousandsSeparator),
			utils.IntToString(result.TotalNumberOfLines, thousandsSeparator),
			utils.Int64ToString(result.TotalSize, thousandsSeparator),
		)
	}
}
