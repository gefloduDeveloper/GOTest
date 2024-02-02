package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value p points to is %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)
	p = &i
	*p = 1
	fmt.Printf("\nThe value p points to is %v", *p)
	fmt.Printf("\nThe value if i is: %v\n", i)

	//slices copy the pointers that's why it changes the value for both
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)

	//to optimize in case we have big data sets we can use pointers to not duplicate the memory usage
	var thing3 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing3 array is: %p", &thing3)
	var result2 [5]float64 = squarePoint(&thing3)
	fmt.Printf("\nThe result is: %v", result2)
	fmt.Printf("\nThe value of thing1 is: %v", thing3)
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

func squarePoint(thing4 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing4 array is: %p", &thing4)
	for i := range thing4 {
		thing4[i] = thing4[i] * thing4[i]
	}
	return *thing4
}
