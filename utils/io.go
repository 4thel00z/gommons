package utils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type LineIterator struct {
	reader   *bufio.Reader
	capacity int64
}

func NewLineIteratorFromPath(path string, capacity int) *LineIterator {
	file, err := os.Open(path)
	Cry(err)
	defer file.Close()
	size, err := GetFileSize(path)
	Cry(err)

	return &LineIterator{
		reader:   bufio.NewReader(file),
		capacity: size,
	}
}

func NewLineIterator(reader io.Reader, capacity int64) *LineIterator {
	return &LineIterator{
		reader:   bufio.NewReader(reader),
		capacity: capacity,
	}
}

func (lineIterator *LineIterator) Next() ([]byte, error) {
	var bytes []byte
	for {
		line, isPrefix, err := lineIterator.reader.ReadLine()
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, line...)
		if !isPrefix {
			break
		}
	}
	return bytes, nil
}

func (lineIterator *LineIterator) ReadUntilEOF() string {
	var builder strings.Builder
	builder.Grow(int(lineIterator.capacity))
	for {
		line, err := lineIterator.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				Cry(err)
			}
		}
		builder.Write(line)
	}
	return builder.String()
}

func GetFileSize(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return -1, err
	}
	return fi.Size(), nil
}

func ReadStdin() ([]byte, error) {
	return ioutil.ReadAll(os.Stdin)
}
