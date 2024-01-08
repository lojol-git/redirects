package redirects

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRun(t *testing.T) {
	// Load the redirects
	r := Load("redirects_test.yml")
	if r != nil {
		t.Fatalf("Load was incorrect, got: %s, want: nil.", r.Error())
	}

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/old-path/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the Run function
	result := Run(rr, req)

	// Check the result
	if !result {
		t.Error("Run was incorrect, got: false, want: true.")
	}
}
