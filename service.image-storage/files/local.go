package files

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

// Local struct represents a local storage
type Local struct {
	maxFileSize int
	basePath    string
}

// NewLocal creates a new local struct with given basePath and maxSize of
// the files
func NewLocal(basePath string, maxSize int) (*Local, error) {
	path, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	return &Local{basePath: path}, nil
}

// Save is implemented from the Storage interface and it will save contents to
// a specific file.
func (l *Local) Save(path string, contents io.Reader) error {
	localFilePath := l.fullPath(path)

	// check dir of localFilePath and create it
	dir := filepath.Dir(localFilePath)
	err := os.MkdirAll(dir, os.ModeDir)
	if err != nil {
		return xerrors.Errorf("Unable to create directory: %w", err)
	}

	// get the file info of the localFilePath
	_, err = os.Stat(localFilePath)
	if err == nil {
		// try to remove the file for saving later
		err = os.Remove(localFilePath)
		if err != nil {
			return xerrors.Errorf("Unable to delete file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return xerrors.Errorf("Unable to get file info: %w", err)
	}

	// create new file with localFilePath
	file, err := os.Create(localFilePath)
	if err != nil {
		return xerrors.Errorf("Unable to create file: %w", err)
	}
	defer file.Close()

	// copy the given contents into the created file
	_, err = io.Copy(file, contents)
	if err != nil {
		return xerrors.Errorf("Unable to write to file: %w", err)
	}

	return nil
}

// Get returns contents of a specific path or file
func (l *Local) Get(path string) (*os.File, error) {
	filepath := l.fullPath(path)

	// open file and catch error
	file, err := os.Open(filepath)
	if err != nil {
		return nil, xerrors.Errorf("Unable to open file: %w", err)
	}

	return file, nil
}

func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basePath, path)
}
