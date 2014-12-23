// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
 * utils.go
 *
 */
package utils

// Importing libraries
import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Constants
const (
	windowsNewline = "\r\n"
	unixNewline    = "\n"
	macNewline     = "\r"
)

// ------------------------------------------
// DetermineNewline
// ------------------------------------------
func DetermineNewline(stringWithNewline string) string {
	if strings.Contains(stringWithNewline, windowsNewline) {
		return windowsNewline
	}

	if strings.Contains(stringWithNewline, unixNewline) {
		return unixNewline
	}

	if strings.Contains(stringWithNewline, macNewline) {
		return macNewline
	}

	return windowsNewline
}

// ------------------------------------------
// IsInString
// ------------------------------------------
func IsInString(stringToSearch string, stringsToSearchFor []string) bool {
	var isFound = false

	for _, searchItem := range stringsToSearchFor {
		if strings.Contains(stringToSearch, searchItem) {
			isFound = true
			break
		}
	}

	return isFound
}

// ------------------------------------------
// IsInSlice
// ------------------------------------------
func IsInSlice(sliceToSearch []string, stringToSearchFor string) bool {
	var isFound = false

	for _, searchItem := range sliceToSearch {
		if searchItem == stringToSearchFor {
			isFound = true
			break
		}
	}

	return isFound
}

// ------------------------------------------
// IntToString
// ------------------------------------------
func IntToString(n int, separator rune) string {
	return Int64ToString(int64(n), separator)
}

// ------------------------------------------
// Int64ToString
// ------------------------------------------
func Int64ToString(n int64, separator rune) string {
	s := strconv.FormatInt(n, 10)
	startOffset := 0
	var buff bytes.Buffer

	if n < 0 {
		startOffset = 1
		buff.WriteByte('-')
	}

	length := len(s)
	commaIndex := 3 - ((length - startOffset) % 3)

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

// ------------------------------------------
// GetDirectory
// ------------------------------------------
func GetDirectory(pathFromFlag string, pathOfExecutable string) string {
	var err error

	// First non-flag argument should be the starting directory
	var path = pathFromFlag

	// If no directory given, use the current directory
	if len(path) == 0 {
		path = filepath.Dir(pathOfExecutable)
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
