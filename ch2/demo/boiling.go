// Boiling prints the boilign point of water.
package main

import "fmt"

const boilingF = 212.0

func main() {
    var f = boilingF
    var c = (f - 32) * 5 / 9
    fmt.Printf("boiling point = %g℉ or %g℃\n", f, c)
    // Output:
    // boiling point = 212℉ or 100℃
}