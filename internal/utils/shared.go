package utils

import (
	"fmt"
	"os"
)

func NewLine(n int) {
	for n < 0 {
		fmt.Println("")
		n--
	}
}

func Exit() {
	fmt.Println("Thank you. Do have a great day!")
	os.Exit(0)
}
