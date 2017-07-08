// Copyright 2014-2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

// fileSize contains file size
type fileSize struct {
	Name  string
	Size  int64
	Lines int
}

// FileSizes contains a slice of FileSize
type fileSizes []fileSize

//func (p FileSizes) Add(name string, size int64) { p = append(p, FileSize{name, size}) }
func (p fileSizes) Len() int           { return len(p) }
func (p fileSizes) Less(i, j int) bool { return p[i].Lines > p[j].Lines }
func (p fileSizes) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
