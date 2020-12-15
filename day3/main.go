package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// aTree is the character symbolizing a tree
const aTree = "#"

// Slope represents an x-y path down a hill
type Slope struct {
	x int
	y int
}

func main() {
	input := ReadInput("day3/data.txt")

	// Part One
	slope := Slope{3, 1}
	trees := DoTobogganRun(input, &slope)
	// Part One Result => 178
	fmt.Printf("Number of trees found on Slope {%v %v}: %v\n", slope.x, slope.y, trees)

	// Part Two
	slopes := []Slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	treesFound := DoMultipleSlopes(input, &slopes)
	getProduct := func(x, y int) int { return x * y }
	// Part Two => 3492520200
	fmt.Printf("Product of the number of trees encountered: %v\n", ReduceInts(treesFound, getProduct, 1))
}

// IsATree takes a string and returns true if it is aTree
func IsATree(s string) bool {
	return s == aTree
}

// ReadInput parses puzzle input into a slice of strings
func ReadInput(s string) []string {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

// CheckTobogganRow tests a row (s string) for the presence of a tree at col x
func CheckTobogganRow(s string, x int) bool {
	if x >= len(s) {
		log.Fatalf("tried to index past string length")
	}
	return IsATree(string(s[x]))
}

// DoTobogganRun performs a simulated run down the mountain, returning the number of tree found
func DoTobogganRun(s []string, slope *Slope) int {
	// pos: current col position
	pos, rowLength, treesFound := 0, len(s[0]), 0
	for i := 0; i < len(s); i += slope.y {
		found := CheckTobogganRow(s[i], pos)
		if found {
			treesFound++
		}
		// update position for next row
		pos += slope.x
		if pos >= rowLength {
			pos -= rowLength
		}
	}
	return treesFound
}

// DoMultipleSlopes tests the probability of a sudden arboreal stop against multiple slopes
func DoMultipleSlopes(s []string, slopes *[]Slope) []int {
	treesFound := make([]int, len(*slopes))
	for i, slope := range *slopes {
		treesFound[i] = DoTobogganRun(s, &slope)
	}
	return treesFound
}

// ReduceInts performs an array reduce method on the slice of ints
func ReduceInts(i []int, fn func(int, int) int, acc int) int {
	for _, n := range i {
		acc = fn(acc, n)
	}
	return acc
}
