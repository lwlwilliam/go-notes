package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
	"regexp"
)

// 之所以用 []byte 而不是 string 来存储 Body，是因为 io 库使用的是这种类型
type Page struct {
	Title string
	Body []byte
}

// 模板缓存：每次渲染页面时，renderTemplate 都会调用 ParseFiles；
// 更好的方法是在程序初始化时就调用 ParseFiles 一次，把所有模板解析到 *Template，然后就可以使用 ExecuteTemplate 方法渲染指定模板；
// template.Must 是一个包装器，当传入非空的错误值时，会引起 panic 错误，否则返回 *Template；
// ParseFiles 可以传入任意数量标识模板文件的字符串参数，然后解析文件到 templates；
// 如果要添加更多的模板到程序中，就把模板名添加到 ParseFile 参数中
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// 程序有一个严重的瑕疵：用户可以使用任意路径在服务器上读写；
// 为了减轻这种情况，可以用正则表达式验证 title；
// 创建一个全局变量保存验证表达式；
// regexp.MustCompile 方法会解析编译正则表达式，返回 regexp.Regexp；
// MustCompile 与 Compile 的区别是：如果表达式错误，会导致 panic 错误，而 Compile 会返回一个作为第二个参数的 error 值
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// 通过 URL 路径获取 title
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	// 验证 URL 路径，提取 title
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil
}

// Page 描述了页面数据在内存中是如何存储的，然而怎么永久存储？
// 可以用 save 方法把页面的 Body 部分保存到 text 文件中，文件名为 Title
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)  // 往文件写入 []byte 类型的数据
}

// 除了保存页面外，还想加载页面
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 允许用户查看 wiki 页面。处理以 "/view/" 为前缀的 URL
// 从 r.URL.Path 提取页面 Title
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	// 如果请求的页面不存在，重定向到编辑页
	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// 加载页面（如果不存在就创建空页面结构），并显示 HTML 表单
func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

// 测试
func main() {
	//var page = Page{Title: "test", Body: []byte("This is just a test.")}
	//page.save()
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
