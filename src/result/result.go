// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
 * result.go
 *
 */
package result

// Importing libraries
import (
	"entry"
)

/*
 * Result
 */
type Result struct {
	Directory          string
	Extensions         map[string]*entry.Entry
	TotalNumberOfFiles int
	TotalNumberOfLines int
	TotalSize          int64
}
