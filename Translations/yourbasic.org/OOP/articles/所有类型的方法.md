### 所有类型的方法

> 原文：[https://yourbasic.org/golang/methods-explained/](https://yourbasic.org/golang/methods-explained/)

**在类型定义中声明的任何类型都可以有附加的方法**。

*	方法是拥有接收器参数的函数；
*	接收器位于关键词`func`和方法名之间；
*	可以为任何在类型定义中声明的类型定义方法；

在该示例中，Value 方法与 MyType 关联。方法接收器称为 p。

```
type MyType struct {
	n int
}

func (p *MyType) Value() int { return p.n }

func main() {
	m := new(MyType)
	fmt.Println(pm.Value())		// 0 (零值)
}
```

如果把值转换为其他类型，新的值会拥有新类型的方法，而不是旧类型的。

```
type MyInt int

func (m MyInt) Positive() bool { return m > 0 }

func main() {
	var m MyInt = 2
	m = m * m	// 基本类型的运算符仍然适用

	fmt.Println(m.Positive())			// true
	fmt.Println(MyInt(-1).Positive())	// false

	var n int
	n = int(m)	// 需要转换
	n = m		// 非法
}
```

```
../main
# command-line-arguments
../src/method2.go:20:4: cannot use m (type MyInt) as type int in assignment
```

source code: [method1.go](../src/method1.go)、[method2.go](../src/method2.go)
