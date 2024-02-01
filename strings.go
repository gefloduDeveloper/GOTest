package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[1]
	fmt.Printf("%v, %T", indexed, indexed)
	//uses the extra byte given the UTF-8 encoding
	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	//Runes (basically int-32 based chars)
	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"g", "e", "r", "m", "a", "n"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)

	//Cannot change value assigned
	//catStr[0] = 'a'

	//last method is innefficient since it creates a new string each time we add the new rune
	//we can improve using the string builder
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("\n%v", catStr2)
}
