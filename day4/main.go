// Package main provides a function to Valid Passports
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// regexp variables to match passport fields
var ecl = regexp.MustCompile("ecl:\\w+")
var hcl = regexp.MustCompile("hcl:#\\w+")
var hgt = regexp.MustCompile("hgt:\\w+")
var iyr = regexp.MustCompile("iyr:\\d+")
var eyr = regexp.MustCompile("eyr:\\d+")
var byr = regexp.MustCompile("byr:\\d+")
var pid = regexp.MustCompile("pid:\\d+")
var nums = regexp.MustCompile("\\d+")

// Variables to perform passport field validation
var isValidHairColor = regexp.MustCompile("^#[a-fA-F0-9]{6}$")
var validEyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
var isValidPassportId = regexp.MustCompile("^\\d{9}$")
var isHeightIn = regexp.MustCompile("^\\d{2}in$")
var isHeightCm = regexp.MustCompile("^\\d{3}cm$")
var hasAllFields = regexp.MustCompile("(ecl)|(hcl)|(hgt)|(iyr)|(eyr)|(byr)|(pid)")

// Constants to perform passport field validation
const minBirthYear = 1920
const maxBirthYear = 2002
const minIssueYear = 2010
const maxIssueYear = 2020
const minExpYear = 2020
const maxExpYear = 2030
const minHeightCm = 150
const maxHeightCm = 193
const minHeightIn = 59
const maxHeightIn = 76

// Passport represents a passport object
type Passport struct {
	eyeColor     string
	hairColor    string
	height       Height
	issueYear    int
	expYear      int
	birthYear    int
	id           int
	isValid      bool
	hasAllFields bool
}

// Height represents a height value, along with the measurement unit
type Height struct {
	value int
	unit  string
}

// ParsePassport returns a Passport struct with fields parsed
func ParsePassport(s string) Passport {
	eyeColor, eclOk := ParseEyeColor(ecl.FindString(s))
	hairColor, hclOk := ParseHairColor(hcl.FindString(s))
	height, hgtOk := ParseHeight(hgt.FindString(s))
	issueYear, iyrOk := ParseYear(iyr.FindString(s), minIssueYear, maxIssueYear)
	expYear, eyrOk := ParseYear(eyr.FindString(s), minExpYear, maxExpYear)
	birthYear, byrOk := ParseYear(byr.FindString(s), minBirthYear, maxBirthYear)
	id, pidOk := ParseId(pid.FindString(s))
	hasFields := len(hasAllFields.FindAllString(s, -1)) == 7

	return Passport{
		eyeColor,
		hairColor,
		height,
		issueYear,
		expYear,
		birthYear,
		id,
		eclOk && hclOk && hgtOk && iyrOk && eyrOk && byrOk && pidOk && hasFields,
		hasFields,
	}
}

// GetValue returns the value from a key:value pair
func GetValue(s string) string {
	if len(s) > 0 {
		return strings.Split(s, ":")[1]
	}
	return ""
}

// ParseEyeColor returns a parsed eye color, and its validity
func ParseEyeColor(s string) (string, bool) {
	eyeColor := GetValue(s)
	for _, c := range validEyeColors {
		if c == eyeColor {
			return eyeColor, true
		}
	}
	return eyeColor, false
}

// ParseHairColor returns a parsed hair color, and its validity
func ParseHairColor(s string) (string, bool) {
	hairColor := GetValue(s)
	if isValidHairColor.MatchString(hairColor) {
		return hairColor, true
	}
	return hairColor, false
}

func ParseHeight(s string) (Height, bool) {
	height := GetValue(s)
	n := nums.FindString(height)
	h, err := strconv.Atoi(n)
	if err != nil {
		return Height{}, false
	}

	if isHeightCm.MatchString(height) {
		return Height{h, "cm"}, h >= minHeightCm && h <= maxHeightCm
	}

	if isHeightIn.MatchString(height) {
		return Height{h, "in"}, h >= minHeightIn && h <= maxHeightIn
	}
	return Height{}, false
}

// ParseYear returns a parsed year, validated to be between low and high bounds
func ParseYear(s string, low, high int) (int, bool) {
	n := GetValue(s)
	year, err := strconv.Atoi(n)
	if err != nil {
		return 0, false
	}
	return year, year >= low && year <= high
}

// ParseId returns a parsed Passport ID
func ParseId(s string) (int, bool) {
	n := GetValue(s)
	id, err := strconv.Atoi(n)
	if err != nil {
		return 0, false
	}
	return id, isValidPassportId.MatchString(n)
}

// ReadInput processes a textfile of puzzle input and returns a slice of strings
func ReadInput(f string) []string {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n\n")
}

// ValidatePassports returns the number of passports with all fields present, and all fields valid
func ValidatePassports(s []string) (int, int) {
	invalidPassportsOne, invalidPassportsTwo := 0, 0
	for _, p := range s {
		passport := ParsePassport(p)
		if passport.hasAllFields {
			invalidPassportsOne++
		}
		if passport.isValid {
			invalidPassportsTwo++
		}
	}
	return invalidPassportsOne, invalidPassportsTwo
}

func main() {
	input := ReadInput("day4/data.txt")
	partOne, partTwo := ValidatePassports(input)
	// Part One => 260
	fmt.Printf("Valid passports (Part one): %v\n", partOne)
	// Part Two => 153
	fmt.Printf("Valid passports (Part two): %v\n", partTwo)
}
