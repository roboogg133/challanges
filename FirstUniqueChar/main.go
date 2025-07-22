package main

import (
	"fmt"
	"strings"
)

var input string = "aabbccddeefg"

func FirstUniqueChar(CompleteString string) *string {

	NewString := []rune(CompleteString)

	for i := range NewString {
		if len(NewString) == 0 {
			break
		}
		many := strings.Count(CompleteString, string(NewString[i]))
		if many > 1 {
			NewString = []rune(strings.ReplaceAll(CompleteString, string(NewString[i]), ""))
			NewString = []rune(CompleteString)
		} else {

			thischar := string(NewString[i])
			return &thischar
		}
	}

	return nil
}

func main() {

	//	i don't know if the way i did is the best way to do, but it works!

	char := FirstUniqueChar(input)

	if char == nil {
		fmt.Println("This string don't have an unique char")
		return
	}

	fmt.Println(*char)

}
