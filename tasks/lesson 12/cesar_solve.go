/*
Program created by anpir
This program demonstrated how you can encode, decode and break Cesar cypher.
Real magic happens in shift() nad uncesar()
cesar() and uncesar() are resiliant to symbols other than a-z A-Z
uncesar() doesn't use special symbols to solve the cypher, only a-z and A-Z
both functions can take k which is bigger than 26, in which case k = k%26
key can contain space or special symbols, but they will not be used in unscrambler
*/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog!"
	k := 12
	key := "x j"

	fmt.Printf("Starting string:   %#v\n", s)

	enc := cesar(s, k)

	fmt.Printf("Encoded:           %#v\n", enc)
	fmt.Printf("Decoded (using k): %#v\n", cesar(enc, -k))

	dec_s, dec_k := uncesar(enc, key)
	fmt.Println("All possible solutions:")
	for i := range dec_s {
		fmt.Printf("%2d: \"%s\" (k = %2d)\n", i+1, dec_s[i], dec_k[i])
	}
}

// simple cesar scramble, to unscramble -k can be used
func cesar(s string, k int) string {
	res := ""
	for _, n := range s {
		res += string(shift(n, k))
	}
	return res
}

// uncesar returns cesar cypher decyphered and k used to encrypt s
// it returns slices with all possible solutions
// in error case, uncesar will return "" and 0
// keyword must be at least 2 symbols long, for proper unscramble
func uncesar(s string, keyword string) ([]string, []int) {
	r := []rune(s)         // we need to work with runes, not bytes!
	key := []rune(keyword) // same thing

	// cut all non letters in the beggining of key
	// because we need a letter to generate proper offsets
	for i := 0; i < len(key); i++ {
		if key[i] >= 'a' && key[i] <= 'z' {
			key = key[i:]
			break
		}
	}

	// contains offsets dependent from first letter of key
	offset := gen_offset(string(key))

	// used to track, whether we found a solution or not
	var match bool

	var k []int

	// for every index of r, try to match up key (using offsets)
	for i := 0; i <= len(r)-len(key); i++ {

		match = true

		for j := i; j < i+len(key); j++ {
			// if we are looking at a non-letter and key also contains a non-letter
			// skip this index
			if offset[j-i] == -66 && (r[j] > 'z' || r[j] < 'a') {
				continue
			}
			// if current rune doesn't match with first one with aplied offset, key wont fit, abort
			if shift(r[j], -offset[j-i]) != r[i] {
				match = false
				break
			}
		}

		// we have found a solution, save k so we could unscramble whole string later
		if match {
			k = append(k, int(key[0]-r[i]))
		}

	}

	// if we didn't find any solutions, return empty slices
	if len(k) == 0 {
		return []string{}, []int{}
	}

	// if we did find at least one solution
	// generate strings for every solution
	// and return them
	res := make([]string, len(k))
	for i, n := range k {
		res[i] = cesar(s, n)
		k[i] = normalize_offset(-k[i])
	}

	return res, k
}

// returns shifted rune r, k symbols to the right
// k can be any number, even negative
// can work both with small and large letters
// for any other symbols doesn't do anything, returns without changes
func shift(r rune, k int) rune {
	k = normalize_offset(k)

	// remember, whether letter was big and make it small
	wasUpper := r != unicode.ToLower(r)
	r = unicode.ToLower(r)

	// if we are not working with a letter, return it
	if r > 'z' || r < 'a' {
		return r
	}

	// get number of letter r, add k on top, normalize and convert back to proper rune
	r = (r-'a'+rune(k))%26 + 'a'

	// if letter was big, return big letter, else just return the small variant
	if wasUpper {
		return unicode.ToUpper(r)
	}
	return r
}

// return slice with numbers for each symbol of s, which is needed to be added
// to convert from s[0] to s[n]
// example: abacaba -> [0, 1, 0, 2, 0, 1, 0]
// for every non-letter offset will be equal to 66
func gen_offset(s string) []int {
	// we don't need to be working with big letters
	r := []rune(strings.ToLower(s))

	origin := r[0]

	result := make([]int, len(r))

	for i := 1; i < len(result); i++ {
		if r[i] < 'a' || r[i] > 'z' { // not a letter
			result[i] = -66
		} else {
			result[i] = int(r[i] - origin)
		}
	}

	return result
}

// converts any offset to range [0, 25]
// a bit of weird math
func normalize_offset(n int) int {
	if n < 0 {
		return 26 - ((-n) % 26)
	}
	return n % 26
}
