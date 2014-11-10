package main
 
import (
    "fmt"
    "time"
    "math"
    "math/rand"
)

func main(){
    fmt.Println("Good morning, xiujiao !")
    fmt.Println("Now it is ", time.Now())
    
    fmt.Println("My favorate number is:", rand.Intn(10))
    fmt.Println("Pi=",math.Pi)
}
