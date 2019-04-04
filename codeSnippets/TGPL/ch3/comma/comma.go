package comma

import "bytes"

// comma inserts comma
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return Comma(s[:n-3]) + "," + s[n-3:]
}

// practice3.10
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
