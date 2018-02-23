package main

import (
	"fmt"
)

func main() {
	months := [...]string{"January", "February", "March", "April", "May",
	                      "June", "July", "August", "September", "October",
	                      "November", "December"}

	summer := months[6:9]
	fmt.Println(summer)
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
	fmt.Println(summer)
}
