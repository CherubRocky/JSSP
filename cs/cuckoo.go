package cs

type Egg interface {
	Neighbour() Egg
	Fly() Egg
	Cost() float64
	Feasible() bool
}

type CuckooSearch struct {
	Nests []Egg
	timesHatched float64
	
}

func (c *CuckooSearch) Search() Egg {
	return nil
}

func (c *CuckooSearch) Hatch() Egg {
	return nil
}
