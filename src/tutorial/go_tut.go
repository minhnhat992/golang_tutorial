package main

import (
	"bytes"
	"fmt"
	"math"
	"golang_turorial/src/github.com/einsteinis/go_tutorial/string_util"


)

func main() {
	// bytes_tutorial()
	// use_math_pck()
	fmt.Println(string_util.Reverse("reverse me"))
	fmt.Println(add(5,6))
	var fruits [2]string
	fruits[0] = "banana"
	fruits[1] = "pineapple"
	fmt.Println(fruits)
}

// func main() {
// 	my_byte := byte('b')
// 	my_rune := '~'
// 	fmt.Printf("%c=%d %c=%U\n", my_byte, my_rune, my_rune, my_rune)
// }

func bytes_tutorial() {
	b1 := byte('a')
	b2 := []byte("A")
	b3 := []byte{'a', 'b', 'c'}
	fmt.Printf("b1 = %c\n", b1)
	fmt.Printf("b2 = %c\n", b2)
	fmt.Printf("b3 = %s\n", b3)
	s1 := []byte("Hello")
	s2 := []byte("World")
	s3 := [][]byte{s1, s2}
	s4 := bytes.Join(s3, []byte(","))
	s5 := []byte{}
	s5 = bytes.Join(s3, []byte("--"))
	s6 := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("s1 = %s\n", s1)
	fmt.Printf("s2 = %s\n", s2)
	fmt.Printf("s3 = %s\n", s3)
	fmt.Printf("s4 = %s\n", s4)
	fmt.Printf("s5 = %s\n", s5)
	fmt.Printf("%s\n", bytes.Join(s6, []byte(", ")))
}

func use_math_pck() {
	c := true
	fmt.Println(c, math.Floor(1.7), math.Ceil(10.5))
}

func add(x,y int) int {
	return x + y
}

