package tests

import (
	"go_tuts/utils"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := utils.Add(2, 3)
	expected := 5

	if sum != expected {
		t.Errorf("Expected :%d  and got : %d ", expected, sum)
	}

}