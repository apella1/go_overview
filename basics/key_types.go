package main

import "fmt"

type Key struct {
	Path, Country string
}

// map[KeyType]ValueType
func keyTypes() {
	// * string key -> (string key -> int value)
	// ! map keys must be comparable
	hits := make(map[string]map[string]int)
	fmt.Println(hits)

	myMap := make(map[Key]int)
	fmt.Println(myMap)
}

// ! slices, arrays are not comparable as they are just pointers to memory addresses

func getCounts(userIDs []string) map[string]int {
	userFreq := make(map[string]int)

	// * range can also be used to loop through the user ids
	// * range works similarly to Python's enumerate()
	for i := 0; i < len(userIDs); i++ {
		_, ok := userFreq[userIDs[i]]
		if !ok {
			userFreq[userIDs[i]] = 0
		}
		userFreq[userIDs[i]]++

	}
	return userFreq
}

// accessing values in maps is a safe call

func nestedMaps() {
	// * int -> (string -> (string -> string))
	firstMap := make(map[int]map[string]map[string]string)
	fmt.Printf("%T\n", firstMap)
}

// * individual characters are of type rune

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		char := rune(name[0])
		_, ok := nameCounts[char]
		if !ok {
			nameCounts[char] = make(map[string]int)
		}
		nameCounts[char][name]++
	}
	return nameCounts
}
