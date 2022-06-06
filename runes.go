package main

import (
	"fmt"
	"unicode/utf8"
	// "unicode/utf8"
)

func runeRepresentation(){
	const runeRepresentated = "स्वप्निल"
	const normalEnglishText = "swopnil"
	fmt.Println(len(runeRepresentated))
	fmt.Println(len(normalEnglishText))

	for i := 0; i < len(runeRepresentated); i++{
		fmt.Printf("%x\n", runeRepresentated[i])
	}

	fmt.Println("How many runes in स्वप्निल?: ", utf8.RuneCountInString(runeRepresentated))

	fmt.Println("Runes can be used with range")

	for idx, runeValue := range runeRepresentated{
		fmt.Printf("%x starts at %d \n", runeValue, idx)
	}
}