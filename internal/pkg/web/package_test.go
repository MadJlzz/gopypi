package web

import "testing"

func TestNormalize(t *testing.T) {
	uri := "https://frodo.lotr.dev/bad_path-should.be_normalized/"
	got := normalize(uri)
	want := "https://frodo-lotr-dev/bad-path-should-be-normalized/"
	if got != want {
		t.Errorf("normalize uri doesn't follow PEP503 spec.\ngot: [%s], want: [%s]", got, want)
	}
}
