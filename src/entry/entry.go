/*
 * entry.go
 *
 * Resultcontainer for extensions
 */
package entry

// Entry for each extension, counting
type Entry struct {
	ExtensionName string
	IsBinary      bool
	NumberOfFiles int
	NumberOfLines int
	Filesize      int64
}
