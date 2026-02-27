package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type number interface {
	~int | ~int64 | ~int32 | ~float32 | ~float64 |
		~uint | ~uint32 | ~uint64
}

func sum[T number](a T, b T) T {
	return a + b
}

func process[T any](value T) {
	/*
		mitavan meghdar value ke az type T yani generic hast ro check konim key
		type oon key hast
	*/
	switch v := any(value).(type) {
	case int:
		fmt.Println("value T is int", v)
	case string:
		fmt.Println("value T is string", v)
	default:
		fmt.Println("Unknown type (only string or int)")
	}
}

func main() {
	// process(23)
	// fmt.Println(sum(23, 23.2))

	if err := processFile("test.txt"); err != nil {
		switch e := err.(type) {
		case *os.PathError:
			fmt.Println("Error os path:", e)
		case *ReadError:
			fmt.Println("Error read:", e)
		default:
			fmt.Println("unknown error:", e)
		}
	}
}

type ReadError struct {
	Filename string
	Op       string
}

func (re *ReadError) Error() string {
	return fmt.Sprintf("Error reading file %s: %s", re.Filename, re.Op)
}

func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return &os.PathError{
			Op:   "open",
			Path: filename,
			Err:  err,
		}
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	_, err = file.Read(buffer)
	if err != nil && !errors.Is(err, io.EOF) {
		return &ReadError{
			Op:       "read",
			Filename: filename,
		}
	}

	return nil
}
