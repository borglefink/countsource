// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package count

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"config"
	"result"
	"utils"
)

const (
	pathSeparator = "\\"
)

// Exclusions
type Exclusions struct {
	ExcludeDirectories []string
	ExcludeFiles       []string
}

var countResult result.Result
var exclusions Exclusions

// ------------------------------------------
// setupExclusions
// ------------------------------------------
func setupExclusions(sc config.Config) Exclusions {
	var exclusions = Exclusions{sc.ExcludeDirectories, sc.ExcludeFiles}

	// Initialise ExcludeDirectories with searchable criteria
	for index, _ := range exclusions.ExcludeDirectories {
		exclusions.ExcludeDirectories[index] = pathSeparator + exclusions.ExcludeDirectories[index] + pathSeparator
	}

	return exclusions
}

// ------------------------------------------
// setupResult
// ------------------------------------------
func setupResult(sc config.Config) result.Result {
	var r = result.Result{Extensions: make(map[string]*result.Entry)}

	// Add extensions
	for _, k := range sc.Extensions {
		r.Extensions[k] = &result.Entry{k, false, 0, 0, 0}
	}

	// Add binary extensions
	for _, k := range sc.BinaryExtensions {
		r.Extensions[k] = &result.Entry{k, true, 0, 0, 0}
	}

	return r
}

// ------------------------------------------
// isExcluded
// ------------------------------------------
func isExcluded(filename string) bool {
	// Get full path of file
	var fulldir, _ = filepath.Abs(filename)

	var excluded = utils.IsInString(fulldir, exclusions.ExcludeDirectories)

	if !excluded {
		excluded = utils.IsInSlice(exclusions.ExcludeFiles, filename)
	}

	return excluded
}

// ------------------------------------------
// Initialize
// ------------------------------------------
func Initialize() {
	var sc = config.LoadConfig()

	exclusions = setupExclusions(sc)
	countResult = setupResult(sc)
}

// ------------------------------------------
// CountExtension
// ------------------------------------------
func CountExtension(filename string, f os.FileInfo) {
	// Default excluded if it is a directory
	// If not, check for exclusions
	var excluded = f.IsDir() || isExcluded(filename)

	if !excluded {
		// Extension for the entry we're looking at
		var ext = filepath.Ext(filename)

		// Is the extension one of the relevant ones?
		var _, willBeCounted = countResult.Extensions[ext]

		// If yes, proceed with counting
		if willBeCounted {
			countResult.Extensions[ext].NumberOfFiles += 1
			countResult.TotalNumberOfFiles += 1

			var size = f.Size()
			countResult.Extensions[ext].Filesize += size
			countResult.TotalSize += size

			// Binary files will not have "number of lines"
			if !countResult.Extensions[ext].IsBinary {
				// Slurp the whole file into memory
				var contents, err = ioutil.ReadFile(filename)

				// Ok, count lines
				if err == nil {
					var stringContents = string(contents)
					var newline = utils.DetermineNewline(stringContents)

					var numberOfLines = len(strings.Split(stringContents, newline))
					countResult.Extensions[ext].NumberOfLines += numberOfLines
					countResult.TotalNumberOfLines += numberOfLines
				} else {
					fmt.Println("Problem reading inputfile %s, error:%v", filename, err)
				}
			}
		}
	}
}

// ------------------------------------------
// Print
// ------------------------------------------
func Print(root string) {
	result.PrintResult(root, countResult)
}
