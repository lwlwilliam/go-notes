package printints

import (
	"bytes"
	"fmt"
)

func IntsToString(values []byte) string  {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
		buf.Write(values)
	}
	buf.WriteByte(']')

	//var res = make([]byte, 5)
	//n, _ := buf.Read(res)
	//fmt.Println(res, n)

	return buf.String()
}
