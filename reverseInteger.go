package main

import (
	"fmt"
	"strconv"
	"bytes"
)

func reverse(x int) int {
	ind := ""
	if x < 0 {
		x = -1 * x
		ind = "-"
	}

	s := strconv.Itoa(x)

	var b bytes.Buffer

	for i := len(s) - 1; i >= 0; i-- {
		b.WriteString(string(s[i]))
	}

	t, _ := strconv.Atoi(ind + b.String())

	if t > 2147483647 {
		return 0
	} else if t < -2147483648 {
		return 0
	}
	return t
}

func main() {
	r := reverse(1563847412)
	fmt.Println(r)
}
