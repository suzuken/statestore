package statestore

import (
	"encoding/gob"
	"os"
)

// ReadWriter is interface for storing state.
type ReadWriter interface {
	Write(interface{}) error
	Read(interface{}) error
}

// FileStateStore is store for state in a file.
type FileStateStore struct {
	filename string
}

// NewFileStateStore creates FileStateStore for saving state.
// You must give a filepath to save.
func NewFileStateStore(path string) *FileStateStore {
	return &FileStateStore{filename: path}
}

// Write writes state to file. Overwrite old state.
// In FileStateStore, saving state as gob.
func (fs *FileStateStore) Write(e interface{}) error {
	f, err := os.Create(fs.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	return enc.Encode(e)
}

// Read reads state from file.
func (fs *FileStateStore) Read(e interface{}) error {
	finfo, err := os.Stat(fs.filename)
	if err != nil {
		// if file not found, create new one.
		ff, err := os.Create(fs.filename)
		if err != nil {
			return err
		}
		defer ff.Close()
		return nil
	}
	// file is empty
	if finfo.Size() == 0 {
		return nil
	}
	f, err := os.Open(fs.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	dec := gob.NewDecoder(f)
	return dec.Decode(e)
}
