package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)



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

