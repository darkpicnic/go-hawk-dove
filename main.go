package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	Active = "active"
	Asleep = "asleep"
	Dead   = "dead"
	Hawk   = "hawk"
	Dove   = "dove"
)

const (
	StartingDoves                 = 150
	StartingHawks                 = 5
	StartingEnergy                = 100
	Rounds                        = 20
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
	currentRound := 0
	totalCulled := 0
	totalDoveBabies := 0
	totalHawkBabies := 0
	for currentRound < Rounds && len(agents) > 2 {
		awakenAgents()

		for {
			agent, nemesis, err := getRandomAgents()
			if err != nil {
				break
			}
			compete(agent, nemesis, createFood())
		}

		atrophyAgents()
		roundDeaths := cull()
		roundDoveBabies, roundHawkBabies := breed()
		totalDoveBabies += roundDoveBabies
		totalHawkBabies += roundHawkBabies
		totalCulled += roundDeaths
		currentRound += 1
		fmt.Printf("Round births: %d\n", roundDoveBabies+roundHawkBabies)
		fmt.Printf("Round %d deaths: %d\n", currentRound, roundDeaths)
	}

	fmt.Printf("Total agents: %d\n", len(agents))
	fmt.Printf("Total babies: %d\n", totalDoveBabies+totalHawkBabies)
	fmt.Printf("Total agents killed: %d\n", totalCulled)
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
		if agent.Status != Dead {
			agent.Energy -= EnergyLossPerRound
		}
	}
}

func createAgent(agentType string, startingEnergy int, status string) *Agent {
	agent := new(Agent)
	agent.Energy = startingEnergy
	agent.Status = status
	agent.Type = agentType
	return agent
}

func breed() (int, int) {
	doveCount := 0
	hawkCount := 0
	for _, agent := range agents {
		if agent.Status != Dead && agent.Energy >= EnergyRequiredForReproduction {
			babyA := createAgent(agent.Type, (agent.Energy / 2), Active)
			babyB := createAgent(agent.Type, (agent.Energy / 2), Active)
			agents = append(agents, babyA, babyB)
			agent.Energy /= 2

			if agent.Type == Hawk {
				hawkCount += 2
			}

			if agent.Type == Dove {
				doveCount += 2
			}
		}
	}

	return doveCount, hawkCount
}

func createFood() int {
	return MinFoodPerRound + rand.Intn(MaxFoodPerRound)
}

func awakenAgents() {
	for _, agent := range agents {
		if agent.Status != Dead {
			agent.Status = Active
		}
	}
}

func cull() int {
	totalCulled := 0
	for _, agent := range agents {
		if agent.Status != Dead && agent.Energy < EnergyRequiredForLiving {
			agent.Status = Dead
			totalCulled += 1
		}
	}

	return totalCulled
}

// getEnergyFromFood computes the energy value from food
// Right now, it's 1:1
func getEnergyFromFood(food int) int {
	return food
}

func compete(agent *Agent, nemesis *Agent, food int) {
	var winner Agent
	var loser Agent

	if rn := rand.Intn(1); rn == 0 {
		winner = *agent
		loser = *nemesis
	} else {
		winner = *nemesis
		loser = *agent
	}

	switch {

	case agent.Type == Hawk && nemesis.Type == Hawk:
		// Hawk / Hawk
		winner.Energy += getEnergyFromFood(food)
		loser.Energy -= EnergyLossFromFighting

	case agent.Type == Hawk && nemesis.Type == Dove:
		// Hawk / Dove
		agent.Energy += getEnergyFromFood(food)
		nemesis.Energy -= EnergyLostFromBluffing

	case agent.Type == Dove && nemesis.Type == Hawk:
		// Dove / Hawk
		nemesis.Energy += getEnergyFromFood(food)
		agent.Energy -= EnergyLostFromBluffing

	case agent.Type == Dove && nemesis.Type == Dove:
		// Dove / Dove
		winner.Energy += getEnergyFromFood(food)
		loser.Energy -= EnergyLostFromBluffing
	}

	agent.Status = Asleep
	nemesis.Status = Asleep
	return
}

func getAgentCountByStatus(status string) int {
	counter := 0
	for _, agent := range agents {
		if agent.Status == status {
			counter += 1
		}
	}

	return counter
}

func getRandomAgents() (agent *Agent, nemesis *Agent, err error) {
	if getAgentCountByStatus(Active) < 2 {
		err = errors.New("Agent or nemesis not found")
	}
	agentIndex := rand.Intn(len(agents) - 1)
	agent = agents[agentIndex]
	for nemesis == nil {
		if nemesisIndex := rand.Intn(len(agents) - 1); nemesisIndex != agentIndex {
			nemesis = agents[nemesisIndex]
		}
	}
	return
}
