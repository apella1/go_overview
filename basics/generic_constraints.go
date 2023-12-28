package main

import "fmt"

type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type gnUser struct {
	UserEmail string
}

func (u gnUser) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin gnUser
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.UserEmail
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o.Admin,
		Amount:   amount,
	}
}

type userBiller struct {
	Plan string
}

func (ub userBiller) Name() string {
	return fmt.Sprintf("%s org biller", ub.Plan)
}

func (ub userBiller) Charge(u gnUser) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}
