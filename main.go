package main

import (
	"fmt"
	"regexp"
)

func ValiddateEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9._!#]+@[a-zA-Z0-9][a-z0-9.]*?[.]+[a-z]*[^.]$")

	return re.MatchString(email)
}

func main() {
	fmt.Println(ValiddateEmail("nobby.phala@gmail.com"))
}
