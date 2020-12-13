package main

import "testing"

func TestReadFile(t *testing.T) {
	got := len(ReadInput())
	want := 200
	if got != want {
		t.Errorf("Reading file, got %v want %v", got, want)
	}
}

func TestConvertToNums(t *testing.T) {
	ints := ConvertToNums("12", "35")
	n1 := ints[0]
	n2 := ints[1]
	g1, g2 := 12, 35

	if g1 != n1 || g2 != n2 {
		t.Errorf("Converting numbers, got n1 %v n1 %v, want n1 %v n2 %v", g1, g2, n1, n2)
	}
}

func TestFindTwoNumbersFor(t *testing.T) {

	compareSlices := func(t *testing.T, got, want []int) {
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
	numbers := []string{"1721", "979", "366", "299", "675", "1456"}

	t.Run("Two Numbers", func(t *testing.T) {
		got := FindTwoNumbersFor(2020, numbers)
		want := []int{1721, 299}
		compareSlices(t, got, want)
	})

	t.Run("Three Numbers", func(t *testing.T) {
		got := FindThreeNumbersFor(2020, numbers)
		want := []int{979, 366, 675}
		compareSlices(t, got, want)
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
