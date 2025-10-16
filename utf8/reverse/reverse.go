//go:build !solution

package reverse

import (
	"strings"
)

func Reverse(input string) string {
	var sb strings.Builder
	m := make(map[int]string)
	i := 1
	for _, r := range input {
		m[i] = string(r)
		i++
	}
	var str_list []string
	for j := i; j > 0; j-- {
		str_list = append(str_list, m[j])
	}
	for _, str := range str_list {
		sb.WriteString(str)
	}
	return sb.String()
}
