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

	"utils"
)

// Exclusions
type Exclusions struct {
	ExcludeDirectories []string
	ExcludeFiles       []string
}

var countResult Result
var exclusions Exclusions
var rootPath = ""
var pathSeparator = "/"
var showDirectoryStatus = false
var showFileStatus = false
var showOnlyIncluded = true
var showOnlyExcluded = true
var showBigFiles = 0
var bigFiles = make(FileSizes, 0)

type FileSize struct {
	Name  string
	Size  int64
	Lines int
}
type FileSizes []FileSize

//func (p FileSizes) Add(name string, size int64) { p = append(p, FileSize{name, size}) }
func (p FileSizes) Len() int           { return len(p) }
func (p FileSizes) Less(i, j int) bool { return p[i].Lines > p[j].Lines }
func (p FileSizes) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// ------------------------------------------
// Initialize
// ------------------------------------------
func Initialize(root string, showDirectories, showFiles, showOnlyInc, showOnlyExcl bool, showBig int) {
	rootPath = root
	showDirectoryStatus = showDirectories
	showFileStatus = showFiles

	showOnlyIncluded = showOnlyInc
	showOnlyExcluded = showOnlyExcl
	showBigFiles = showBig

	pathSeparator = utils.GetPathSeparator()
	var sc = LoadConfig()

	exclusions = setupExclusions(sc)
	countResult = setupResult(sc)
}

// ------------------------------------------
// setupExclusions
// ------------------------------------------
func setupExclusions(sc Config) Exclusions {
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
func setupResult(sc Config) Result {
	var r = Result{Extensions: make(map[string]*Entry)}

	// Add extensions
	for _, k := range sc.Extensions {
		r.Extensions[k] = &Entry{k, false, 0, 0, 0}
	}

	// Add binary extensions
	for _, k := range sc.BinaryExtensions {
		r.Extensions[k] = &Entry{k, true, 0, 0, 0}
	}

	return r
}

// ------------------------------------------
// isExcluded
// ------------------------------------------
func isExcluded(filename string) bool {
	// Get full path of file
	var fulldir, _ = filepath.Abs(filename)

	var excluded = utils.IsInString(fulldir+pathSeparator, exclusions.ExcludeDirectories)

	if !excluded {
		excluded = utils.IsInSlice(exclusions.ExcludeFiles, filepath.Base(filename))
	}

	return excluded
}

// ------------------------------------------
// ShowDirectoryOrFile
// ------------------------------------------
func ShowDirectoryOrFile(isDir bool, filename string, excluded bool) {
	var status = ""

	if showDirectoryStatus && isDir {
		if excluded {
			status = " EXCLUDED"
		} else {
			status = ""
		}

		if (showOnlyIncluded && !excluded) || (showOnlyExcluded && excluded) || (!showOnlyIncluded && !showOnlyExcluded) {
			fmt.Printf("Directory %s%s\n", strings.Replace(filename, rootPath+pathSeparator, "", 1), status)
		}
	}

	if showFileStatus && !isDir {
		var indent = "   "
		if !showDirectoryStatus {
			indent = "File "
		}

		if excluded {
			status = " EXCLUDED"
		} else {
			status = ""
		}

		if (showOnlyIncluded && !excluded) || (showOnlyExcluded && excluded) || (!showOnlyIncluded && !showOnlyExcluded) {
			fmt.Printf("%s %s%s\n", indent, strings.Replace(filename, rootPath+pathSeparator, "", 1), status)
		}
	}
}

// ------------------------------------------
// CountExtension
// ------------------------------------------
func CountExtension(filename string, f os.FileInfo) {
	if f == nil {
		return
	}

	// Default excluded if it is a directory
	// If not, check for exclusions
	//var excluded = f.IsDir() || isExcluded(filename)
	var excluded = isExcluded(filename)

	ShowDirectoryOrFile(f.IsDir(), filename, excluded)

	if !f.IsDir() && !excluded {
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
					bigFiles = append(bigFiles, FileSize{f.Name(), size, numberOfLines}) // filename
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
func Print() {
	PrintResult(rootPath, countResult, bigFiles)
}
