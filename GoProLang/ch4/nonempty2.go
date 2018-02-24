package main

import "fmt"

func nonempty2(strings []string) []string {
    out := strings[:0]
    for _, s := range strings {
        if s != "" {
            out = append(out, s)
        }
    }
    return out
}

func main() {
    var strings = []string{"one", "", "three"}

    fmt.Println(strings)
    strings = nonempty2(strings)
    fmt.Println(strings)
}