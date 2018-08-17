package main
import (
	"os"
	"text/template"
)

func main() {
	t := template.New("test")
	// 模板变量 $，可以在模板里创建局部变量，使用 $ 前缀即可。变量名必须由字母、数字和下划线构成。
	t = template.Must(t.Parse("{{with $3 := `hello`}}{{$3}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)

	t = template.Must(t.Parse("{{with $x3 := `hola`}}{{$x3}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)

	t = template.Must(t.Parse("{{with $x_1 := `hey`}}{{$x_1}} {{.}} {{$x_1}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)
}
