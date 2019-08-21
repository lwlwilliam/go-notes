### map 概述

#### 基础

map 是键值对的无序集合，键具有唯一性。

```
var m map[string]int                // m == nil, len(m) == 0
m1 := make(map[string]float64)      // empty map of string-float64 pairs
m2 := make(map[string]float64, 100) // preallocate room for 100 entries
m3 := map[string]float64{
    "e": 2.71828,
    "pi": 3.1416,
}

fmt.Println(len(m1), len(m2), len(m3))  // 0 0 2
```

*   map 的默认零值是 nil。一个 nil map 除了不能添加任何元素外，等价于一个空的 map。
*   可以通过字面量或者调用 make 函数来创建 map，make 创建时容量参数是可选的。
*   内置的 len 函数可以查询键值对的数量。

#### 添加、查找和删除

[map.go](../src/map.go)

#### 迭代

注意，这里打印出来是无序的。

[map2.go](../src/map2.go)

*   迭代的顺序是不确定的，每次迭代可能都不一样。
*   如果在迭代期间删除了尚未到达的条目，则不会生成相应的迭代值（译注：也就是说 map 类型不是并发安全的）。
*   如果在迭代期间创建了条目，则在迭代期间可能生成该条目，也可能不生成该条目（译注：与上一条一样，说明了 map 类型不是并发安全的）。

#### 实现细节

*   map 底层是 [hash table](https://yourbasic.org/algorithms/hash-tables-explained/) 数据结构。
*   在恒定的开销时间内提供了查询、插入和删除操作（TODO: 这翻译怪怪的，有待修改）。
*   必须为 key 类型定义 == 和 != 比较操作符。
