package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "(name)is(age)yearsold"
	knowledge := [][]string{{"name", "bob"}, {"age", "two"}}
	fmt.Println(evaluate(s, knowledge))
}

func evaluate(s string, knowledge [][]string) string {
	hash := make(map[string]string, len(knowledge))
	for _, kd := range knowledge {
		hash[kd[0]] = kd[1]
	}
	ans, start := &strings.Builder{}, -1
	for i, c := range s {
		if c == '(' {
			start = i
		} else if c == ')' {
			tmp := string(s[start+1 : i])
			if v, ok := hash[tmp]; ok {
				ans.WriteString(v)
			} else {
				ans.WriteByte('?')
			}
			start = -1
		} else if start < 0 {
			ans.WriteRune(c)
		}
	}
	return ans.String()
}
