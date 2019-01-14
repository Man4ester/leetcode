package main

import (
	"fmt"
	"strings"
)

func main() {
	b := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(b))
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)
	for _, val := range emails {
		parts := strings.Split(val, "@")
		localName := parts[0]

		localName = strings.Replace(localName, ".", "", -1)

		plusParts := strings.Split(localName, "+")

		finalPart := plusParts[0] + parts[1]
		m[finalPart] = true
	}
	return len(m)
}
