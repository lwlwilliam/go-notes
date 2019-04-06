package main

import "fmt"

func reverse(s []int)  {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(s)
	reverse(s)
	fmt.Println(s)

	s2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(s2[:3])
	reverse(s2[3:])
	reverse(s2)
	fmt.Println(s2)

	str := "abcdefghijklmnopqrstuvwxyz"
	str1 := str[:3]
	str2 := str[:4]
	fmt.Println(&str, &str1, &str2)

	s3 := []int{0, 1, 2, 3, 4, 5}
	s4 := s3[:3]
	s5 := s3[:4]
	fmt.Println(&s3[0], &s4[0], &s5[0])
}
