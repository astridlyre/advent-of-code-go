package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const target = 2020

func main() {
	nums := ReadInput()
	// Function for Reducing Int Slice

	fn := func(x, y int) int { return x * y }
	// Part One => 211899
	result := FindTwoNumbersFor(target, nums)
	fmt.Printf("Numbers found: %v\n", result)
	fmt.Printf("Product of Numbers: %v\n", ReduceInts(result, fn, 1))

	// Part Two => 275765682
	result = FindThreeNumbersFor(target, nums)
	fmt.Printf("Numbers found: %v\n", result)
	fmt.Printf("Product of Numbers: %v\n", ReduceInts(result, fn, 1))
}

func ReadInput() []string {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func FindTwoNumbersFor(n int, s []string) []int {
	result := make([]int, 2)

	for i, n1 := range s {
		for _, n2 := range s[i:] {
			ints := ConvertToNums(n1, n2)
			if ints[0]+ints[1] == n {
				result[0] = ints[0]
				result[1] = ints[1]
				return result
			}
		}
	}
	return nil
}

func FindThreeNumbersFor(n int, s []string) []int {
	result := make([]int, 3)

	for i, n1 := range s {
		for j, n2 := range s[i:] {
			for _, n3 := range s[i+j:] {
				ints := ConvertToNums(n1, n2, n3)
				x := ints[0]
				y := ints[1]
				z := ints[2]
				if x+y+z == n {
					result[0] = x
					result[1] = y
					result[2] = z
					return result
				}
			}
		}
	}
	return nil
}

func ConvertToNums(s ...string) []int {
	result := make([]int, len(s))

	for i, c := range s {
		n, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		result[i] = n
	}
	return result
}

func ReduceInts(i []int, fn func(int, int) int, acc int) int {
	for _, n := range i {
		acc = fn(acc, n)
	}
	return acc
}
