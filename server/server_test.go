package server_test

import (
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/diegodesousas/testsuite/server"
)

var _ = Describe("Server", func() {
	var ms *server.MockServer

	BeforeEach(func() {
		ms = server.NewMockServer()
		ms.SetRequestPattern([]server.RequestPattern{
			{
				Path:   "/test",
				Method: http.MethodGet,
				Body:   nil,
				Response: server.Response{
					Status: http.StatusOK,
					Body:   server.ReadString(`{"test": "suite"}`),
				},
			},
			{
				Path:   "/test",
				Method: http.MethodPost,
				Body:   server.ReadString(`{"test": "suite"}`),
				Response: server.Response{
					Status: http.StatusOK,
					Body:   server.ReadString(`{"test": "suite", "method": "post"}`),
				},
			},
			{
				Path:   "/test/empty-body",
				Method: http.MethodPost,
				Body:   server.ReadString(`{"empty": "body"}`),
				Response: server.Response{
					Status: http.StatusOK,
					Body:   nil,
				},
			},
		})
	})

	It("Should response properly based on params", func() {
		path := ms.Url() + "/test"
		resp, _ := http.Get(path)
		expectedBody := `{"test": "suite"}`

		bytes, _ := ioutil.ReadAll(resp.Body)

		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(string(bytes)).To(Equal(expectedBody))
	})

	It("Should response when pass body param", func() {
		path := ms.Url() + "/test"
		body := server.ReadString(`{"test": "suite"}`)
		resp, _ := http.Post(path, "application/json", body)
		expectedBody := `{"test": "suite", "method": "post"}`

		bytes, _ := ioutil.ReadAll(resp.Body)

		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(string(bytes)).To(Equal(expectedBody))
	})

	It("Should response when pass body param", func() {
		path := ms.Url() + "/test/empty-body"
		body := server.ReadString(`{"empty": "body"}`)
		resp, _ := http.Post(path, "application/json", body)
		expectedBody := ""

		bytes, _ := ioutil.ReadAll(resp.Body)

		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(string(bytes)).To(Equal(expectedBody))
	})

})
