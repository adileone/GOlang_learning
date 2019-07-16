package main

import (
	"testing"
)

// TestOpenAccount : test
func TestOpenAccount(t *testing.T) {

	acc := OpenAccount(7)

	if acc.active == false {
		t.Error("Account inactive")
	}

	if acc.currentMoney != 7 {
		t.Error("wrong balance")
	}

}
