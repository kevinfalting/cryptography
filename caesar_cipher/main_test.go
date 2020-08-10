package main

import (
	"strings"
	"testing"
)

func TestCeasarCipher(t *testing.T) {
	testCases := []struct {
		plainText  string
		offset     int
		cipherText string
	}{
		{
			plainText:  "a",
			offset:     0,
			cipherText: "a",
		},
		{
			plainText:  "i love you vikki",
			offset:     6,
			cipherText: "orubkeuaboqqo",
		},
	}

	for _, test := range testCases {
		got := encrypt(test.offset, test.plainText)
		if got != test.cipherText {
			t.Fatalf("expected %q got %q", test.cipherText, got)
		}

		got = decrypt(test.offset, test.cipherText)
		p := strings.ReplaceAll(test.plainText, " ", "")
		if got != p {
			t.Fatalf("expected %q got %q", p, got)
		}
	}
}
