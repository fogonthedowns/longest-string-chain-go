package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	in := []string{"a", "b", "ba", "bca", "bda", "bdca"}
	out := longestStrChain(in)
	if out != 4 {
		t.Errorf("got %d, want %d", out, 4)
	}

}
