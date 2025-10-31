package formatter

import "testing"

func TestNew(t *testing.T) {
	f := New()
	if f == nil {
		t.Fatal("New() returned nil")
	}
}
