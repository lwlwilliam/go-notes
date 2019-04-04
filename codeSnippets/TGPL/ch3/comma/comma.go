package comma

import (
	"bytes"
	"strconv"
	"strings"
)

// comma inserts comma
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return Comma(s[:n-3]) + "," + s[n-3:]
}

// practice3.10
// 编写一个非递归版本的 comma 函数，使用 bytes.Buffer 代替字符串链接操作
func Comma2(s string) string  {
	var concat, res bytes.Buffer
	for i, border := len(s) - 1, len(s); i >= 0; i -- {
		concat.WriteByte(s[i])

		if (border - i) % 3 == 0 && i != 0 {
			concat.WriteByte(',')
		}
	}

	for i := len(concat.String()) - 1; i >= 0; i -- {
		res.WriteByte(concat.String()[i])
	}

	return res.String()
}

// practice3.11
// 完善 comma 函数，以支持浮点数处理和一个可选的正负号的处理
func Comma3(s string) string {
	dots := bytes.Count([]byte(s), []byte("."))

	if dots == 1 {
		parts := strings.Split(s, ".")
		isNum := true

		for _, p := range parts {
			if _, err := strconv.Atoi(p); err != nil {
				isNum = false
				break
			}
		}

		if isNum {
			return Comma(parts[0]) + "." + parts[1]
		} else {
			return Comma(s)
		}
	} else {
		return Comma(s)
	}
}