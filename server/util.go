package server

import (
	"bytes"
)

func ReadString(s string) *bytes.Reader {
	jsonPattern := []byte(s)
	return bytes.NewReader(jsonPattern)
}
