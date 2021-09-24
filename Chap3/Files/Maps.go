package main

import "fmt"

func main(){
	iMap := make(map[string]int)
	fmt.Println(iMap)

	anotherMap := map[string]int{
		"k1": 12,
		"k2": 13,
		"k3": 14,
		"k4": 15,
	}

	// The iteration is not in order
	// for key, value := range anotherMap{
	// 	fmt.Println(key, value)
	// }

	delete(anotherMap, "k1")
	for key, value := range anotherMap{
		fmt.Println(key, value)
	}

	_, ok := iMap["doesItExist"]
	if ok {
		fmt.Println("Exists")
	} else {
		fmt.Println("Does NOT Exist")

	}
}