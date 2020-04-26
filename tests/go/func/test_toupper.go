package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	table := z01.MultRandASCII()

	table = append(table,
		"Hello! How are you?",
	)

	for _, arg := range table {
		z01.Challenge("ToUpper", student.ToUpper, correct.ToUpper, arg)
	}
}