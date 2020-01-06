package potato

import (
	"math"
	"strings"
)

func StringToName(s string) (val []byte, err error) {
	// ported from the potato codebase, libraries/chain/include/potato/chain/name.hpp
	var bit uint32 = 255 - 4
	sLen := uint32(len(s))
	val = make([]byte, 32)

	var c byte
	var i uint32 = 0
	for ; i < sLen; i++ {
		c = charToSymbol(s[i])
		j := 0
		for ; j < 6; j++ {
			if bit >= 0 {
				b := (c >> j) & 1
				b <<= (7 - (bit % 8))
				val[31-uint8(math.Floor(float64(bit)/float64(8.0)))] |= b
			}
			bit--
		}
	}

	return
}

func charToSymbol(c byte) byte {
	if c == '~' {
		return (c - '~') + 49
	}
	if c == '`' {
		return (c - '`') + 48
	}
	if c == ')' {
		return (c - ')') + 47
	}
	if c == '(' {
		return (c - '(') + 46
	}
	if c == '}' {
		return (c - '}') + 45
	}
	if c == '{' {
		return (c - '{') + 44
	}
	if c == ']' {
		return (c - ']') + 43
	}
	if c == '[' {
		return (c - '[') + 42
	}
	if c == '>' {
		return (c - '>') + 41
	}
	if c == '<' {
		return (c - '<') + 40
	}
	if c == ':' {
		return (c - ':') + 39
	}
	if c == '_' {
		return (c - '_') + 38
	}
	if c >= 'A' && c <= 'Z' {
		return (c - 'A') + 12
	}
	if c >= 'a' && c <= 'z' {
		return (c - 'a') + 12
	}
	if c >= '0' && c <= '9' {
		return (c - '0') + 2
	}
	if c == '-' {
		return (c - '-') + 1
	}
	return 0
}

var base32Alphabet = []byte(".-0123456789abcdefghijklmnopqrstuvwxyz_:<>[]{}()`~")

func NameToString(in []byte) string {
	// ported from libraries/chain/name.cpp in potato
	a := []byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'}

	var i uint32 = uint32(0)
	var bit uint32 = 255 - 4
	for bit >= 0 {
		var c byte = 0
		var j byte = 0
		for ; j < 6; j++ {
			if bit >= 0 {
				var b byte = in[31-uint8(math.Floor(float64(bit)/float64(8)))]
				b = (b >> (7 - (bit % 8))) & 1
				c |= (b << j)
			}
			bit--
		}
		a[i] = base32Alphabet[c&0x3f]
	}

	return strings.TrimRight(string(a), ".")
}
