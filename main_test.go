package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	if Print() != "Hello World" {
		t.Error("Print() doesn't return Hello World")
	}
}
