### 各种复合数据类型的声明和初始化方法

> 数组

1.  先声明，后初始化
    ```go
    var arr [3]int
    arr[0] = 1
    arr[1] = 2
    arr[2] = 3
    ```
    
2.  声明并初始化
    ```go
    var arr [3]int = [3]int{1, 2, 3}
    ```
    
3.  声明并初始化，数据类型由 Go 语言判断 
    ```go
    var arr = [3]int{1, 2, 3}
    ```
    
4.  `...`表示根据实际元素个数来设置数组大小 
    ```go
    var arr = [...]int{1, 2, 3}
    ```
    
5. 这种只能在函数内部使用 
    ```go
    arr := [3]int{1, 2, 3}
    ```
    
> slice

1.  先声明，后初始化 
    ```go
    var slice []int
    slice = make([]int, 3, 3)
    ```
    
2.  声明并初始化
    ```go
    var slice []int = make([]int, 3, 3)
    ```
    
3.  由 Go 判断数据类型
    ```go
    var slice = make([]int, 3, 3)
    ```
    
4.  只能在函数内部使用
    ```go
    slice := make([]int, 3, 3)
    ```
    
5.  只能在函数内部使用，跟 4 的区别就是没有用 make 函数（其实貌似创建 slice 不一定非要用 make 函数，以上都可以把 make 去掉）
    ```go
    slice := []int{1, 2, 3}
    ```

> map

map 也可以用 make 函数创建：`var m = make(map[string]int)`。

1.  先声明，后初始化
    ```go
    var m map[string]int
    m = map[string]int{"a":1, "b":2, "c":3}
    ```
    
2.  声明并初始化
    ```go
    var m map[string]int = map[string]int{"a":1, "b":2, "c":3}
    ```
    
3.  由 Go 判断数据类型
    ```go
    var m = map[string]int{"a":1, "b":2, "c":3}
    ```
    
4.  只能在函数内部使用
    ```go 
    m := map[string]int{"a":1, "b":2, "c":3}
    ```
    
> struct

```go
type S struct {
    a int
    b string
    // ...
}
```

1.  先声明，后初始化 
    ```go
    var s S
    s = S{1, "test"}
    ```
    
2.  声明并初始化
    ```go
    var s S = S{1, "test"}
    ```
    
3.  由 Go 判断类型
    ```go
    var s = S{1, "test"}
    ``` 
    
4.  只能在函数内部使用
    ```go
    s := S{1, "test"}
    ```
