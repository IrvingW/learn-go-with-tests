package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 4)
	want := "aaaa"
	if got != want {
		t.Errorf("got %q but expected %q", got, want)
	}
}

// This is a benchmark test(function name starts with Benchmark)
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}
