package files

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

type Local struct {
	maxFileSize int
	basePath    string
}

func NewLocal(basePath string, maxSize int) (*Local, error) {
	path, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	return &Local{basePath: path}, nil
}

func (l *Local) Save(path string, contents io.Reader) error {
	localFilePath := l.fullPath(path)

	dir := filepath.Dir(localFilePath)
	err := os.MkdirAll(dir, os.ModeDir)
	if err != nil {
		return xerrors.Errorf("Unable to create directory: %w", err)
	}

	_, err = os.Stat(localFilePath)
	if err == nil {
		err = os.Remove(localFilePath)
		if err != nil {
			return xerrors.Errorf("Unable to delete file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return xerrors.Errorf("Unable to get file info: %w", err)
	}

	file, err := os.Create(localFilePath)
	if err != nil {
		return xerrors.Errorf("Unable to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, contents)
	if err != nil {
		return xerrors.Errorf("Unable to write to file: %w", err)
	}

	return nil
}

func (l *Local) Get(path string) (*os.File, error) {
	filepath := l.fullPath(path)

	file, err := os.Open(filepath)
	if err != nil {
		return nil, xerrors.Errorf("Unable to open file: %w", err)
	}

	return file, nil
}

func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basePath, path)
}
