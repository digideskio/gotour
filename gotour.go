package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Good morning, xiujiao !")
	fmt.Println("Now it is ", time.Now())

	fmt.Println("My favorate number is:", rand.Intn(10))
	fmt.Println("Pi=", math.Pi)
	fmt.Println("2+4=", add(2, 4))
}

func add(x, y int) int {
	return x + y
}
