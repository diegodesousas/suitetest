package server

import "io"

type RequestPattern struct {
	Path     string
	Method   string
	Body     io.Reader
	Response Response
}
