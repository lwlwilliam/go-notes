package basename

import "strings"

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func Basename1(s string) string {
	// Discard las '/' and everything before.
	for i := len(s) - 1; i >= 0; i -- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before las '.'.
	for i := len(s) - 1; i >= 0; i -- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// 简化版本，使用了 strings.LastIndex 库函数
func Basename2(s string) string  {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
