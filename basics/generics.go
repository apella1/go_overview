package main

import (
	"errors"
	"fmt"
)

func generics() {
	bs := bookStore{
		booksSold: []book{},
	}
	sellProducts[book](&bs, []book{
		{
			title:  "Deep Work",
			author: "Cal Newport",
			price:  56.99,
		},
		{
			title:  "How To be an A student",
			author: "Cal Newport",
			price:  56.99,
		},
		{
			title:  "So good to ignore",
			author: "Cal Newport",
			price:  56.99,
		},
		{
			title:  "Slow Productivity",
			author: "Cal Newport",
			price:  56.99,
		},
		{
			title:  "Win in High School",
			author: "Cal Newport",
			price:  56.99,
		},
	})
	fmt.Println(bs.booksSold)
	ts := toyStore{
		toysSold: []toy{},
	}
	sellProducts[toy](&ts, []toy{
		{
			name:  "Car",
			price: 40.00,
		},
		{
			name:  "Crayon",
			price: 40.00,
		},
	})
	fmt.Println(ts.toysSold)
}

// ! DRYing code

func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

type gnEmail struct {
	message     string
	senderEmail string
	recEmail    string
}

func getLast[T any](s []T) T {
	if len(s) == 0 {
		var zeroVal T
		return zeroVal
	}
	return s[len(s)-1]
}

// * constraints
type stringer interface {
	String() string
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		// ! the need for a more specific constraint rather than using `any`
		result += val.String()
	}
	return result
}

type lineItem interface {
	GetCost() float64
	GetName() string
}

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	itemCost := newItem.GetCost()
	if balance < itemCost {
		// zero value of slices is nil
		return nil, 0.0, errors.New("insufficient funds")
	}
	oldItems = append(oldItems, newItem)
	return oldItems, balance - itemCost, nil
}

// parametric constraints
// * represents a store that sells products
type store[P product] interface {
	Sell(P)
}

type product interface {
	Price() float64
	Name() string
}

type book struct {
	title  string
	author string
	price  float64
}

func (b book) Price() float64 {
	return b.price
}

func (b book) Name() string {
	return b.title
}

type toy struct {
	name  string
	price float64
}

func (t toy) Price() float64 {
	return t.price
}

func (t toy) Name() string {
	return t.name
}

// store that sells books
type bookStore struct {
	booksSold []book
}

type toyStore struct {
	toysSold []toy
}

// Sell adds books to the book store's inventory
func (bs *bookStore) Sell(b book) {
	bs.booksSold = append(bs.booksSold, b)
}

func (ts *toyStore) Sell(t toy) {
	ts.toysSold = append(ts.toysSold, t)
}

func sellProducts[P product](s store[P], products []P) {
	for _, p := range products {
		s.Sell(p)
	}
}
