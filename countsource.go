// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MichaelTJones/walk"
)

var (
	root             = getDirectory(flag.Arg(0), ".")
	showDirectories  = flag.Bool("dir", false, "show exclusion status of directories in path.")
	showFiles        = flag.Bool("file", false, "show exclusion status of files in path.")
	showOnlyIncluded = flag.Bool("inc", false, "show only included files/directories.")
	showOnlyExcluded = flag.Bool("excl", false, "show only excluded files/directories.")
	showBigFiles     = flag.Int("big", 0, "show the x largest files")
	help             = flag.Bool("?", false, "this help information")
	config           Config
)

var countResult Result
var exclusions Exclusions
var pathSeparator = getPathSeparator()
var bigFiles = make(fileSizes, 0)
var configFilename string

// init
func init() {
	flag.Usage = usage
	flag.Parse()

	if *help {
		usage()
	}

	// Load config and prepare for parsing directory
	configFilename = getConfigFileName()
	config = loadConfig()
	exclusions = config.getExclusions()
	countResult = config.setupResult()
	printAnalyticsHeader()
}

// usage
func usage() {
	var executableName = filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "\nCOUNTSOURCE (C) Copyright 2014-2015 Erlend Johannessen\n")
	fmt.Fprintf(os.Stderr, "%s counts sourcecode lines for given directory and sub-directories.\n", executableName)
	fmt.Fprintf(os.Stderr, "\nUsage: %s [dirname] \n", executableName)
	fmt.Fprintf(os.Stderr, "  dirname: Name of directory with source code to count lines for. Uses current directory if no directory given.\n")
	flag.PrintDefaults()
	os.Exit(0)
}

// main
func main() {
	// Processing the given directory
	var err = walk.Walk(root, forEachFileSystemEntry)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}

	// Show result
	printResult()
}
