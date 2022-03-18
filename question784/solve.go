package question784

import "strings"

var res []string

func letterCasePermutation(s string) []string {
	res = []string{}
	path := ""
	backtrace(path, s, 0)
	return res
}

func backtrace(path string, s string, start int) {
	if len(path) == len(s) {
		res = append(res, path)
		return
	}
	ch := s[start]
	if ch >= 'a' && ch <= 'z' ||ch>='A'&&ch<='Z'{
		path += strings.ToUpper(string(ch))
		backtrace(path, s, start+1)
		path = path[:len(path)-1]
		path += strings.ToLower(string(ch))
		backtrace(path, s, start+1)
		path = path[:len(path)-1]
	} else {
		path += string(ch)
		backtrace(path, s, start+1)
	}
}
