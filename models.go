package main

type Perk struct {
	Description string
	Like        int
	Heart       int
}

type Company struct {
	Id     int
	Name   string
	Detail string
	Perks  []Perk
}
