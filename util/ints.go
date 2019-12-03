package util

import (
	"bytes"
	"io"
	"strconv"
	"strings"
)

// IntsFromReader accepts any reader and returns the integers that is contained within it. Integers
// are separated by the given separator.
func IntsFromReader(reader io.Reader, sep string) ([]int, error) {
	buf := bytes.Buffer{}
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	var ints []int
	rows := strings.Split(string(buf.Bytes()), sep)
	for _, row := range rows {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}
		i, err := strconv.Atoi(row)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}
