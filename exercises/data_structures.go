package exercises

import (
	"fmt"
	"maps"
)

func ComparingMaps() {
	m := map[string]int{
		"hello": 5,
		"world": 25,
	}
	n := map[string]int{
		"world": 25,
		"hello": 5,
	}

	fmt.Println(maps.Equal(m, n))

	// deleting from maps
	delete(m, "hello")
	delete(n, "world")

	// emptying a map
	fmt.Println(n, len(n))
	clear(n)
	fmt.Println(n, len(n))

	// the comma ok idiom
	v, ok := m["king"]
	fmt.Println(v, ok)

	// * reading and writing from maps is similar to Python dicts
	scores := make(map[string]int)
	scores["chelsea"] = 6
	scores["arsenal"] = 6
	scores["arsenal"]++
	scores["city"] = 6
	scores["united"] = 6

	fmt.Println(scores)
	fmt.Println(scores["arsenal"])

	// todo : hash maps
}
