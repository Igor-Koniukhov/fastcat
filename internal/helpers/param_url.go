package helpers

func Param(pattern, path string) string {
	var i, j int
	for i < len(path) {
		switch {
		case j >= len(pattern):
			if pattern[len(pattern)-1] == '/' {
				return path[i:]
			}
			return ""
		case pattern[j] == ':':
			var nextc byte
			_, nextc, j = match(pattern, isAlnum, j+1)
			_, _, i = match(path, matchPart(nextc), i)
		case path[i] == pattern[j]:
			i++
			j++
		default:
			return ""
		}
	}
	return ""
}

func isAlpha(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isAlnum(ch byte) bool {
	return isAlpha(ch) || isDigit(ch)
}

func match(s string, f func(byte) bool, i int) (matched string, next byte, j int) {
	j = i
	for j < len(s) && f(s[j]) {
		j++
	}
	if j < len(s) {
		next = s[j]
	}
	return s[i:j], next, j
}

func matchPart(b byte) func(byte) bool {
	return func(c byte) bool {
		return c != b && c != '/'
	}
}

