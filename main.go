package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)

type NRT0201 struct {
	NRT0201Sequence asn1.RawValue
}

func main() {
	var e error
	var a NRT0201
	data, _ := os.ReadFile("data/TDBRACLKENSA02818")
	_, e = asn1.Unmarshal(data, &a)
	fmt.Println(e)
	fmt.Println(a)

}

func printAsn(value asn1.RawValue) {
	fmt.Println("Tag: ", value.Tag)
	fmt.Println("Class: ", value.Class)
	fmt.Println("IsCompound: ", value.IsCompound)
	fmt.Println("Bytes: ", value.Bytes)
}
