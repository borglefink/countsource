// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	windowsNewline       = "\r\n"
	unixNewline          = "\n"
	oldMacNewline        = "\r"
	unixPathSeparator    = "/"
	windowsPathSeparator = "\\"
)

// determineNewline
func determineNewline(stringWithNewline string) string {
	if strings.Contains(stringWithNewline, windowsNewline) {
		return windowsNewline
	}

	if strings.Contains(stringWithNewline, unixNewline) {
		return unixNewline
	}

	if strings.Contains(stringWithNewline, oldMacNewline) {
		return oldMacNewline
	}

	return windowsNewline
}

// isInString
func isInString(stringToSearch string, stringsToSearchFor []string) bool {
	var isFound = false

	for _, searchItem := range stringsToSearchFor {
		if strings.Contains(stringToSearch, searchItem) {
			isFound = true
			break
		}
	}

	return isFound
}

// isInSlice
func isInSlice(sliceToSearch []string, stringToSearchFor string) bool {
	var isFound = false

	for _, searchItem := range sliceToSearch {
		if searchItem == stringToSearchFor {
			isFound = true
			break
		}
	}

	return isFound
}

// intToString
func intToString(n int, separator rune) string {
	return int64ToString(int64(n), separator)
}

// int64ToString
func int64ToString(n int64, separator rune) string {
	var s = strconv.FormatInt(n, 10)
	var startOffset = 0
	var buff bytes.Buffer

	if n < 0 {
		startOffset = 1
		buff.WriteByte('-')
	}

	var length = len(s)
	var commaIndex = 3 - ((length - startOffset) % 3)

	if commaIndex == 3 {
		commaIndex = 0
	}

	for i := startOffset; i < length; i++ {
		if commaIndex == 3 {
			buff.WriteRune(separator)
			commaIndex = 0
		}
		commaIndex++
		buff.WriteByte(s[i])
	}

	return buff.String()
}

// round return rounded version of x with prec precision.
// http://grokbase.com/t/gg/golang-nuts/12ag1s0t5y/go-nuts-re-function-round
func round(x float64, prec int) float64 {
	var rounder float64
	var pow = math.Pow(10, float64(prec))
	var intermed = x * pow
	var _, frac = math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

// getDirectory
func getDirectory(pathFromFlag, defaultPath string) string {
	var err error

	// First non-flag argument should be the starting directory
	var path = pathFromFlag

	// If no directory given, use the current directory
	if len(path) == 0 {
		path = filepath.Dir(defaultPath)
	}

	// Getting the full path, if necessary
	path, err = filepath.Abs(path)
	if err != nil {
		fmt.Printf("Directory [%v] does not exist.\n", path)
		os.Exit(-1)
	}

	// Removing quotes, if any
	path = strings.Replace(path, "\"", "", -1)

	// Checking if directory is ok
	_, err = os.Stat(path)
	if err != nil {
		fmt.Printf("Directory [%v] does not exist.\n", path)
		os.Exit(-1)
	}

	return path
}

// isWindows
func isWindows() bool {
	return strings.Index(os.Getenv("OS"), "Windows") >= 0
}

// getPathSeparator
func getPathSeparator() string {
	if isWindows() {
		return windowsPathSeparator
	}
	return unixPathSeparator
}
