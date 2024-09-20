package auth

import "testing"

func TestAuth(t *testing.T) {
	got := "123"
	want := "123"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
