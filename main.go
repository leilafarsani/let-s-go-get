package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
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