package question557

import "strings"

func reverseString(s []byte)  string{
	n := len(s)
	for i:=0;i<(n>>1);i++{
		s[i],s[n-i-1] = s[n-i-1],s[i]
	}
	return string(s)
}
func reverseWords(s string) string {
	ss := strings.Split(s," ")
	for i:=0;i<len(ss);i++ {
		ss[i] = reverseString([]byte(ss[i]))
	}
	return strings.Join(ss," ")
}