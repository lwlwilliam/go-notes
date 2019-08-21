### 接口概述

> 原文：[https://yourbasic.org/golang/interfaces-explained/](https://yourbasic.org/golang/interfaces-explained/)

#### 基础

**接口类型由一系列方法签名组成。接口变量类型可以保存任何实现了接口方法的值。**

以下例子的 Temp 和`*Point` 都实现了 MyStringer 接口。

[basicInterface.go](../src/basicInterface.go)

#### 结构类型

**一个类型以实现接口方法的形式来实现某个接口。而不需要显式声明（译注：类型实现了接口）**

事实上，上例的 Temp，`*Temp`和`*Point`类型同样实现了标准库的`fmt.Stringer`接口。接口中的 String 方法被当作操作数传到函数中进行打印，例如
fmt.Println 函数。

#### 空接口

没有指定方法的接口类型就是空接口。

```
interface{}
```

空接口可以保存任何类型的值，因为任何类型至少都实现了零个方法。

[emptyInterface.go](../src/emptyInterface.go)

fmt.Println 函数就是一个主要的例子，它可以接受任何类型的参数。

```
func Println(a ...interface{}) (n int, err error)
```

#### 接口值

**接口值由具体的值和动态的类型组成：[Value, Type]。**

调用 fmt.Printf 时，可以使用 %v 打印具体的值以及 %T 打印动态类型。

[interfaceValue.go](../src/interfaceValue.go)

接口类型的零值是 nil，可以用 [nil, nil] 来表示。

nil 接口上调用方法是一个运行时错误，然而，编写能够处理接收器值 [nil, Type] 的方法很常见，其中 Type 的类型不为 nil。

当然，也可以使用类型断言，类型切换和反射来访问接口值的动态类型。

#### 相等性

符合以下条件之一时两个接口值相等。

1.  拥有相同的具体值以及一致的动态类型；
2.  两者都为 nil；

[equalityInterface.go](../src/equalityInterface.go)

上例中，`x = (*Point)(nil)`语句 x 的具体值是 nil，但是它的动态类型是`*Point`。
