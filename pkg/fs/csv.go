package fs

import (
	"errors"
	"os"

	"github.com/gocarina/gocsv"
)

type Csv[T any] struct {
	Path string
}

func NewCsv[T any](filePath string) *Csv[T] {
	return &Csv[T]{
		Path: filePath,
	}
}

func (f *Csv[T]) Read() (T, error) {
	file, err := os.OpenFile(f.Path, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		var zero T
		return zero, err
	}
	defer file.Close()

	var data T
	if err := gocsv.UnmarshalFile(file, &data); err != nil && !errors.Is(err, gocsv.ErrEmptyCSVFile) {
		var zero T
		return zero, err
	}

	return data, nil
}

func (f *Csv[T]) Write(data T) error {
	file, err := os.OpenFile(f.Path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return gocsv.MarshalFile(&data, file)
}
