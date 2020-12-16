package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Pattern to match each password string
var /* const */ rx = regexp.MustCompile("([0-9]+)|([a-z]:)|([a-z]+)")

// Policy represents a password policy
type Policy struct {
	min    int
	max    int
	letter string
}

// Password represents a password type
type Password struct {
	policy Policy
	value  string
}

// isValid returns true if the password contains the required amount of a letter
func (p *Password) isValid() bool {
	result := strings.Count(p.value, p.policy.letter)
	return result >= p.policy.min && result <= p.policy.max
}

// isValidTwo returns true of the password contains the required letter at the proper index
func (p *Password) isValidTwo() bool {
	// Make sure max is within valid bounds
	if p.policy.max > len(p.value) {
		log.Fatalf("invalid policy max: longer than password")
	}
	// Get chars at each policy position
	l, p1, p2 := p.policy.letter, string(p.value[p.policy.min-1]), string(p.value[p.policy.max-1])
	if p1 == l && p2 != l || p1 != l && p2 == l {
		return true
	}
	return false
}

func main() {
	input := ReadInput("day2/data.txt")
	partOne, partTwo := ValidatePasswords(input)
	// Part one => 454
	fmt.Printf("Valid passwords according to Part one rules: %v\n", partOne)

	// Part two => 649
	fmt.Printf("Valid Passwords according to Part two rules: %v\n", partTwo)
}

// ReadInput gets puzzle input from a file
func ReadInput(s string) []string {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

// ValidatePasswords returns the number of valid passwords according to policy
func ValidatePasswords(s []string) (int, int) {
	isValid, isValidTwo := 0, 0
	for _, p := range s {
		pwd := ParsePassword(p)
		if pwd.isValid() {
			isValid++
		}
		if pwd.isValidTwo() {
			isValidTwo++
		}
	}
	return isValid, isValidTwo
}

// ParsePassword takes a string and parses it into a Password struct
func ParsePassword(s string) Password {
	// Match password
	match := rx.FindAllString(s, -1)
	if len(match) != 4 || match == nil {
		log.Fatalf("unable to match password")
	}

	// Attempt to convert numerical min & max values
	min, err1 := strconv.Atoi(match[0])
	max, err2 := strconv.Atoi(match[1])
	if err1 != nil || err2 != nil {
		log.Fatalf("unable to convert min & max bounds")
	}

	// Password literal to return
	return Password{
		Policy{
			min,
			max,
			string(match[2][0]),
		},
		match[3],
	}
}
