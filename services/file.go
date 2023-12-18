package services

import (
	"os"
	"strings"

	uuid "github.com/google/uuid"
)

type File struct {
	connection *os.File
	path       string
}

func NewFile(path string) *File {
	if path == "" {
		return nil
	}
	paths := strings.Split(path, "/")
	file_name := paths[len(paths)-1]
	if file_name == "" {
		file_name = uuid.NewString()
	}
	return &File{
		path:       file_name,
		connection: nil,
	}
}

func (f *File) Open() error {
	var err error
	f.connection, err = os.OpenFile(f.path, os.O_CREATE, 0666)
	return err

}

func (f *File) Write(b *[]byte) (int, error) {
	return f.connection.Write(*b)
}

func (f *File) Close() error {
	return f.connection.Close()
}
