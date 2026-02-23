package internal

import "testing"

func TestPlaceholder(t *testing.T) {
	got := Placeholder()
	if got != "placeholder" {
		t.Errorf("Placeholder() = %q, want %q", got, "placeholder")
	}
}
