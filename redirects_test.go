package redirects

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRun(t *testing.T) {
	// Load the redirects
	err := Load("redirects_test.yml")
	if err != nil {
		t.Fatalf("Load was incorrect, got: %s, want: nil.", err.Error())
	}

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/old-path", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a next handler that should not be called
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Next handler should not be called")
	})

	// Create the handler and serve
	handler := Run(next)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusMovedPermanently {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMovedPermanently)
	}

	// Check the redirect location
	expected := "/new-path"
	if location := rr.Header().Get("Location"); location != expected {
		t.Errorf("handler returned unexpected location: got %v want %v",
			location, expected)
	}
}
