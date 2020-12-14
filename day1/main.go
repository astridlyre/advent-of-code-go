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
	nums := ConvertStringsToInts(ReadInput())
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

// ReadInput processes a textfile of puzzle input and returns a slice of strings
func ReadInput() []string {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

// ConvertStringsToInts converts slice of strings to ints
func ConvertStringsToInts(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ints[i] = n
	}
	return ints
}

// FindTwoNumbersFor searches an array of numbers for two which add to n
func FindTwoNumbersFor(n int, nums []int) []int {
	result := make([]int, 2)

	for i, n1 := range nums {
		for _, n2 := range nums[i:] {
			if n1+n2 == n {
				result[0] = n1
				result[1] = n2
				return result
			}
		}
	}
	return nil
}

// FindThreeNumbersFor search an array of numbers for three which add to n
func FindThreeNumbersFor(n int, nums []int) []int {
	result := make([]int, 3)

	for i, n1 := range nums {
		for j, n2 := range nums[i:] {
			for _, n3 := range nums[i+j:] {
				if n1+n2+n3 == n {
					result[0] = n1
					result[1] = n2
					result[2] = n3
					return result
				}
			}
		}
	}
	return nil
}

// ReduceInts performs an array reduce method on the slice of ints
func ReduceInts(i []int, fn func(int, int) int, acc int) int {
	for _, n := range i {
		acc = fn(acc, n)
	}
	return acc
}
