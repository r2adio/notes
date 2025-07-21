package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHandler(t *testing.T) {
	e := echo.New() //clean echo instance for testing, httptest.NewRequest():creates fake http req

	req := httptest.NewRequest(http.MethodGet, "/", nil) //meathod: get, url:"/", body: nil

	resp := httptest.NewRecorder() //captures response; records: status code, headers, body; no network:everything in memory

	c := e.NewContext(req, resp) //creates echo context; simulates real req handlin' env; no server startup needed

	s := &Server{} //creates server instance

	// Handler Execution and Error Checking
	if err := s.HelloWorldHandler(c); err != nil {
		t.Errorf("handler() error = %v", err)
		return
	}
	// Status Code Assertions
	if resp.Code != http.StatusOK {
		t.Errorf("handler() wrong status code = %v", resp.Code)
		return
	}

	//steps: define expected -> decode actual -> deep comparison -> detailed error
	// Response Body Testing
	expected := map[string]string{"message": "Hello World"}
	var actual map[string]string
	// Decode the response body into the actual map
	//json.NewDecoder?: stream processing; type safety; error handling
	if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
		t.Errorf("handler() error decoding response body: %v", err)
		return
	}
	// Compare the decoded response with the expected value
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("handler() wrong response body. expected = %v, actual = %v", expected, actual)
		return
	}
}
