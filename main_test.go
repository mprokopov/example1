package main

import "testing"

func TestSimpleFactory(t *testing.T) {
	f := SimpleFactory("http://localhost1")

	if f.Url != "http://localhost" {
		t.Errorf("feature incorrect, got %s, want: %s ", f.Url, "http://localhost")
	}
}
