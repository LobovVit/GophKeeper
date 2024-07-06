package repositories

import (
	"bytes"
	"os"
)

type Storage struct {
	dir string
}

func New(dir string) Storage {
	return Storage{
		dir: dir,
	}
}

func (s Storage) Store(file *File) error {
	if err := os.WriteFile(s.dir+file.name, file.buffer.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

type File struct {
	name   string
	buffer *bytes.Buffer
}

func (f *File) Write(chunk []byte) error {
	_, err := f.buffer.Write(chunk)
	return err
}
