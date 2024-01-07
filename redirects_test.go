package redirects

import (
	"testing"
)

func TestLoad(t *testing.T) {
	err := Load("redirects_test.yml")
	if err != nil {
		t.Errorf("Load was incorrect, got: %s, want: nil.", err.Error())
	}
}
