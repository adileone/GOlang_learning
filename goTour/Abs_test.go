package main

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	point := Vertex{3, 4}
	result := point.Abs()

	if result != 5 {
		t.Error("Abs should have been 5 instead of " + fmt.Sprintf("%f", result))
	}
}
