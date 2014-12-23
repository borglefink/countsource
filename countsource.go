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

// ------------------------------------------
// init
// ------------------------------------------
func init() {
	flag.Usage = usage
	flag.Parse()

	// Load config and prepare for parsing directory
	count.Initialize()
}

// ------------------------------------------
// usage
// ------------------------------------------
func usage() {
	var executableName = filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "%s counts sourcecode lines for given directory and sub-directories.\n", executableName)
	fmt.Fprintf(os.Stderr, "Usage: %s [dirname] \n", executableName)
	fmt.Fprintf(os.Stderr, "       dirname: Name of directory with source code to count lines for.\n")
	fmt.Fprintf(os.Stderr, "                Uses current directory if no directory given.\n")
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
	root = utils.GetDirectory(flag.Arg(0), os.Args[0])

	// Walking the given directory
	err := filepath.Walk(root, forEachEntry)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}

	// Show result
	count.Print(root)
}
