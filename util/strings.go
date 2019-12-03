package util

import (
	"bytes"
	"io"
)

// StringFromReader accepts any reader and returns the string that is contained within it.
func StringFromReader(reader io.Reader) (string, error) {
	buf := bytes.Buffer{}
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}
