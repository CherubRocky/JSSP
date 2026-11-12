package cs

import "math/rand/v2"

type Egg interface {
	Neighbour() Egg
	Fly() Egg
	Cost() float64
	//	Feasible() bool
}

type EggFactory func() Egg

type CuckooSearch struct {
	nestsNumber int
	maxGen int
	eggRejection float64
	Hatch EggFactory
	Nests []Egg
	BestSol Egg
	rng *rand.Rand	
}

func NewCuckooSearch(numberNests, numberGenerations int, rejectProb float64, f EggFactory) *CuckooSearch {
	pcg := rand.NewPCG(4, 0)
	generator := rand.New(pcg)
	cs := &CuckooSearch{
		nestsNumber: numberNests,
		maxGen: numberGenerations,
		eggRejection: rejectProb,
		Hatch: f,
		rng: generator,
	}
	cs.InitNests()
	return cs
}

func (c *CuckooSearch) InitNests() {
	c.Nests = make([]Egg, c.nestsNumber)
	for i, _ := range c.Nests {
		c.Nests[i] = c.Hatch()
		if i == 0 {
			c.BestSol = c.Nests[i]
		} else if c.Nests[i].Cost() < c.BestSol.Cost() {
			c.BestSol = c.Nests[i]
		}
	}
	
}

func (c *CuckooSearch) Search() Egg {
	for iter := 0; iter < c.maxGen; iter++ {
		c.layEggs()
	}
	return c.BestSol
}


func (c *CuckooSearch) layEggs() {
	for i := 0; i < c.nestsNumber; i++ {
		nEgg := c.flight(c.Nests[i])
		indComp := c.rng.IntN(c.nestsNumber)
		for indComp == i {
			indComp = c.rng.IntN(c.nestsNumber)
		}
		if nEgg.Cost() < c.Nests[indComp].Cost() {
			c.Nests[indComp] = nEgg
			if nEgg.Cost() < c.BestSol.Cost() {
				c.BestSol = nEgg
			}
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
		if c.Nests[i].Cost() < c.BestSol.Cost() {
			c.BestSol = c.Nests[i]
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
