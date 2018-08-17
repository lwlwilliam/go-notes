package main
import (
	"text/template"
	"fmt"
)

func main() {
	tOk := template.New("ok")

	// a valid template, so no panic with Must:
	template.Must(tOk.Parse("/* and a comment */ some static text: {{ .Name}}"))
	fmt.Println("The first one parse OK.")
	fmt.Println("The next one ought to fail.")
	//tErr := template.New("error_template")
	//template.Must(tErr.Parse(" some static text {{ .Name }"))
}
