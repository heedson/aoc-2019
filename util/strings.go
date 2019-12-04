package util

import (
	"bytes"
	"io"
	"strings"
)

// StringsFromReader accepts any reader and returns the strings that is contained within it. Strings
// are separated by the given separator.
func StringsFromReader(reader io.Reader, sep string) ([]string, error) {
	buf := bytes.Buffer{}
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	var strs []string
	segments := strings.Split(string(buf.Bytes()), sep)
	for _, seg := range segments {
		if len(seg) == 0 {
			continue
		}
		strs = append(strs, seg)
	}
	return strs, nil
}
