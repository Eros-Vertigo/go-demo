package test

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{2, 5})
	want := []int{6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expect:%d, but got %d", want, got)
	}
}
