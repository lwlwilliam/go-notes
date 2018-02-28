### 方法

从 90 年代早期开始，面向对象编程(OOP)就成为了称霸工程界和教育界的编程范式，所以之后几乎所有大规模被应用的语言都包含了对 OOP 的支持，Go
语言也不例外。

尽管没有被大众所接受的明确的 OOP 的定义，从我们的理解来讲，一个对象其实也就是一个简单的值或者一个变量，在这个变量中会包含一些方法，而
一个方法则是一个和特殊类型关联的函数。一个面向对象的程序会用方法来表达其属性和对应的操作，这样使用这个对象的用户就不需要直接去操作对象，
而是借助方法来做这些事情。

标准库提供的一些方法：time.Duratin 类型的 Seconds 方法：

```go
const day = 24 * time.Hour
fmt.Println(day.Seconds())  // "86400"
```

自定义方法：Celsius 类型的 String 方法：

```go
func (c Celsius) String() string {
	return fmt.Sprintf("%g", c)
}
```

>> 方法声明

在函数声明时，在其名字之前放上一个变量，即是一个**方法**。这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的
方法。

[geometry.go](geometry.go)

上面的代码里那个附加的参数 p，叫做方法的**接收器(receiver)**，早期的面向对象语言留下的遗产将调用一个方法称为"向一个对象发送消息"。

在 Go 语言中，我们并不会像其它语言那样用 this 或者 self 作为接收器；我们可以任意的选择接收器的名字。由于接收器的名字经常会被使用到，所
以保持其在方法间传递时的一致性和简短性是不错的主意。这里的建议是可以使用其类型的第一个字母。在方法调用过程中，接收器参数一般会在方法名
之前出现。这和方法声明是一样的，都是接收器参数在方法名字之前。

```go
p := Point{1, 2}
q := Point{4, 6}
fmt.Println(Distance(p, q))  // "5", function call
fmt.Println(p.Distance(q))   // "5", method call
```

可以看到，上面的两个函数调用都是 Distance，但是却没有发生冲突。第一个 Distance 的调用实际上用的是包级别的函数 geometry.Distance，而第
二个则是使用刚刚声明的 Point，调用的是 Point 类下声明的 Point.Distance 方法。这种 p.Distance 的表达式叫做**选择器**，因为它会选择合适
的对应 p 这个对象的 Distance 方法来执行。选择器也会被用来选择一个 struct 类型的字段，比如 p.X。由于方法和字段都是在同一命名空间，所以
如果在这里声明一个 X 方法的话，编译器会报错，因为在调用 p.X 时会有歧义（这里有点奇怪）。

因为每种类型都有其方法的命名空间，在用 Distance 这个名字的时候，不同的 Distance 调用指向了不同类型里的 Distance 方法。

```go
// A path is a journey connecting the points with straight lines.
type Path []Point
// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
```

Path 是一个命名的 slice 类型，而不是 Point 那样的 struct 类型，然而依然可以为它定义方法。

> 基于指针对象的方法

函数的其中一个参数实在太大，希望能避免进行这种黑夜的拷贝，这种情况下就需要用到指针了。

```go
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
```

这个方法的名字是 **(*Point).ScaleBy**。这里的括号是必须的；没有括号的话这个表达式可能会被理解为**\*(Point.ScaleBy)**。

为了避免歧义，在声明方法时，如果一个类型名本身是一个指针的话，是不允许其出现在接收器中的，比如：

```go
type P *int
func (P) f() { /*...*/ }  // compile error: invalid receiver type
```

### TODO: 未完待续


