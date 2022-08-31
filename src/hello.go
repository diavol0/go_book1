package main

import "fmt"

func main() {
	var line string
	fmt.Println("Hello world , go!")
	fmt.Println("What's your name?")
	fmt.Scanf("%s\n", &line)
	fmt.Printf("Привет , %s!\n", line)

}
