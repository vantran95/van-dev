package main

import (
	"fmt"
	"time"
)

func main() {
	//User does: Print,then increase counter...
	//GetInstance().Increment()

	time.Sleep(1 * time.Second)

	//GetInstance().Increment()

	fmt.Println(GetInstance().GetCounter())
}
