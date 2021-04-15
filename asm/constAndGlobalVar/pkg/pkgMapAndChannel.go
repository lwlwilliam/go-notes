package pkg

var m map[string]int
var ch chan int

func PrintMapChannel() {
	for k := range m {
		println(m[k])
	}

	println()
	println(ch)
}
