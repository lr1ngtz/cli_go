package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("w1 w2 w3 w4\n")
	exp := 4
	res := count(b, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}


func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("w1 w2 w3\nl2\nl4 w1")
	exp := 3
	res := count(b, true)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}