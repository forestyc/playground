package utils

import (
	"os"
	"path"
)

type File struct {
	f        *os.File
	filename string
}

func NewFile(filename string) *File {
	dumpPath := path.Dir(filename)
	os.MkdirAll(dumpPath, os.ModePerm)
	return &File{
		filename: filename,
	}
}

// SaveFile saves buf to a file
func (f *File) SaveFile(buf []byte) error {
	var err error
	f.f, err = os.Create(f.filename)
	if err != nil {
		return err
	}
	defer f.f.Close()
	_, err = f.f.Write(buf)
	return err
}

// LoadFile reads all content from file
func (f *File) LoadFile() ([]byte, error) {
	return os.ReadFile(f.filename)
}

// AppendFile appends buf to a file
func (f *File) AppendFile(buf []byte) error {
	var err error
	if f != nil {
		f.f, err = os.OpenFile(f.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}
	if err != nil {
		return err
	}
	_, err = f.f.Write(buf)
	return err
}

func (f *File) Close() {
	f.f.Close()
}
