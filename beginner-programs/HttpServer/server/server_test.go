package server

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
)

func  TestMyRouterAndHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hello/nazar", nil)
	res := httptest.NewRecorder()
	NewServer().ServeHTTP(res, req)

	expectedResponse := "Hello, nazar"
	errorMessage := fmt.Sprintf("Expected %s but got ", expectedResponse)

	if res.Body.String() != expectedResponse {
		t.Errorf(errorMessage + res.Body.String())
	}
}