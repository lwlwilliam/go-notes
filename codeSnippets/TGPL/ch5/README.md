### 函数声明

函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。

```go
func name(parameter-list) (result-list) {
	body
}
```

如果函数返回一个无名变量或者没有返回值，返回值列表的括号是可以省略的。

返回值可以像形式参数一样被命名。在这种情况下，每个返回值被声明成一个局部变量，并根据该返回值的类型，将其初始化为 0。

如果一组形参或返回值有相同的类型，就不必为每个形参都写出参数类型。下面两个声明是等价的：

```go
func f(i, j, k int, s, t string)                { /* ... */ }
func f(i int, j int, k int, s string, t string) { /* ... */ }
```

下面用 4 种方法声明拥有 2 个 int 型参数和 1 个 int 型返回值的函数。blank identifier(即`_`符号)可以强调某个参数未被使用。

```go
func add(x int, y int) int   {return x + y}
func sub(x, y int) (z int)   {z = x - y; return}
func first(x int, _ int) int {return x}
func zero(int, int) int      {return 0}

fmt.Printf("%T\n", add)     // "func(int, int) int"
fmt.Printf("%T\n", sub)     // "func(int, int) int"
fmt.Printf("%T\n", first)   // "func(int, int) int"
fmt.Printf("%T\n", zero)    // "func(int, int) int"
```

函数的类型被称为函数的`标识符`。如果两个函数形式参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型和
标识符。形参和返回值的变量名不影响函数标识符也不影响它们是否可以以省略参数类型的形式表示。

在函数调用时，Go 语言没有默认参数值，也没有任何方法可以通过参数名指定形参，因此形参和返回值的变量名对于函数调用者而言没有意义。

在函数体中，函数的形参作为局部变量，被初始化为调用者提供的值。函数的形参和有名返回值作为函数最外层的局部变量，被存储在相同的词法块中。

实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但是，如果实参包括引用类型，如指针、slice、map、
function、channel 等类型，实参可能会由于函数的间接引用被修改。

可能会偶尔遇到没有函数体的函数声明，这表示该函数不是以 Go 实现的。这样的声明定义了函数标识符。

```go
package math

func Sin(x float64) float64 // implemented in assembly language
```
