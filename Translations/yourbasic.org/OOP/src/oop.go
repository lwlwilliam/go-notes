package main

import (
	"fmt"
)

type Sleeper interface {
	Sleep()
}

type Animal struct {
	// ...
}

func (a *Animal) Eat() { fmt.Println("Animal Eat") }
func (a *Animal) Sleep() { fmt.Println("Animal Sleep") }
func (a *Animal) Breed() { fmt.Println("Animal Breed") }

type Dog struct {
	Animal
	// ...
}
func (d Dog) Sleep() { fmt.Println("Dog Sleep") }

type Cat struct {
	Animal
	// ...
}
func (c Cat) Sleep() { fmt.Println("Cat Sleep") }

func main() {
	dog := Dog{}
	cat := Cat{}

	dog.Eat()
	dog.Sleep()
	dog.Breed()
//	fmt.Printf("%T\n", dog)
	
	pets := []Sleeper{dog, cat}
	for _, v := range pets {
		v.Sleep()
	}
}
