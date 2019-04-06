# 复合数据类型

复合数据类型，它是以不同的方式组合基本类型可以构造出来的复合数据类型。数组和结构体是聚合类型；它们的
值由许多元素或成员字段的值组成。数组是由同构的元素组成——每个数组元素都是完全相同的类型——结构体则是
由异构的元素组成的。数组和结构体都是有固定内存大小的数据结构。相比之下，slice 和 map 则是动态的
数据结构，它们将根据需要动态增长。

### 数组

数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。因为数组的长度是
固定的，因此在 Go 语言中很少直接使用数组。和数组对应的类型是 Slice（切片），它是可以增长和收缩
的动态序列，slice 功能也更灵活，但是要理解 slice 工作原理的话需要先理解数组。

默认情况下，数组的每个元素都被初始化为元素类型对应的零值，对于数字类型来说就是 0。也可以使用数组
字面值语法用一组值来初始化数组：

```
var q [3]int = [3]int{1, 2, 3}
var r [3]int = [3]int{1, 2}
fmt.Println(r[2])  // "0"
```

在数组字面值中，如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来
计算。因此，上面 q 数组的定义可以简化为：

```go
q := [...]int{1, 2, 3}
fmt.Printf("%T\n", q)  // "[3]int"
```

数组的长度是数组类型的一个组成部分，因此 [3]int 和 [4]int 是两种不同的数组类型。数组的长度必须
是常量表达式，因为数组的长度需要在编译阶段确定。

```
q := [3]int{1, 2, 3}
q = [4]int{1, 2, 3, 4}  // compile error: cannot assign [4]int to [3]int
```

以上的形式是直接提供顺序初始化值序列，但是也可以指定一个索引和对应值列表的方式初始化，如下：

```go
type Currency int

const (
	USD Currenty = iota     // 美元
	EUR                     // 欧元
	GBP                     // 英镑
	RMB                     // 人民币
)

symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

fmt.Println(RMB, symbol[RMB])
```

如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的。只有当两个数组的所有
元素都是相等的时候数组才是相等的。

```
a := [2]int{1, 2}
b := [...]int{1, 2}
c := [2]int{1, 3}
fmt.Println(a == b, a == c, b == c) // true, false, false
d := [3]int{1, 2}
fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
```

### Slice

Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。slice 的语法和数组很像，只是没有
固定长度而已。

数组和 slice 之间有着紧密的联系。一个 slice 是一个轻量级的数据类型，提供了访问数组子序列（
或者全部）元素的功能，而且 slice 的底层确实引用一个数组对象。一个 slice 由三个部分构成：
指针，长度和容量。指针指向第一个 slice 元素对应的底层数组元素的地址，要注意的是 slice 的第
一个元素并不一定就是数组的第一个元素。长度对应 slice 中元素的数目；长度不能超过容量，容量
一般是从 slice 的开始位置到底层数组的结尾位置。

多个 slice 之间可以共享底层的数据，并且引用的数组部分区间可能重叠。

```
months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 
                    5: "May", 6: "June", 7: "July", 8: "August", 9: "September",
                    10: "October", 11: "November", 12: "December"}
```

以上数组第 0 个元素为空字符串。下面定义表示第二季度和北方夏天月份的 slice，它们有重叠部分：

```go
Q2 := months[4:7]
summer := months[6:9]
fmt.Println(Q2)     // ["April" "May" "June"]
fmt.Println(summer) // ["June" "July" "August"]
```

两个 slice 之间的关系如下图：

![slice](assets/20190406213449.png)

如果切片操作超出 cap(s) 的上限将导致一个 panic 异常，但是超出 len(s) 则是意味着扩展了
slice，因为新 slice 的长度会变大：

```go
fmt.Println(summer[:20])    // panic: out of range

endlessSummer := summer[:5] // extend a slice (within capacity)
fmt.Println(endlessSummer)  // "[June July August September October]"
```

字符串的切片操作和 []byte 字节类型切片的切片操作是类似的。都写作 x[m:n]，并且都是返回一个
原始字节系列的子系列，底层都是共享之前的底层数组，因此这种操作都是常量时间复杂度。x[m:n]
切片操作对于字符串则生成一个新字符串，如果 x 是 []byte 的话则生成一个新的 []byte。

