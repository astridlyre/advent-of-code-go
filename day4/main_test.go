package main

import (
	"strings"
	"testing"
)

// Contains two valid passports according to Part One
var testInputOne = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

// Contains all invalid passports according to Part Two
var allInvalidPassports = `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`

// Contains all valid passports according to Part Two
var allValidPassports = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`

func TestGetValue(t *testing.T) {
	test := "ecl:gry"
	got := GetValue(test)
	want := "gry"

	if got != want {
		t.Errorf("Testing GetValue: got %v, want %v", got, want)
	}
}

func TestParseEyeColor(t *testing.T) {
	t.Run("ParseEyeColor returns true on valid eye color", func(t *testing.T) {
		test := "ecl:gry"
		got, gotOk := ParseEyeColor(test)
		want, wantOk := "gry", true

		if got != want || gotOk != wantOk {
			t.Errorf("Testing ParseEyeColor: got %v %v, want %v %v", got, gotOk, want, wantOk)
		}
	})

	t.Run("ParseEyeColor returns false on invalid eye color", func(t *testing.T) {
		test := "ecl:wrong"
		got, gotOk := ParseEyeColor(test)
		want, wantOk := "wrong", false

		if got != want || gotOk != wantOk {
			t.Errorf("Testing ParseEyeColor: got %v %v, want %v %v", got, gotOk, want, wantOk)
		}
	})
}

func TestParseHairColor(t *testing.T) {
	t.Run("ParseHairColor returns true on valid hair color", func(t *testing.T) {
		test := "hcl:#fffffd"
		got, gotOk := ParseHairColor(test)
		want, wantOk := "#fffffd", true

		if got != want || gotOk != wantOk {
			t.Errorf("Testing ParseHairColor: got %v %v, want %v %v", got, gotOk, want, wantOk)
		}
	})

	t.Run("ParseHairColor returns false on invalid hair color", func(t *testing.T) {
		test := "hcl:#abxxd822"
		_, gotOk := ParseHairColor(test)
		wantOk := false

		if gotOk != wantOk {
			t.Errorf("Testing ParseHairColor: got %v, want %v", gotOk, wantOk)
		}
	})
}

func TestParseHeight(t *testing.T) {
	t.Run("ParseHeight returns true on valid height in cm", func(t *testing.T) {
		test := "hgt:183cm"
		got, gotOk := ParseHeight(test)
		want, wantOk := Height{183, "cm"}, true

		if got != want || gotOk != wantOk {
			t.Errorf("Testing ParseHeight: got %v %v, want %v %v", got, gotOk, want, wantOk)
		}
	})

	t.Run("ParseHeight returns false on invalid height in cm", func(t *testing.T) {
		test := "hgt:540cm"
		got, gotOk := ParseHeight(test)
		want, wantOk := Height{540, "cm"}, false

		if got != want || gotOk != wantOk {
			t.Errorf("Testing ParseHeight: got %v %v, want %v %v", got, gotOk, want, wantOk)
		}
	})

	t.Run("ParseHeight returns true on valid height in inches", func(t *testing.T) {
		test := "hgt:75in"
		got, gotOk := ParseHeight(test)
		want, wantOk := Height{75, "in"}, true

		if got != want || gotOk != wantOk {
			t.Errorf("Testing ParseHeight: got %v %v, want %v %v", got, gotOk, want, wantOk)
		}
	})

	t.Run("ParseHeight returns false on invalid height in inches", func(t *testing.T) {
		test := "hgt:2in"
		_, gotOk := ParseHeight(test)
		wantOk := false

		if gotOk != wantOk {
			t.Errorf("Testing ParseHeight: got %v, want %v", gotOk, wantOk)
		}
	})
}

func TestParseYear(t *testing.T) {
	t.Run("ParseYear returns true with a valid year within range", func(t *testing.T) {
		test := "byr:2000"
		got, gotOk := ParseYear(test, 1920, 2002)
		want, wantOk := 2000, true

		if got != want || gotOk != wantOk {
			t.Errorf("Testing ParseYear: got %v %v, want %v %v", got, gotOk, want, wantOk)
		}
	})

	t.Run("ParseYear returns false with an invalid year", func(t *testing.T) {
		test := "iyr:1000"
		got, gotOk := ParseYear(test, 2010, 2020)
		want, wantOk := 1000, false

		if got != want || gotOk != wantOk {
			t.Errorf("Testing ParseYear: got %v %v, want %v %v", got, gotOk, want, wantOk)
		}
	})
}

func TestParseId(t *testing.T) {
	t.Run("ParseId returns true with a valid passport id", func(t *testing.T) {
		test := "pid:860033327"
		got, gotOk := ParseId(test)
		want, wantOk := 860033327, true

		if got != want || gotOk != wantOk {
			t.Errorf("Testing ParseId: got %v %v, want %v %v", got, gotOk, want, wantOk)
		}
	})

	t.Run("ParseId returns false with an invalid passport id", func(t *testing.T) {
		test := "pid:35643775"
		_, got := ParseId(test)
		want := false

		if got != want {
			t.Errorf("Testing ParseId: got %v, want %v", got, want)
		}
	})
}

func TestParsePassport(t *testing.T) {
	t.Run("ParsePassport returns a valid passport with valid input", func(t *testing.T) {
		test := "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\n byr:1937 iyr:2017 cid:147 hgt:183cm"
		got := ParsePassport(test)
		want := Passport{"gry", "#fffffd", Height{183, "cm"}, 2017, 2020, 1937, 860033327, true, true}

		if got != want {
			t.Errorf("Testing ParsePassport: got %v, want %v", got, want)
		}
	})

	t.Run("ParsePassport returns an invalid passport with invalid input", func(t *testing.T) {
		test := "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\n hcl:#cfa07d byr:1929"
		got := ParsePassport(test).isValid
		want := false

		if got != want {
			t.Errorf("Testing ParsePassport: got %v, want %v", got, want)
		}
	})

	t.Run("ParsePassport returns false when not enough fields", func(t *testing.T) {
		test := "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\n hcl:#cfa07d byr:1929"
		got := ParsePassport(test).hasAllFields
		want := false

		if got != want {
			t.Errorf("Testing ParsePassport: got %v, want %v", got, want)
		}
	})
}

func TestReadFile(t *testing.T) {
	got := len(ReadInput("data.txt"))
	want := 291
	if got != want {
		t.Errorf("Reading file, got %v want %v", got, want)
	}
}

func TestValidatePassports(t *testing.T) {
	t.Run("Passports according to Part One", func(t *testing.T) {
		got, _ := ValidatePassports(strings.Split(testInputOne, "\n\n"))
		want := 2

		if got != want {
			t.Errorf("Testing ValidatePassports: got %v, want %v", got, want)
		}
	})

	t.Run("Passports according to Part Two: all invalid", func(t *testing.T) {
		_, got := ValidatePassports(strings.Split(allInvalidPassports, "\n\n"))
		want := 0

		if got != want {
			t.Errorf("Testing ValidatePassports: got %v, want %v", got, want)
		}
	})

	t.Run("Passports according to Part Two: all valid", func(t *testing.T) {
		_, got := ValidatePassports(strings.Split(allValidPassports, "\n\n"))
		want := 4

		if got != want {
			t.Errorf("Testing ValidatePassports: got %v, want %v", got, want)
		}
	})
}
