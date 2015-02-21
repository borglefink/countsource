// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"count"
	"utils"
)

var root string
var showDirectories *bool
var showFiles *bool
var showOnlyIncluded *bool
var showOnlyExcluded *bool

// ------------------------------------------
// init
// ------------------------------------------
func init() {
	showDirectories = flag.Bool("dir", false, "show exclusion status of directories in path.")
	showFiles = flag.Bool("file", false, "show exclusion status of files in path.")
	showOnlyIncluded = flag.Bool("inc", false, "show only included files/directories.")
	showOnlyExcluded = flag.Bool("excl", false, "show only excluded files/directories.")

	var help = flag.Bool("?", false, "this help information")

	flag.Usage = usage
	flag.Parse()

	if *help {
		usage()
	}

	count.PrintAnalyticsHeader(*showDirectories, *showFiles, *showOnlyIncluded, *showOnlyExcluded)
}

// ------------------------------------------
// usage
// ------------------------------------------
func usage() {
	var executableName = filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "\nCOUNTSOURCE (C) Copyright 2014 Erlend Johannessen\n")
	fmt.Fprintf(os.Stderr, "%s counts sourcecode lines for given directory and sub-directories.\n", executableName)
	fmt.Fprintf(os.Stderr, "\nUsage: %s [dirname] \n", executableName)
	fmt.Fprintf(os.Stderr, "  dirname: Name of directory with source code to count lines for. Uses current directory if no directory given.\n")
	flag.PrintDefaults()
	os.Exit(0)
}

// ------------------------------------------
// forEachEntry
// ------------------------------------------
func forEachEntry(filename string, f os.FileInfo, err error) error {
	count.CountExtension(filename, f)
	return nil
}

// ------------------------------------------
// main
// ------------------------------------------
func main() {
	root = utils.GetDirectory(flag.Arg(0), ".")

	// Load config and prepare for parsing directory
	count.Initialize(root, *showDirectories, *showFiles, *showOnlyIncluded, *showOnlyExcluded)

	// Processing the given directory
	var err = filepath.Walk(root, forEachEntry)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}

	// Show result
	count.Print()
}
