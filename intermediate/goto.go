package intermediate

import (
	"fmt"
	"math/rand"
)

// goto is also a controlled statement in Go
// it should be avoided when necessary however it has its use cases
// Edgar Dijkstra

/* func IllegalGoTo() {
	a := 10
	goto skip
	b := 20
skip:
	c := 30
	fmt.Println(a, b, c)
	if c > a {
		goto inner
	}
	if a < b {

	inner:
		fmt.Println("%v is less than %v", a, b)
	}

}
*/

// * goto is an option to make the code more readable

func LegalGoto() {
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("Do something when the loop completes normally.")
done:
	fmt.Println("Do some complicated logic no matter why we left the loop.")
}
