package cs

type Egg interface {
	Revolve() Egg
	Neighbour() Egg
	Fly() Egg
	Cost() float64
	Feasible() bool
}

type EggFactory func() Egg

type CuckooSearch struct {
	nestNumber int
	eggRejection float64
	Nests []Egg
	BestSol Egg
	timesHatched float64
	maxGen int
	Hatch EggFactory
	rng *rand.Rand
	
}

func (c *CuckooSearch) Search() Egg {
	for iter := 0; iter < c.maxGen; iter++ {
		c.layEggs()
	}
	return nil
}

func (c *CuckooSearch) layEggs() {
	for i := 0; i < c.nestNumber; i++ {
		nEgg := c.flight(c.Nests[i])
		indComp := c.rng.IntN(c.nestNumber)
		for indComp == i {
			indComp = c.rng.IntN(c.nestNumber)
		}
		if nEgg.Cost() < c.Nests[indComp].Cost() {
			c.Nests[indComp] = nEgg
		}
	}
	c.rejectEggs()
}

// rejectEggs does the rejection of cuckoos step
// according to c.eggRejection for each egg
func (c *CuckooSearch) rejectEggs() {
	for i, _ := range c.Nests{
		prob := c.rng.Float64()
		if prob <= c.eggRejection {
			c.Nests[i] = c.Hatch()
		}
	}
}


// flight does a Lévy Flight on an egg
func (c *CuckooSearch) flight(e Egg) Egg {
	num := c.rng.Float64()

	switch {
	case num < 0.2:
		return e.Neighbour()
	case num < 0.4:
		return e.Neighbour().Neighbour()
	case num < 0.8:
		return e.Neighbour().Neighbour().Neighbour()
	default:
		return e.Fly()
	}
}
