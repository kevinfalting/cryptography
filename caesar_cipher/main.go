package main

import (
	"flag"
	"fmt"
	"strings"
)

var flagOffset = flag.Int("o", 0, "ceasar cipher offset")
var flagEncrypt = flag.String("e", "", "encrypt text")
var flagDecrypt = flag.String("d", "", "decrypt text")

var englishAlphabet = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i",
	"j", "k", "l", "m", "n", "o", "p", "q", "r",
	"s", "t", "u", "v", "w", "x", "y", "z",
}

var englishAlphabetMap = map[string]int{
	"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9,
	"j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18,
	"s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
}

func main() {
	flag.Parse()

	if len(*flagEncrypt) == 0 && len(*flagDecrypt) == 0 ||
		len(*flagEncrypt) != 0 && len(*flagDecrypt) != 0 ||
		*flagOffset < 0 {
		flag.Usage()
		return
	}

	if len(*flagEncrypt) != 0 {
		fmt.Println(encrypt(*flagOffset, *flagEncrypt))
	} else {
		fmt.Println(decrypt(*flagOffset, *flagDecrypt))
	}
}

// implement a ceasar cipher encryptor
func encrypt(offset int, plaintext string) string {
	cipherText := ""
	plaintext = strings.ToLower(plaintext)
	plaintext = strings.ReplaceAll(plaintext, " ", "")

	offset = offset % len(englishAlphabet)

	for _, c := range plaintext {
		pos := englishAlphabetMap[string(c)]
		cipherIndex := (pos + offset) % len(englishAlphabet)
		cipherText += englishAlphabet[cipherIndex-1]
	}

	return cipherText
}

// TODO: implement a ceasar cipher decryptor

func decrypt(offset int, cipherText string) string {
	plainText := ""

	for _, c := range cipherText {
		pos := englishAlphabetMap[string(c)]
		plainIndex := mod((pos - offset), len(englishAlphabet))
		plainText += englishAlphabet[plainIndex-1]
	}

	return plainText
}

func mod(a, b int) int {
	return (a%b + b) % b
}
