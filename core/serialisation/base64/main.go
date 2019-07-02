package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {
	eightBitData := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	buff := &bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, buff)
	encoder.Write(eightBitData)
	encoder.Close()
	fmt.Println(buff.String())

	dbuff := make([]byte, 12)
	decoder := base64.NewDecoder(base64.StdEncoding, buff)
	decoder.Read(dbuff)
	for _, ch := range dbuff[:len(eightBitData)] {
		fmt.Print(ch)
	}
}
