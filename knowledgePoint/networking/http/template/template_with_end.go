package main
import (
	"os"
	"text/template"
)

func main() {
	t := template.New("test")

	// with 语句用来设置 dot 的值
	t, _ = t.Parse("{{with `hello`}}{{.}}{{end}}!\n")
	//t, _ = t.Parse("{{with `hello`}}{{end}}!\n")
	t.Execute(os.Stdout, nil)

	t, _ = t.Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}}{{end}}!\n")
	t.Execute(os.Stdout, nil)
}
