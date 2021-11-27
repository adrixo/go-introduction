package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (m MyReader) Read(c []byte) (n int, err error) {
	c[0] = "A"
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
