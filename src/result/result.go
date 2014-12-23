/*
 * result.go
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
