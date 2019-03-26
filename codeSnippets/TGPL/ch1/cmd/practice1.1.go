// 能够打印 os.Args[0]
package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	fmt.Println(strings.Join(os.Args, " "))
}
