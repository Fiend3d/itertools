package itertools

import (
	"fmt"
	"slices"
	"testing"
)

func TestCombinations(t *testing.T) {
	items := []int{1, 2, 3}
	got := slices.Collect(Combinations(items, 2))
	want := [][]int{{1, 2}, {1, 3}, {2, 3}}
	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %v want %v", got, want)
	}
	for i := range got {
		if !slices.Equal(got[i], want[i]) {
			t.Fatalf("mismatch at index %d: got %v want %v", i, got[i], want[i])
		}
	}
}

func TestPermutations(t *testing.T) {
	items := []int{1, 2, 3}
	got := slices.Collect(Permutations(items, 2))
	want := [][]int{
		{1, 2}, {1, 3},
		{2, 1}, {2, 3},
		{3, 1}, {3, 2},
	}

	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %v want %v", got, want)
	}
	for i := range got {
		if !slices.Equal(got[i], want[i]) {
			t.Fatalf("mismatch at index %d: got %v want %v", i, got[i], want[i])
		}
	}
}

func TestProduct(t *testing.T) {
	items := []int{1, 2}
	got := slices.Collect(Product(items, 2))
	want := [][]int{
		{1, 1}, {1, 2},
		{2, 1}, {2, 2},
	}
	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %v want %v", got, want)
	}
	for i := range got {
		if !slices.Equal(got[i], want[i]) {
			t.Fatalf("mismatch at index %d: got %v want %v", i, got[i], want[i])
		}
	}
}

func TestCombinationsWithReplacement(t *testing.T) {
	items := []int{1, 2}
	got := slices.Collect(CombinationsWithReplacement(items, 2))
	want := [][]int{
		{1, 1}, {1, 2},
		{2, 2},
	}
	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %v want %v", got, want)
	}
	for i := range got {
		if !slices.Equal(got[i], want[i]) {
			t.Fatalf("mismatch at index %d: got %v want %v", i, got[i], want[i])
		}
	}
}

func TestCombinationsWithReplacement2(t *testing.T) {
	items := []string{"A", "B", "C"}
	got := slices.Collect(combinationsWithReplacement(items, 2))

	want := [][]string{
		{"A", "A"}, {"A", "B"}, {"A", "C"},
		{"B", "B"}, {"B", "C"},
		{"C", "C"},
	}
	t.Log("Say bye")

	fmt.Printf("Got: %v\n", got)
	fmt.Printf("Want: %v\n", want)
}
