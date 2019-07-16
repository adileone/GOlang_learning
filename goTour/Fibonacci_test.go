package main

import (
	"testing"
)

//TestFibonacci : closure func test
func TestFibonacci(t *testing.T) {
	result := 34
	f := Fibonacci()
	var num int

	for i := 0; i < 10; i++ {
		num = f()
	}
	if num != result {
		t.Error("Errore Logico")
	}
}
