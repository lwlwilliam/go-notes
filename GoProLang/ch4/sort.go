package main

import (
    "sort"
    "fmt"
)

func main() {
    var ages = map[string]int{
        "charlie": 34,
        "alice": 31,
    }
    var names []string

    for name := range ages {
        names = append(names, name)
    }

    sort.Strings(names)
    for _, name := range names {
        fmt.Printf("%s\t%d\n", name, ages[name])
    }
}