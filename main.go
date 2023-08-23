package main

import "fmt"


func main(){
	var gender = "female"
	fmt.Println("Hello World! Hello GO!")
	sayHello("Leila")
	fmt.Print("You are ", gender)

}

func sayHello(name string){
	fmt.Println("Hello", name)
}