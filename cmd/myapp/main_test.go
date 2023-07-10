package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"
)

func TestHelloHandler(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	// Define test cases
	testCases := []struct {
		path     string
		expected string
	}{
		{path: "World", expected: "Hello, World!"},
		{path: "Gopher", expected: "Hello, Gopher!"},
	}

	for _, tc := range testCases {
		resp, err := http.Get(server.URL + "/" + tc.path)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status OK; got %v", resp.Status)
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		body := strings.TrimSpace(string(bodyBytes))

		if body != tc.expected {
			t.Fatalf("Expected %s, but got %s", tc.expected, body)
		}
	}
}
