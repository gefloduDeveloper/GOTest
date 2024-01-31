package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum++
	fmt.Println(intNum)

	var floatNum float32 = 12345678.9
	fmt.Println(floatNum)

	/*Not valid - As expected
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + intNum32*/

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello \n world"
	fmt.Println(myString)
	myString = `Hello
world`
	fmt.Println(myString)
	myString = "Hello" + " " + "world"
	fmt.Println(myString)

	//Len is string number of bytes
	fmt.Println(len("y"))

	//lenght in character terms use the utf-8 function
	fmt.Println(utf8.RuneCountInString(myString))

	//Runes-a new concept
	var myRune rune = 'a'
	fmt.Println(myRune)

	//as usual
	var myBoolean bool = false
	fmt.Println(myBoolean)

	//vars with auto assigned type
	myVar := "text"
	fmt.Println(myVar)

	//multiple definitions in one line
	var1, var2 := 1, "Something"
	fmt.Println(var1, var2)

	const myConst string = "const value"
	//Cannot change
	//myConst = "Another value"
	fmt.Println(myConst)
}
