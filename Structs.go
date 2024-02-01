package main

import "fmt"

func main() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	//prints a piece of array
	fmt.Println(intArr[1:3])

	//print pointers
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	//declare and print array with just one line
	intArr2 := [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	//Slices
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	//when capacity is exceeded the slice appends a new array with the same length as the original
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = []int32{10}
	intSlice = append(intSlice, intSlice3...)
	fmt.Println(intSlice)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	//Careful, maps always return something despite not existing
	fmt.Println(myMap2["Jason"])

	//it's possible to check if key exists
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid Name\n")
	}

	//loops
	//kinda a foreach
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	//like a while
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	i = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}

	//regular for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
