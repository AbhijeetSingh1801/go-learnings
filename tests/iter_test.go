package tests

import (
	"testing"
	"go_tuts/utils"
)
func TestRepeat(t *testing.T) {
	repeated := utils.Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		utils.Repeat("a")
	}
}