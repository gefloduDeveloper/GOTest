package main

import "fmt"

func main(){

	var intSlice = []int{1,2,3}
	fmt.Println(sumSlice(intSlice))
	
	var float32Slice = []float32{1,2,3}
	fmt.Println(sumSlice(float32Slice))

	var float64Slice = []float64{1,2,3}
	fmt.Println(sumSlice(float64Slice))

	var slice = []int{}
	fmt.Println(isEmpty(slice))

	fmt.Println(isEmpty(float32Slice))

//	var intSlice = []int{1,2,3}
//	fmt.Println(sumIntSlice(intSlice))
//
//	var float32Slice = []float32{1,2,3}
//	fmt.Println(sumFloat32Slice(float32Slice))
//
//	var float64Slice = []float64{1,2,3}
//	fmt.Println(sumFloat64Slice(float64Slice))
}

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _, v:= range slice{
		sum+=v
	}
	return sum
}

func isEmpty[T any](slice []T) bool{
	return len(slice) == 0
}



func sumIntSlice(slice []int) int{
	var sum int
	for _, v:= range slice{
		sum+=v
	}
	return sum
}

func sumFloat32Slice(slice []float32) float32{
	var sum float32
	for _, v:= range slice{
		sum+=v
	}
	return sum
}

func sumFloat64Slice(slice []float64) float64{
	var sum float64
	for _, v:= range slice{
		sum+=v
	}
	return sum
}
