package main

import (
	"fmt"
	"strings"
)

func main() {
	pointer_tut()
	fmt.Println(Vertex{1,2})
	v:=Vertex{1,2}
	v.X = 4
	fmt.Println(v.X)

	//pointers and Struct
	p:=&v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1,v2,v3,p)

	arrays_tut()

	slices_tut()

	slices_tut_2()

	slice_literal()

	slice_len_cap()

	make_slice()

	slice_of_slice()

	slice_append()
	
	range_tut()
	
}

func pointer_tut(){
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Vertex struct{
	X int
	Y int
}

var(
	v1 = Vertex{1,2}
	v2= Vertex{X:1}
	v3 = Vertex{}
	p = &Vertex{1,2}

)

//arrays
func arrays_tut(){
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a[0])

	primes := [6]int{1,2,3,4,5,7}
	fmt.Println(primes)
}

//slices
func slices_tut(){
	primes:= [6]int{2,3,4,5,6,7}
	var s []int = primes[1:4]
	fmt.Println(s)
}

func slices_tut_2(){
	names := [4]string{"test","asdas","asdasd","123asdhas"}
	fmt.Println(names)

	a:= names[0:1]
	b:= names[2:3]
	fmt.Println(a,b)

	b[0]="XXXX"
	fmt.Println(a,b)
	fmt.Println(names)
}

func slice_literal(){
	q:=[]int{2,3,21}
	fmt.Println(q)

	r:=[]bool{true,false,true,true}
	fmt.Println(r)

	s:=[]struct{
		X int
		Y bool
	}{
		{2,true},
		{3,false},
		{4,true},
	}
	fmt.Println(s)
}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func slice_len_cap(){

	s:= []int{2,3,4,6,7,8}
	printSlice(s)

	//Slice the slice to give it zero len
	s = s[:0]
	printSlice(s)

	//Extend its length
	s = s[:4]
	printSlice(s)

	//Drop first two values
}

func printSlice_2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}



func make_slice(){
	a:=make([]int, 5)
	printSlice_2("a",a)

	b:=make([]int,0,5)	
	printSlice_2("b",b)

	c:=b[:2]
	printSlice_2("c",c)
	
	d:=c[2:5]
	printSlice_2("d",d)

}

func slice_of_slice(){
	//create a tic tac toe board
	board:=[][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//The player take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i:=0;i<len(board);i++{
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func slice_append() {
	var s []int
	printSlice_3(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice_3(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice_3(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice_3(s)
}

func printSlice_3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func range_tut(){
	var pow = []int{1,2,3,4,5,6,7}
	for i,v:= range pow{
		fmt.Printf("2**%d = %d\n", i, v)
	}
}