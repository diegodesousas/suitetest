package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
)

type MockServer struct {
	server  *httptest.Server
	handler *mockHandler
}

type mockHandler struct {
	patterns []RequestPattern
}

func (mh mockHandler) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal("testsuite: unable to read request body")
		return
	}

	for _, rp := range mh.patterns {
		if rp.Body == nil {
			rp.Body = ioutil.NopCloser(ReadString(""))
		}

		expectedBody, _ := ioutil.ReadAll(rp.Body)
		if reflect.DeepEqual(bodyBytes, expectedBody) && r.URL.String() == rp.Path && r.Method == rp.Method {
			if rp.Response.Body == nil {
				rp.Response.Body = ioutil.NopCloser(ReadString(""))
			}

			resBytes, _ := ioutil.ReadAll(rp.Response.Body)
			res.WriteHeader(rp.Response.Status)
			_, _ = res.Write(resBytes)
			return
		}
	}

	log.Fatal("testsuite: pattern not found in mock server")
}

func (ms MockServer) Url() string {
	return ms.server.URL
}

func (ms MockServer) SetRequestPattern(patterns []RequestPattern) *MockServer {
	ms.handler.patterns = patterns
	return &ms
}

func NewMockServer() *MockServer {
	mh := &mockHandler{}
	return &MockServer{
		server:  httptest.NewServer(mh),
		handler: mh,
	}
}
