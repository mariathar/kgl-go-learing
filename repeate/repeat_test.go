package repeate

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatCount(t *testing.T) {
	repeated := RepeatCount("a", 5)
	expectedCount := 5

	if strings.Count(repeated, "a") != expectedCount {
		t.Errorf("expectedCount %q but got %q", expectedCount, strings.Count(repeated, "a"))
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
