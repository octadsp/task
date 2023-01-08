package main

import (
	"fmt"
)

func main() {
	huruf("a")
	huruf("b")
}

func huruf(huruf string) {
	if huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" {
		fmt.Printf("%s adalah vocal\n", huruf)
	} else {
		fmt.Printf("%s adalah konsonan\n", huruf)
	}
}
