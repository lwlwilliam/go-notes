### 结构体

#### 结构体标签

结构体中的字段除了 name 和 type 之外，还有可选的 tag，该 tag 是附属于字段的了字符串，它可以当成文档或者其他重要的标注。
tag 内容不能在一般的程序中使用，只有 reflect 库可以访问它。

#### 结构体嵌套

结构体可以嵌套使用，类点类似于面向对象中的继承。当存在命名冲突时，外层的命名会覆盖内层的命名；如果冲突的命名处于同一层，
当命名被使用时会报错，如果不使用则没有影响。

#### 方法

结构体看起来就像 class 的简单形式，那么它的方法在哪里呢？Go 也有与 OO 语言方法意思基本一样的概念：Go 的方法是表示切类类型
变量行为的函数，该类型被称为接收器。方法是特殊的函数。

接收器类型几乎可以是任何类型，而不仅限于结构类型：甚至于函数类型或者 int、bool、string 或者 array
的别名类型。接收器不能是接口类型，因为接口类型是抽象定义的，而方法是具体实现，如果非要这么做，会产生编译错误：`invalid receiver type...`。

接口器也不能是指针类型，但可以是指向所有允许的类型指针。

在 Go 中的(struct)类型和它的方法的结合体等同于面向对象中的类。它们最主要的不同点是 Go 中与类型绑定的方法并不聚合在一起，这些方法可以分布
不同的源文件中，唯一的要求是它们必须在同一个 package 中。**由于方法和类型必须定义在同一个 package 中，这就是不能为 int、float 等类型定义方法
的原因**。如果尝试为 int 定义方法会产生编译错误：`cannot define new methods on non-local type int`。有一种折中的实现方法，那就是为类型(int, 
float, ...)定义别名，然后为该别名类型定义方法。或者把类型作为匿名类型嵌入 struct 中。当然，为别名类型定义的方法也只对该别名类型有效。

方法的一般形式：


```
func (recv receiver_type) methodName(parameter_list) (return_value_list) {...}
```


调用方法：`recv.methodName()`。 如果 recv 是指针，会自动解引用。如果方法不需要 recv，可以用`_`代替它：


```
func (_ receiver_type) methodName(parameter_list) (return_value_list) {...}
```

recv 类似面向对象中的`this`或者`self`，但在 Go 中并没有为它指定关键字，如果喜欢的话，可以使用 self 或者 this 作为 接收器的变量名。

#### 值接收器和指针接收器

出于性能的考虑，接收器大部分都使用指针作为接收器。

如果想让方法修改接收器指向的数据，就使用指针接收器。否则，使用值接收器为更清晰。Go 会自动检测接收器是值接收器还是指针接收器，
并且会自动转换，不需要我们特意指出。

#### 方法和非暴露字段

可以参考面向对象提供 getter 和 setter 方法令非暴露字段可以被访问。


#### 嵌套类型的方法和继承方法

当匿名类型嵌入 struct 中，该类型中的可见方法也会被嵌入，外部的类型则会继承这些方法。该机制提供了一个简单的方法来模拟
面向对象的子类继承。