// 微型数学运算解释器
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	_ = iota
	MOV
	SUB
	ADD
	MUT
	DIV
	OUT
)

// 命令读取
func read() ([]byte, error) {
	if len(os.Args) == 1 {
		//return []byte("MOV x 3\nMOV y 4\nADD x y\nOUT x"), nil
		return []byte(`MOV x 45
MOV y 67
ADD x y
OUT x
SUB x y
OUT x
MUT x y
OUT x
DIV x y
OUT x
MOV u 2
MOV i 7
MUT u i
OUT u`), nil
	} else {
		// ./microCalc instruction
		f, err := os.Open(os.Args[1])
		if err != nil {
			return nil, err
		}
		defer f.Close()

		return ioutil.ReadAll(f)
	}
}

// lexer
type token struct {
	action int
	left   string
	right  string
}

var cmds []token

func line(str string) (token, bool) {
	strs := strings.Split(str, " ")

	var cmd token
	cmd.action = setcmd(strs[0])
	if cmd.action == 0 {
		return cmd, false
	}

	cmd.left = strs[1]
	if cmd.action != OUT {
		cmd.right = strs[2]
	}

	return cmd, true
}

func setcmd(str string) int {
	switch str {
	case "MOV":
		return MOV
	case "SUB":
		return SUB
	case "ADD":
		return ADD
	case "MUT":
		return MUT
	case "DIV":
		return DIV
	case "OUT":
		return OUT
	}

	return 0
}

func isNumber(str string) (int, bool) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, false
	}
	return num, true
}

var errorList []string

func parser(str string) {
	errorList = make([]string, 0)
	strs := strings.Split(str, "\n")

	cmds = make([]token, len(strs))
	for i, val := range strs {
		var rel bool

		val = strings.TrimSpace(val)
		if val == "" {
			continue
		}

		cmds[i], rel = line(val)
		if rel == false {
			var str = "Error: line " + strconv.Itoa(i+1) + ": unknown command: " + val
			errorList = append(errorList, str)
		}
	}
}

// runtime
var heap map[string]int

func runtime() {
	heap = make(map[string]int)
	for _, cmd := range cmds {
		switch cmd.action {
		case MOV:
			mov(cmd.left, cmd.right)
		case SUB:
			sub(cmd.left, cmd.right)
		case ADD:
			add(cmd.left, cmd.right)
		case MUT:
			mut(cmd.left, cmd.right)
		case DIV:
			div(cmd.left, cmd.right)
		case OUT:
			out(cmd.left)
		}
	}
}

func mov(l string, r string) {
	num, bo := isNumber(r)
	if bo == true {
		heap[l] = num
	} else {
		heap[l] = heap[r]
	}
}

func sub(l string, r string) {
	num, bo := isNumber(r)
	if bo == true {
		heap[l] = heap[l] - num
	} else {
		heap[l] = heap[l] - heap[r]
	}
}

func add(l string, r string) {
	num, bo := isNumber(r)
	if bo == true {
		heap[l] = heap[l] + num
	} else {
		heap[l] = heap[l] + heap[r]
	}
}

func div(l string, r string) {
	num, bo := isNumber(r)
	if bo == true {
		heap[l] = heap[l] / num
	} else {
		heap[l] = heap[l] / heap[r]
	}
}

func mut(l string, r string) {
	num, bo := isNumber(r)
	if bo == true {
		heap[l] = heap[l] * num
	} else {
		heap[l] = heap[l] * heap[r]
	}
}

func out(l string) {
	fmt.Println(heap[l])
}

func main() {
	code, err := read()
	if err != nil {
		fmt.Println("Fatal error: no input")
		return
	}

	parser(string(code))

	if len(errorList) != 0 {
		for _, val := range errorList {
			fmt.Println(val)
		}
		return
	}
	runtime()

	//for k, v := range heap {
	//	fmt.Println(k, ":", v)
	//}
}
