package main

import (
	"reflect"
	"testing"
)

//TestWordCount : Test
func TestWordCount(t *testing.T) {

	m := map[string]int{"Ciao": 1, "Mondo": 1}

	mappa := WordCount("Ciao Mondo")

	eq := reflect.DeepEqual(m, mappa)

	if !eq {
		t.Error("Mappe non uguali")
	}
}
