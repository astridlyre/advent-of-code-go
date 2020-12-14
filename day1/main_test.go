package main

import "testing"

func TestReadFile(t *testing.T) {
	got := len(ReadInput())
	want := 200
	if got != want {
		t.Errorf("Reading file, got %v want %v", got, want)
	}
}

func CompareSlices(t *testing.T, got, want []int) {
	t.Helper()
	for _, n := range got {
		found := false
		for _, w := range want {
			if n == w {
				found = true
			}
		}
		if !found {
			t.Errorf("Could not find the right number: %v", n)
		}
	}
}

func TestConvertToNums(t *testing.T) {
	test := []string{"20", "30", "40", "1", "69"}
	want := []int{20, 30, 40, 1, 69}
	got := ConvertStringsToInts(test)
	CompareSlices(t, got, want)
}

func TestFindTwoNumbersFor(t *testing.T) {
	numbers := []int{1721, 979, 366, 299, 675, 1456}

	t.Run("Two Numbers", func(t *testing.T) {
		got := FindTwoNumbersFor(2020, numbers)
		want := []int{1721, 299}
		CompareSlices(t, got, want)
	})

	t.Run("Two Nums, no matches shows empty slice", func(t *testing.T) {
		got := FindTwoNumbersFor(69, numbers)
		want := []int{}
		CompareSlices(t, got, want)
	})

	t.Run("Three Numbers", func(t *testing.T) {
		got := FindThreeNumbersFor(2020, numbers)
		want := []int{979, 366, 675}
		CompareSlices(t, got, want)
	})

	t.Run("Three Nums, no matches shows empy slice", func(t *testing.T) {
		got := FindThreeNumbersFor(23, numbers)
		want := []int{}
		CompareSlices(t, got, want)
	})
}

func TestReduceInts(t *testing.T) {
	ints := []int{5, 10, 2}
	fn := func(x, y int) int {
		return x * y
	}
	got := ReduceInts(ints, fn, 1)
	want := 100

	if got != want {
		t.Errorf("Reducing Ints, got %v want %v", got, want)
	}
}

func BenchmarkFindNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindThreeNumbersFor(2020, ConvertStringsToInts(ReadInput()))
	}
}
