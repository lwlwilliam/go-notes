package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  Name    `json:"name"`
	Email []Email `json:"email"`
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

type Name struct {
	Family   string `json:"family"`
	Personal string `json:"personal"`
}

type Email struct {
	Kind    string `json:"kind"`
	Address string `json:"address"`
}

func main() {
	person := Person{
		Name:  Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "j.newmarch@gmail.com"}, Email{Kind: "work", Address: "jan@gmail.com"}},
	}

	saveJSON("person.json", person)

	// 这个算是一次性的
	//jn, _ := json.MarshalIndent(person, "", "\t")
	//fmt.Println(string(jn))

	man := new(Person) // man := new(interface{}) // 可以不指定数据类型
	loadJSON("person.json", man)
	fmt.Println(man.String())
}

func saveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	//encoder.SetIndent("", "\t")
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func loadJSON(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}
