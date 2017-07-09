// Copyright 2014-2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Exclusions contains the excluded directories and files
type Exclusions struct {
	ExcludeDirectories []string
	ExcludeFiles       []string
}

// forEachFileSystemEntry
func forEachFileSystemEntry(filename string, f os.FileInfo, err error) error {
	countExtension(filename, f)
	return nil
}

// isExcluded
func isExcluded(filename string) bool {
	// Get full path of file
	var fulldir, _ = filepath.Abs(filename)

	var excluded = isInString(fulldir+pathSeparator, exclusions.ExcludeDirectories)

	if !excluded {
		excluded = isInSlice(exclusions.ExcludeFiles, filepath.Base(filename))
	}

	return excluded
}

// showDirectoriesOrFile
func showDirectoriesOrFile(isDir bool, filename string, excluded bool) {
	var status string

	if *showDirectories && isDir {
		if excluded {
			status = " EXCLUDED"
		} else {
			status = ""
		}

		if (*showOnlyIncluded && !excluded) || (*showOnlyExcluded && excluded) || (!*showOnlyIncluded && !*showOnlyExcluded) {
			fmt.Printf("Directory %s%s\n", strings.Replace(filename, root+pathSeparator, "", 1), status)
		}
	}

	if *showFiles && !isDir {
		var indent = "   "
		if !*showDirectories {
			indent = "File "
		}

		if excluded {
			status = " EXCLUDED"
		} else {
			status = ""
		}

		if (*showOnlyIncluded && !excluded) || (*showOnlyExcluded && excluded) || (!*showOnlyIncluded && !*showOnlyExcluded) {
			fmt.Printf("%s %s%s\n", indent, strings.Replace(filename, root+pathSeparator, "", 1), status)
		}
	}
}

// countExtension
func countExtension(filename string, f os.FileInfo) {
	if f == nil {
		return
	}

	// Default excluded if it is a directory
	// If not, check for exclusions
	//var excluded = f.IsDir() || isExcluded(filename)
	var excluded = isExcluded(filename)

	showDirectoriesOrFile(f.IsDir(), filename, excluded)

	if !f.IsDir() && !excluded {
		// Extension for the entry we're looking at
		var ext = filepath.Ext(filename)

		// Is the extension one of the relevant ones?
		var _, willBeCounted = countResult.Extensions[ext]

		// If yes, proceed with counting
		if willBeCounted {
			countResult.Extensions[ext].NumberOfFiles++
			countResult.TotalNumberOfFiles++

			var size = f.Size()
			countResult.Extensions[ext].Filesize += size
			countResult.TotalSize += size

			// Slurp the whole file into memory
			var contents, err = ioutil.ReadFile(filename)

			if err != nil {
				fmt.Printf("Problem reading inputfile %s, error:%v\n", filename, err)
				return
			}

			var isBinary = isBinaryFormat(contents)

			// Binary files will not have "number of lines"
			// but might need to have the binary flag set
			if isBinary && !countResult.Extensions[ext].IsBinary {
				countResult.Extensions[ext].IsBinary = true
			} else {
				var stringContents = string(contents)
				var newline = determineNewline(stringContents)

				var numberOfLines = len(strings.Split(stringContents, newline))
				countResult.Extensions[ext].NumberOfLines += numberOfLines
				countResult.TotalNumberOfLines += numberOfLines
				bigFiles = append(bigFiles, fileSize{f.Name(), size, numberOfLines})
			}
		}
	}
}
