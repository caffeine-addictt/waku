package utils

import "strings"

// Invokes CleanString with terminal-specific operators
// like ;|&$.
func EscapeTermString(s string) string {
	return CleanString(s, '|', '&', ';', '`', '$')
}

// Escapes the given string from common escape sequences
// in 0(n) time, no regex.
//
//	\x1b[31mfoo\x1b[0m -> foo
//	\n\r\x00\x01\x02\x03\x04\x05\x06\x07\x08\x0B\x0C\x0E\x0F\x7f -> empty
func CleanString(s string, othersToEscape ...rune) string {
	var newS strings.Builder

	esc := make(map[byte]struct{}, len(othersToEscape))
	for _, v := range othersToEscape {
		esc[byte(v)] = struct{}{}
	}

	var i int
	for i < len(s) {
		switch s[i] {
		case '\x1b': // handle ANSI escape sequences
			i++

			if i < len(s) && s[i] == '[' {
				i++

				seek := true
				for i < len(s) && seek {
					switch s[i] {
					case 'm', 'A', 'B', 'C', 'D', 'H', 'f', 'J', 'K', 'c', 'n', 's', 'u':
						seek = false
					default:
						i++
					}
				}

				i++
			}
			continue

		case '\x00', '\x01', '\x02', '\x03', '\x04', '\x05', '\x06', '\x07', // Control characters
			'\x08', '\x0B', '\x0C', '\x0E', '\x0F', '\n', '\r',
			'\x7f': // Delete character
			i++
			continue
		}

		if _, ok := esc[s[i]]; !ok {
			newS.WriteByte(s[i])
		}

		i++
	}

	return newS.String()
}
