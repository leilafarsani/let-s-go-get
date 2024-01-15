package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
	"math/cmplx"
)

var c, python, java bool

func main(){
	var gender = "female"
	fmt.Println("Hello World! Hello GO!")
	sayHello("Leila")
	fmt.Print("You are ", gender)
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(17))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var l int
	fmt.Println(l, c, python, java)
	var i, j int = 1, 2
	//var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java, p)
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	zero()
	typing()
}



func sayHello(name string){
	fmt.Println("Hello", name)
}


func add(x int, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var l, p int  = 9 , 11

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func zero() {
	var n int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", n, f, b, s)
}

func typing() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}