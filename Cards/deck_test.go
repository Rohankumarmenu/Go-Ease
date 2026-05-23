package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()
	//len(d) != 16 for pass case because we have 4 suits and 4 values, so total cards will be 4*4=16
	//Other than 16 it will fail the test case,let say we put 15 or 17 it will fail the test case
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
