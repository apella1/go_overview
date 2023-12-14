package main

import (
	"errors"
	"fmt"
)

// maps are similar to python dicts
func maps() {
	// * using the make function
	ages := make(map[string]int)
	ages["Bernard"] = 98
	ages["Kaitlin"] = 56

	// * using the literal
	teamManagers := map[string]string{
		"Manchester United": "Eric Ten Hag",
		"Manchester City":   "Pep Guardiola",
		"Liverpool":         "Jurgen Klopp",
	}
	fmt.Println(teamManagers)
	fmt.Println(len(teamManagers))
}

type user struct {
	name        string
	phoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := map[string]user{}
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		name, phoneNumber := names[i], phoneNumbers[i]
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil
}

type newUser struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

// ! maps are passed by reference to functions
// ! we don't have a copy thus can modify the original map
func deleteIfNecessary(users map[string]newUser, name string) (deleted bool, err error) {
	// * shortened syntax for below
	/* if _, ok := users[name]; !ok {
		return false, errors.New("not found")
	} */

	user, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}

	if user.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}

	return false, nil
}
