// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty return a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
    i := 0
    for _, s := range strings {
        if s != "" {
            strings[i] = s
            i ++
        }
    }
    return strings[:i]
}

func main() {
    var strings = []string{"one", "", "three"}
    // var result []string
    // result = nonempty(strings)
    // fmt.Println(result)

    // 通常这样使用
    strings = nonempty(strings)
    fmt.Println(strings)
}