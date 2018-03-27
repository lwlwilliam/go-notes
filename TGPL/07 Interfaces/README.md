### 接口

接口类型是对其他类型行为的抽象和概括；因为接口类型不会和特定的实现细节绑定在一起，通过这种抽象的方式让函数更加灵活和更具有适应能力。

Go 语言中接口类型的独特之处在于它是满足隐式实现的。也就是说，没有必要对于给定的具体类型定义所有满足的接口类型；简单地拥有一些必需的方法就足够了。
这种设计可以让你创建一个新的接口类型满足已经存在的具体类型却不会去改变这些类型的定义；当使用的类型来自于不受我们控制的包时这种设计尤其有用。

> 接口约定

一个具体的类型可以准确地描述它所代表的值，并且展示出对类型本身的一些操作方式：就像数字类型的算术操作，切片类型的取下标、添加元素和范围获取操作。
具体的类型还可以通过它的内置方法提供额外的行为操作。总的来说，当拿到一个具体的类型时就知道它的本身是什么和可以用它来做什么。

Go 语言中还存在另一种类型：接口类型。接口类型是一种抽象的类型。它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；它们只
会展示出它们自己的方法。也就是说当看到一个接口类型的值时，并不知道它是什么，唯一知道的就是可以通过它的方法来做什么。

有两个相似的函数可以进行字符串的格式化：fmt.Printf（把结果写到标准输出） 和 fmt.Sprintf（把结果以字符串形式返回）。得益于使用接口，不必因为返
回结果在使用方式上的一些浅显不同就把格式化这个最困难的过程复制一份。实际上，这两个函数都使用了另一个函数 fmt.Fprintf 来进行封装。