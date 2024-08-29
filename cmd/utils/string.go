package utils

// Escapes the given string for use in a shell command,
// should ensure that only ever 1 command is executed
func EscapeTermString(s string) string {
	var newS string
	for _, v := range s {
		switch v {
		case '|', '\n', '\r', '&', ';', '`', '$':
			newS += " "
		default:
			newS += string(v)
		}
	}
	return newS
}
