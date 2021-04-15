package pkg
var helloworldslice []byte

func PrintSlice() {
	for k := range helloworldslice {
		println(helloworldslice[k])
	}
}