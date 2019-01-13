package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Welcome         string
	Health          int
	CurrentLocation string
}

func (g *Game) Play() {
	fmt.Println(g.Welcome)
	for {
		fmt.Println(locationMap[g.CurrentLocation].Description)
		g.ProcessEvents(locationMap[g.CurrentLocation].Events)
		if g.Health <= 0 {
			fmt.Println("You are dead, game over!!!")
			return
		}
		fmt.Printf("Health: %d\n", g.Health)
		fmt.Println("You can go to these places:")
		for index, loc := range locationMap[g.CurrentLocation].Transitions {
			fmt.Printf("\t%d - %s\n", index+1, loc)
		}
		i := 0
		for i < 1 || i > len(locationMap[g.CurrentLocation].Transitions) {
			fmt.Printf("%s%d%s\n", "Where do you want to go (0 - to quit), [1...", len(locationMap[g.CurrentLocation].Transitions), "]: ")
			fmt.Scan(&i)
		}
		newLoc := i - 1
		g.CurrentLocation = locationMap[g.CurrentLocation].Transitions[newLoc]

	}
}

func (g *Game) ProcessEvents(events []string) {
	for _, evtName := range events {
		g.Health += evts[evtName].ProcessEvent()
	}
}

type Event struct {
	Type        string
	Chance      int
	Description string
	Health      int
	Evt         string
}

func (e *Event) ProcessEvent() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	if e.Chance >= r1.Intn(100) {
		hp := e.Health
		if e.Type == "Combat" {
			fmt.Println("Combat Event")
		}
		fmt.Printf("\t%s\n", e.Description)
		if e.Evt != "" {
			hp = hp + evts[e.Evt].ProcessEvent()
		}
		return hp
	}
	return 0
}

type Location struct {
	Description string
	Transitions []string
	Events      []string
}

var evts = map[string]*Event{
	"alienAttack":     {Type: "Combat", Chance: 20, Description: "An alien beams in front of you and shoots you with a ray gun.", Health: -50, Evt: "doctorTreatment"},
	"doctorTreatment": {Type: "Story", Chance: 10, Description: "The doctor rushes in and inject you with a health boost.", Health: +30, Evt: ""},
	"android":         {Type: "Story", Chance: 50, Description: "Data is in the turbo lift and says hi to you", Health: 0, Evt: ""},
	"relaxing":        {Type: "Story", Chance: 100, Description: "In the lounge you are so relaxed that your health improves.", Health: +10, Evt: ""},
}

var locationMap = map[string]*Location{
	"Bridge":      {"You are on the bridge of a spaceship sitting in the Captain's chair.", []string{"Ready Room", "Turbo Lift"}, []string{"alienAttack"}},
	"Ready Room":  {"The Captain's ready room.", []string{"Bridge"}, []string{}},
	"Turbo Lift":  {"A Turbo Lift that takes you anywhere in the ship.", []string{"Bridge", "Lounge", "Engineering"}, []string{"android"}},
	"Engineering": {"You are in engineering where you see the star drive", []string{"Turbo Lift"}, []string{"alienAttack"}},
	"Lounge":      {"You are in the lounge, you feel very relaxed", []string{"Turbo Lift"}, []string{"relaxing"}},
}

func main() {
	g := &Game{Health: 100, Welcome: "Welcome to the Starship Enterprise\n\n", CurrentLocation: "Bridge"}
	g.Play()
}
