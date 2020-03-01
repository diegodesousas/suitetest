package server

import "io"

type Response struct {
	Status int
	Body   io.Reader
}
