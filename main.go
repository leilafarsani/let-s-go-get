package main

import "fmt"


func main(){
	fmt.Println("Hello World! Hello GO!")
	sayHello("Leila")

}

func sayHello(name string){
	fmt.Println("Hello", name)
}