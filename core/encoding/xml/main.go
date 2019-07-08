package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	Name    Name    `xml:"name"`
	Email   []Email `xml:"email"`
}

type Name struct {
	Family   string `xml:"family"`
	Personal string `xml:"personal"`
}

type Email struct {
	Type    string `xml:"type,attr"`
	Address string `xml:",chardata"`
}

func main() {
	person := Person{
		Name: Name{Family: "William", Personal: "A"},
		Email: []Email{
			Email{Type: "personal", Address: "lwlwilliam.a@gmail.com"},
			Email{Type: "work", Address: "lwlwilliam@foxmail.com"},
		},
	}

	saveXML("test.xml", person)

	p := new(Person)
	loadXML("test.xml", p)
	fmt.Println(p)
}

func saveXML(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := xml.NewEncoder(outFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func loadXML(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := xml.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}

func test() {
	str := `<?xml version="1.0" encoding="utf-8"?>
<person>
	<name>
		<family>Newmarch</family>
		<personal>Jan</personal>
	</name>
	<email type="personal">jan@newmarch.name</email>
	<email type="work">j.newmarch@boxhill.edu.au</email>
</person>`

	var person Person

	err := xml.Unmarshal([]byte(str), &person)
	checkError(err)

	fmt.Println("Family name: \"" + person.Name.Personal + "\"")
	fmt.Println("Second email address: \"" + person.Email[1].Address + "\"")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}
