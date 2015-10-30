// hello_world_test.go
package main

import "testing"

func TestHello(t *testing.T) {
	exp := "hello go world"
	got := hello("Go")

	if exp != got {
		t.Errorf("\nExp: %s\nGot: %s", exp, got)
	}
}
