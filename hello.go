package main // import "github.com/nicolaferraro/vgo-test"

import (
	"fmt"
	"time"
	"rsc.io/quote"
)

func main() {

	fmt.Println(quote.Hello())

	flows := NewFlows()

	ch := flows.Timer(2 * time.Second)
	i := <- ch

	fmt.Println("First message: ", i)
}

