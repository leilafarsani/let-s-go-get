package main

import (
	"fmt"
	"time"
)



func main(){
	var gender = "female"
	fmt.Println("Hello World! Hello GO!")
	sayHello("Leila")
	fmt.Print("You are ", gender)
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

}

func sayHello(name string){
	fmt.Println("Hello", name)
}

