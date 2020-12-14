package main

import "testing"

func TestReadInput(t *testing.T) {
	input := ReadInput("data.txt")
	got := len(input)
	want := 1000

	if got != want {
		t.Errorf("Reading Input: got length %v, want %v", got, want)
	}
}

func TestParseLine(t *testing.T) {
	input := "4-6 b: bbbdbtbbbj"
	password := ParsePassword(input)
	wantPolicyMin := 4
	wantPolicyMax := 6
	wantLetter := "b"
	wantPassword := "bbbdbtbbbj"

	if password.policy.min != wantPolicyMin || password.policy.max != wantPolicyMax {
		t.Errorf("Password policy: got %v, want %v", password.policy, Policy{4, 6, "b"})
	}

	if password.policy.letter != wantLetter {
		t.Errorf("Password policy letter: got %v, want%v", password.policy.letter, wantLetter)
	}

	if password.value != wantPassword {
		t.Errorf("Password value: got %v, want %v", password.value, wantPassword)
	}
}

func TestPasswordisValid(t *testing.T) {
	pwd := Password{
		Policy{
			4,
			6,
			"b",
		},
		"bbbdttbbbj",
	}

	t.Run("Part One Password is Valid", func(t *testing.T) {
		got := pwd.isValid()
		want := true

		if got != want {
			t.Errorf("Password isValid: got %v, want %v", got, want)
		}
	})

	t.Run("Part Two Password is Invalid", func(t *testing.T) {
		got := pwd.isValidTwo()
		want := false

		if got != want {
			t.Errorf("Password isValidTwo: got %v, want %v", got, want)
		}
	})
}

func TestValidatePasswords(t *testing.T) {
	pwds := []string{"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc"}
	got1, got2 := ValidatePasswords(pwds)
	want1, want2 := 2, 1

	if got1 != want1 || got2 != want2 {
		t.Errorf("Validate Passwords: got1 %v, want1 %v | got2 %v, want2 %v", got1, want1, got2, want2)
	}
}
