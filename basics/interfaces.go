package main

import "math"

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}
type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .05
	}
	return float64(len(s.body)) * .01
}

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()
	}

	s, ok := e.(sms)
	if ok {
		return s.toPhoneNumber, s.cost()
	}

	return "", 0.0
}

func getExpensesReportSwitchCase(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}

func interfaces() {
}

// interfaces are a collection of method signatures
type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

func (rect rectangle) area() float64 {
	return rect.height * rect.width
}

func (rect rectangle) perimeter() float64 {
	return 2 * (rect.height + rect.width)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}

// ! a type can implement multiple interfaces
// the empty interface interface{} is implemented by all types since it has no requirements
// * interfaces are implemented implicitly

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name          string
	monthlySalary int
}

func (f fullTime) getName() string {
	return f.name
}

func (f fullTime) getSalary() int {
	return f.monthlySalary * 12
}

// naming interface method arguments
// provides intent for the types implementing the interface

type copier interface {
	copy(sourceFile string, destinationFile string) (bytesCopied int)
}
