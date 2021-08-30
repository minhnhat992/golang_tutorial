// tutorial for loop

package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println(sum)

	for_tut()

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(pow(3, 3, 10),
		pow(3, 2, 10))

	switch_tut()

	switch_tut_2()

	switch_no_condition()

	defer_tut()

	defer_stack_tut()
}

//The init and post statements are optional.
func for_tut() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func switch_tut() {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS x.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switch_tut_2() {
	fmt.Println("When's Monday?")
	today := time.Now().Weekday()
	switch time.Monday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In 2 days")
	default:
		fmt.Println("Too far away")
	}
}

func switch_no_condition(){
	t:= time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func defer_tut(){
	defer fmt.Println("world")
	fmt.Println("Hello")
}

func defer_stack_tut(){
	fmt.Println("counting")
	for i:=0; i<10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}