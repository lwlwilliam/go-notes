### 无继承的无痛面向对象编程

> 原文：[https://yourbasic.org/golang/inheritance-object-oriented/](https://yourbasic.org/golang/inheritance-object-oriented/)

**Go 没有继承——取而代之的是以接口和嵌套支持代码重用和多态**。

继承是传统面向对象的语言提供的三个特性之一。当 Dog 继承自 Animal 时：

1.	Dog 能重用 Animal 的代码；
2.	Animal 类型的变量 x 既可以被 Dog 引用，也可以被 Animal 引用；
3.	x.Eat() 取决于对象 x 引用的是哪种类型而调用对应的 Eat 方法；

在面向对象的术语中，这些特性被称为代码复用、聚合和动态调度。

所以的这些在 Go 中都可使用，使用的是单独的构造。

*	**组合** 和 **嵌套** 提供代码复用；
*	**接口** 负责多态和动态调度；

#### 代码重用：组合

**启动 Go 项目后不用担心类型层次结构，在以后引入多态和动态调度是很容易的**。

如果 Dog 需要 Animal 的部分或者全部功能，只需使用 **组合**。

```
type Animal struct {
	// ...
}

type Dog struct {
	beast Animal
	// ...
}
```

这让你可以自由地根据需要把 Animal 看作是 Dog 的一部分来使用。是的，就是这么简单。

#### 代码复用：嵌套

如果 Dog 继承 Animal 特定的行为，这种方法会导致代码的冗余。

```
type Animal struct {
	// ...
}

func (a *Animal) Eat() {...}
func (a *Animal) Sleep() {...}
func (a *Animal) Breed() {...}

type Dog struct {
	beast Animal
	// ...
}

func (a *Dog) Eat() { a.beast.Eat() }
func (a *Dog) Sleep() { a.beast.Sleep() }
func (a *Dog) Breed() { a.beast.Breed() }
```

这种代码模式就叫 **委托模式**。

这种情况下，Go 使用嵌套。Dog 的声明和它的三个方法可以简化为：

```
type Dot struct {
	Animal
	// ...
}
```

#### 多态和动态调度：接口

**接口保持简短，只有在需要时才采用它们**。

接下来你的项目可能会扩展并包含更多的动物（译者注：关于动物的代码）。这时候你可以采用接口实现多态和动态调度。

```
type Sleeper interface {
	Sleep()
}

func main() {
	pets := []Sleeper{new(Cat), new(Dog)}
	for _, x := range pets {
		x.Sleep()
	}
}
```

没有必要的 Cat 和 Dog 类型的显式声明。任何提供了接口中所有方法的类型都被为实现了该接口。

当看见一只鸟像一只鸭子一样行走和游泳并且叫起来像一只鸭子，就把这只鸟叫做鸭。——James Whitcomb Riley

#### 更多阅读

阅读[Constructors](https://yourbasic.org/golang/constructor-best-practice/)，看看 Go 创建数据结构的最佳实践。

source code: [oop.go](../src/oop.go)
