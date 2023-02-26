package utils

import (
	"testing"
)

func TestVerificationService_DigitCode(t *testing.T) {
	for n := 6; n < 20; n++ {
		t.Log(GenerateDigitCode(n))
	}
}

func TestVerificationService_AlphaDigitCode(t *testing.T) {
	for n := 4; n < 20; n++ {
		t.Log(GenerateAlphaDigitCode(n))
	}
}
