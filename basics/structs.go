package main

import "fmt"

// ! structs are comparable thus can be used for map keys
type messageToSend struct {
	message     string
	phoneNumber int
}

func testMessageSend(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v \n", m.message, m.phoneNumber)
	fmt.Println("===============")
}

func structs() {
	testMessageSend(messageToSend{
		phoneNumber: 07003456741,
		message:     "Thanks for signing up",
	})

	// instantiating a variable using a struct
	// ! all the fields are initialized to their default values
	manchesterUnited := Team{
		Name:        "Manchester United",
		Manager:     "Eric Ten Hag",
		yearFounded: 1876,
		youngestPlayer: Player{
			isMarried: false,
			Name:      "Bobby Mainoo",
			Number:    23,
			Age:       16,
			Weight:    67.78,
		},
		topPlayer: Player{
			isMarried: false,
			Name:      "Bobby Charlton",
			Number:    23,
			Age:       16,
			Weight:    67.78,
		},
	}

	// ! Dot notation can also be used to override the default values of the fields
	/* manchesterUnited.Name = "Manchester United F.C."
	manchesterUnited.yearFounded = 1876
	manchesterUnited.Manager = "Eric Ten Hag"
	manchesterUnited.topPlayer.Name = "Marcus Rashford"
	manchesterUnited.topPlayer.isMarried = false
	manchesterUnited.topPlayer.Age = 26
	manchesterUnited.topPlayer.Weight = 76.56
	manchesterUnited.topPlayer.Number = 10
	manchesterUnited.youngestPlayer.Name = "Alex Garnacho"
	manchesterUnited.youngestPlayer.isMarried = false
	manchesterUnited.youngestPlayer.Age = 17
	manchesterUnited.youngestPlayer.Weight = 70.56
	manchesterUnited.youngestPlayer.Number = 11 */

	// for nested structs, the nested fields cannot be accessed at the top level
	// i.e manchesterUnited.Age (where Age is a field of the player struct)
	fmt.Println(manchesterUnited.youngestPlayer)

	// anonymous structs
	mercedes := struct {
		Name     string
		Category string
	}{
		Name:     "Mercedes Benz",
		Category: "Luxury",
	}
	fmt.Println(mercedes)

	// instantiating a variable using an embedded struct
	myChair := Chair{
		Recliner: false,
		Furniture: Furniture{
			Name:  "Chair",
			Make:  "Luxurious",
			Price: 78.00,
		},
	}
	// embedded fields can be accessed at the top level
	fmt.Println(myChair.Price)

	// struct methods
	r := rect{
		width:  45,
		height: 23,
	}
	fmt.Println(r.area())
}

// nested structs
type Team struct {
	Name           string
	yearFounded    int
	Manager        string
	youngestPlayer Player
	topPlayer      Player
}

// ? what are the rules around naming structs
type Player struct {
	Name      string
	Number    int
	Age       int
	Weight    float64
	isMarried bool
}

// nested anonymous structs
// ! avoid anonymous structs unless they really have to be used
type Book struct {
	Title       string
	Description string
	Author      struct {
		Name string
		Age  int
	}
}

// embedded structs
type Furniture struct {
	Name  string
	Make  string
	Price float64
}

type Chair struct {
	// all the fields are inherited e.g calling chair.Price directly
	// if chair is instantiated using the Chair struct
	// as opposed to chair.Furniture.Price for embedded structs
	Furniture
	Recliner bool
}

// struct methods
type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.height * r.width
}
