package main
import (
"fmt"
"time"
)
func f(n int) {
	for i := 0; i < 5; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(3000)
		time.Sleep(time.Millisecond * amt)
	}
}
func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}