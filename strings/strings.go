package strings

import "regexp"

// Дробит строку на элементы по переданному регулярному выражению
func RegSplit(text string, pattern string) []string {
	regex := regexp.MustCompile(pattern)
	result := regex.Split(text, -1)
	if len(result) == 0 || (len(result) == 1 && result[0] == "") {
		return make([]string, 0)
	}
	return result
}
