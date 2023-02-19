package test

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaa"

	if got != want {
		t.Errorf("expeat %s, but got %s", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
