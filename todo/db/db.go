package db

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	suffix = ".todo"
)

// DB is file base key/value database
type DB struct {
	rootDir string
}

// New returns a new database
func New(rootDir string) (*DB, error) {
	if !isDir(rootDir) {
		err := os.MkdirAll(rootDir, 0700)
		if err != nil {
			return nil, err
		}
	}
	return &DB{rootDir}, nil
}

// Keys returns all the keys in the databases
func (db *DB) Keys() ([]string, error) {
	pattern := db.filePath("*")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	keys := make([]string, len(files))
	for i, fileName := range files {
		name := filepath.Base(fileName)
		keys[i] = name[:len(name)-len(suffix)]
	}

	return keys, nil
}

// Save saves data to key
func (db *DB) Save(key string, data []byte) error {
	file, err := os.Create(db.filePath(key))
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}

// Load loads data for key
func (db *DB) Load(key string) ([]byte, error) {
	return ioutil.ReadFile(db.filePath(key))
}

// Delete deletes key
func (db *DB) Delete(key string) error {
	return os.Remove(db.filePath(key))
}

func (db *DB) filePath(key string) string {
	return filepath.Join(db.rootDir, key+suffix)
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
