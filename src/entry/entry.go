// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entry

type Entry struct {
	ExtensionName string
	IsBinary      bool
	NumberOfFiles int
	NumberOfLines int
	Filesize      int64
}
