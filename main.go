package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {

	sisi := strings.ToUpper(n)
	names := strings.Split(sisi, " ")
	var initials []string

	for _, value := range names {
		initials = append(initials, value[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	ln, mn := getInitials("soccer poker")
	fmt.Println(ln, mn)

	ln1, mn1 := getInitials("cloud strife")
	fmt.Println(ln1, mn1)

	ln2, mn2 := getInitials("baby")
	fmt.Println(ln2, mn2)
}
