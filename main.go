package main

import(
	"github.com/CherubRocky/JSSP/jssp"
	"github.com/CherubRocky/JSSP/cli"
	"github.com/CherubRocky/JSSP/cs"
	
	"math/rand/v2"
	"flag"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprint(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	seed := flag.Uint64("seed", 4, "Usage: --seed <postive int>")
	filePath := flag.String("file", "instances/JSSP_50_20.txt", "Usage: --file <path to file>")
	nests := flag.Int("nests", 20, "Usage: --nests <postive int>")
	gens := flag.Int("gens", 500, "Usage: --gens <postive int>")
	catchProb := flag.Float64("catch-prob", 0.2, "Usage: --gens <rational number in range [0, 1)>")
	flag.Parse()
	
	sins, err := cli.MyReadFile(*filePath)

	if err != nil {
		return fmt.Errorf("Error reading instance file %s: %w", *filePath, err)
	}

	// create instance
	pcg := rand.NewPCG(*seed, 0)
	rng := rand.New(pcg)
	ins, err := jssp.NewInstance(sins, rng)
	
	if err != nil {
		return fmt.Errorf("Error parsing instance file %s: %w", *filePath, err)
	}
	
	var f cs.EggFactory = func() cs.Egg {
		return eggAdapter{ins.RandomSol()}
	}

	cuckooSearch := cs.NewCuckooSearch(*nests, *gens, *catchProb, f)

	bestSolution := cuckooSearch.Search()

	// Missing String Method in interface and implementation
	fmt.Println("Best solution: ")
	fmt.Printf("Cost: %f", bestSolution.Cost())

	return nil
}

// Adapter for solutions
type eggAdapter struct {
	*jssp.Solution
}

func (a eggAdapter) Neighbour() cs.Egg { return eggAdapter{a.Solution.Neighbour()} }
func (a eggAdapter) Fly() cs.Egg { return eggAdapter{a.Solution.Fly()} }
