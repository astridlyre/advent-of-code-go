package main

import "testing"

var testInput = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
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

func TestReadInput(t *testing.T) {
	input := ReadInput("data.txt")
	got := len(input)
	want := 323

	if got != want {
		t.Errorf("Testing ReadInput: got length %v, want %v", got, want)
	}
}

func TestSlope(t *testing.T) {
	slope := Slope{3, 1}

	t.Run("Testing Slope x-axis", func(t *testing.T) {
		got := slope.x
		want := 3

		if got != want {
			t.Errorf("Testing Slope: got %v, want %v", got, want)
		}
	})

	t.Run("Testing Slope y-axis", func(t *testing.T) {
		got := slope.y
		want := 1

		if got != want {
			t.Errorf("Testing Slope: got %v, want %v", got, want)
		}
	})
}

func TestCheckTobogganRow(t *testing.T) {
	t.Run("Test that row find a tree at a local where there is a tree", func(t *testing.T) {
		row := testInput[0]
		got := CheckTobogganRow(row, 3)
		want := true

		if got != want {
			t.Errorf("Testing CheckTobogganRow: got %v, want %v", got, want)
		}
	})

	t.Run("That that a row does not find a tree where there is not tree", func(t *testing.T) {
		row := testInput[0]
		got := CheckTobogganRow(row, 1)
		want := false

		if got != want {
			t.Errorf("Testing CheckTobogganRow: got %v, want %v", got, want)
		}
	})
}

func TestDoTobagganRun(t *testing.T) {
	slope := Slope{3, 1}
	got := DoTobogganRun(testInput, &slope)
	want := 7

	if got != want {
		t.Errorf("Testing DoTobogganRun: got %v, want %v", got, want)
	}
}

func TestMultipleSlopes(t *testing.T) {
	slopes := []Slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	got := DoMultipleSlopes(testInput, &slopes)
	want := []int{2, 7, 3, 4, 2}
	CompareSlices(t, got, want)
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
