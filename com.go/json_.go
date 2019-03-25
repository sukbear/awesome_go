package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Type string
	City string
	Country string
}
type VCard struct{
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private","CQ","China"}
	wa := &Address{"work","BJ","China"}
	vc := &VCard{"Jan","gui",[]*Address{pa,wa},"none"}

	js, _ := json.Marshal(vc)
	fmt.Printf("JSON fomat: %s",js)

	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	enc := json.NewEncoder(file)
	err:=enc.Encode(enc)
	if err!= nil{
		fmt.Print("Error in encoding json")
	}
}