package io_ext

import (
	"io"
	"io/ioutil"
)

func WriterFromReader(writer io.Writer, reader io.Reader) (n int, err error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return len(b), err
	}

	return writer.Write(b)
}
