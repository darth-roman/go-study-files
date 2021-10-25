package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct{
	Name string
	Surname string
	Tel []Telephone
}
type Telephone struct{
	Mobile bool
	Number string
}

func saveToJSON(filename *os.File, key interface{}){
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main(){
	recordOne := Record{
		Name: "Chouchane",
		Surname: "Roman",
		Tel: []Telephone{
			{Mobile: true, Number: "12345-67890"},
			{Mobile: false, Number: "09876-54321"},
			{Mobile: true, Number: "54321-09876"},
		},
	}

	saveToJSON(os.Stdout, recordOne)
}