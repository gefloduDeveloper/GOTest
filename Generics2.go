package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type contactInfo struct{
	Name string
	Email string
}

type purchaseInfo struct{
	Name string
	Price float32
	Amount int
}

func main(){
	var contacts []contactInfo = loadJson[contactInfo]("./contactInfo.json")
	fmt.Printf("\n%+v",contacts)

	var purchases []purchaseInfo = loadJson[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n%+v",purchases)

}

func loadJson[T contactInfo | purchaseInfo](filepath string) []T{
	var data, _ = ioutil.ReadFile(filepath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}