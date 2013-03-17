package main

import "fmt"

const (
	Active = "active"
	Asleep = "asleep"
	Hawk   = "hawk"
	Dove   = "dove"
)

const (
	StartingDoves  = 100
	StartingHawks  = 5
	StartingEnergy = 100
	Rounds         = 100
)

var agents []*Agent

type Agent struct {
	Type   string
	Status string
	Energy int
}

func main() {
	fmt.Printf("Total agents: %d", len(agents))
}

// Init creates the initial doves and hawks
func init() {
	doveCount := 1
	for doveCount < StartingDoves {
		d := new(Agent)
		d.Energy = StartingEnergy
		d.Status = Active
		d.Type = Dove
		agents = append(agents, d)
		doveCount += 1
	}

	hawkCount := 1
	for hawkCount < StartingHawks {
		d := new(Agent)
		d.Energy = StartingEnergy
		d.Status = Active
		d.Type = Hawk
		agents = append(agents, d)
		hawkCount += 1
	}
}
