package domain

type Order struct {
	PersonName string
	Telephone  string
	Positions  []Position

	Done bool
}

func (o Order) AddGood(good Good) {
	if o.Done {
		// TODO: Вернуть ошибку
		return
	}

	for _, position := range o.Positions {
		if position.Good.Name == good.Name {
			position.Amount += 1
			return
		}
	}

	o.Positions =
		append(o.Positions, Position{Good: good, Amount: 1})
}
