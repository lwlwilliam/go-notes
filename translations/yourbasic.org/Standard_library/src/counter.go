package main

import (
	"bufio"
	"strings"
	"fmt"
)

func main() {
	const input = `Beware of bugs in the above code;
I have only proved it correct, not tried it.`

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count ++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(count)
}
