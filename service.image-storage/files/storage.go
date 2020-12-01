package files

import "io"

// Storage interface will represent a storage for this application
type Storage interface {
	Save(path string, file io.Reader) error
}
