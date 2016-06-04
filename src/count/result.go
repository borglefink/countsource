// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package count

import (
	"fmt"
	"sort"
	"strings"

	"utils"
)

const (
	thousandsSeparator = ' '
	// formatString consists of "filetype", "#files", "#lines", "line%", "size", "size%"
	formatString       = "%-11s %10s %12s %6s %13s %6s\n"
	formatStringLength = 11 + 1 + 10 + 1 + 12 + 1 + 6 + 1 + 13 + 1 + 6
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
// PrintAnalyticsHeader
// ------------------------------------------
func PrintAnalyticsHeader(showDirectories, showFiles, showOnlyIncluded, showOnlyExcluded bool) {
	if showDirectories || showFiles {
		fmt.Println()
	}

	if showDirectories {
		if showOnlyIncluded {
			fmt.Printf("Shows included directories.\n")
		}
		if showOnlyExcluded {
			fmt.Printf("Shows excluded directories.\n")
		}
	}

	if showFiles {
		if showOnlyIncluded {
			fmt.Printf("Shows included files.\n")
		}
		if showOnlyExcluded {
			fmt.Printf("Shows excluded files.\n")
		}
	}

	if showDirectories || showFiles {
		fmt.Printf("---------------------------\n")
	}
}

// ------------------------------------------
// PrintResult
// ------------------------------------------
func PrintResult(root string, result Result, bigFiles FileSizes) {
	// Show result header
	printHeader(root)

	// Sorting keys for presentation
	var keys, binaryKeys = getKeys(result.Extensions)

	// Show result for sourcecode, only for extensions found
	for _, ext := range keys {
		printEntry(result.Extensions[ext], result.TotalNumberOfLines, result.TotalSize)
	}

	// For convenience, show result for binaries separately
	for _, ext := range binaryKeys {
		printEntry(result.Extensions[ext], 0, result.TotalSize)
	}

	// Show footer
	printFooter(result)

	if showBigFiles > 0 {
		sort.Sort(bigFiles)
		fmt.Printf("\n\nThe %3d largest files are:                 #lines\n", showBigFiles)
		fmt.Printf("-------------------------------------------------\n")
		for i := 0; i < showBigFiles; i++ {
			fmt.Printf("%-42s %6d\n", bigFiles[i].Name, bigFiles[i].Lines)
		}
	}
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
	fmt.Printf(formatString, "filetype", "#files", "#lines", "line%", "size", "size%")
	fmt.Printf("%s\n", strings.Repeat("-", formatStringLength))
}

// ------------------------------------------
// printEntry
// ------------------------------------------
func printEntry(entry *Entry, totalNumberOfLines int, totalSize int64) {
	var numberOfLinesString = ""
	var percentageString = ""

	if !entry.IsBinary {
		numberOfLinesString = utils.IntToString(entry.NumberOfLines, thousandsSeparator)

		// Show percentage
		if totalNumberOfLines > 0 {
			var percentage = float64(entry.NumberOfLines) * float64(100) / float64(totalNumberOfLines)
			percentageString = fmt.Sprintf("%.1f", utils.Round(percentage, 1))
		}
	}

	var sizePercentage = float64(entry.Filesize) * float64(100) / float64(totalSize)
	var sizePercentageString = fmt.Sprintf("%.1f", utils.Round(sizePercentage, 1))

	fmt.Printf(
		formatString,
		entry.ExtensionName,
		utils.IntToString(entry.NumberOfFiles, thousandsSeparator),
		numberOfLinesString,
		percentageString,
		utils.Int64ToString(entry.Filesize, thousandsSeparator),
		sizePercentageString,
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
			"100.0",
			utils.Int64ToString(result.TotalSize, thousandsSeparator),
			"100.0",
		)
	}
}
