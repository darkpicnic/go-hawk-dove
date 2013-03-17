package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Active = "active"
	Asleep = "asleep"
	Hawk   = "hawk"
	Dove   = "dove"
)

const (
	StartingDoves                 = 100
	StartingHawks                 = 5
	StartingEnergy                = 100
	Rounds                        = 100
	EnergyLossPerRound            = 2
	EnergyLossFromFighting        = 200
	EnergyRequiredForLiving       = 20
	EnergyLostFromBluffing        = 10
	EnergyRequiredForReproduction = 250

	MinFoodPerRound = 20
	MaxFoodPerRound = 70
)

var agents []*Agent

type Agent struct {
	Type   string
	Status string
	Energy int
}

func main() {
	currentRound := 1
	for currentRound < Rounds {
		awakenAgents()

		for {
			agent, nemesis := getRandomAgents()
			compete(agent, nemesis, createFood())
		}

		atrophyAgents()

		currentRound += 1
	}
	fmt.Printf("Total agents: %d", len(agents))
}

// Init creates the initial doves and hawks
func init() {

	rand.Seed(time.Now().UTC().UnixNano())

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

func atrophyAgents() {
	for _, agent := range agents {
		agent.Energy -= EnergyLossPerRound
	}
}

func createFood() int {
	return MinFoodPerRound + rand.Intn(MaxFoodPerRound)
}

func awakenAgents() {
	for _, agent := range agents {
		agent.Status = Active
	}
}

func compete(agent Agent, nemesis Agent, food int) {
	return
}

func getRandomAgents() (agent, nemesis Agent) {
	return
}
