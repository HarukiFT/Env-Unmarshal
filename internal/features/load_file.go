package features

import (
	"os"
	"path/filepath"
)

func LoadFileWD(path string) (*os.File, error) {
	wd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(wd, path))

	return file, err
}
