package main
import (
	"os"
	"text/template"
	"fmt"
)

type Person struct {
	Name	string
	Age		int
}

func main() {
	t := template.New("hello")
	t = t.Delims("--", "==")
	t, _ = t.Parse("My name is --.Name==. I'm --.Age== years old.\n")
	p := Person{Name: "William", Age: 20}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}