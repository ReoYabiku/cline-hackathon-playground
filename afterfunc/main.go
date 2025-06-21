package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("the start of the program...")

	time.AfterFunc(time.Second, func () {
		fmt.Println("inside the afterfunc")
	})

	fmt.Println("the end of the program...")
}
