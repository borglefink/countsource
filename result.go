// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

// ResultEntry is the file entry for result
type ResultEntry struct {
	ExtensionName string
	IsBinary      bool
	NumberOfFiles int
	NumberOfLines int
	Filesize      int64
}

// Result is metadata for the result
type Result struct {
	Directory          string
	Extensions         map[string]*ResultEntry
	TotalNumberOfFiles int
	TotalNumberOfLines int
	TotalSize          int64
}
