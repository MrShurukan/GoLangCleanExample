package domain

type Good struct {
	Name        string
	Price       float64
	Description string
}

type Position struct {
	Good   Good
	Amount int
}
