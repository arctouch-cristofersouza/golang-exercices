package main

type card struct {
}

type cards struct {
	cards []card
}

type deck interface {
	new()
	sort() []card
	shuffle() []card
}
