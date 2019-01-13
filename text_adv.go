package main

import (
	"bufio"
	"cmd/go/testdata/testinternal2/x/y/z/internal/w"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

var Out *os.File
var In *os.File

type Weapon struct {
	minAtt int
	maxAtt int
	Name   string
}

func (w *Weapon) Fire() int {
	return w.minAtt + rand.Intn(w.maxAtt)
}

type Character struct {
	Name    string
	Health  int
	Evasion int
	Alive   bool
	Speed   int
	Weap    int
	Npc     bool
	Items   []int

	Welcome         string
	CurrentLocation string
}

var ennemies = map[int]*Character{
	1: {Name: "Klingon", Health: 50, Alive: true, Weap: 2},
	2: {Name: "Romulan", Health: 55, Alive: true, Weap: 3},
}

func (p *Character) Equip(w int) {
	p.Weap = w
}

func (p *Character) Attack() int {
	return Weaps[p.Weap].Fire()
}

var Weaps = map[int]*Weapon{
	1: {Name: "Phaser", minAtt: 5, maxAtt: 15},
	2: {Name: "Klingon Disruptor", minAtt: 1, maxAtt: 15},
	3: {Name: "Romulan Disruptor", minAtt: 3, maxAtt: 12},
}

type Players []Character

func (slice Players) Len() int {
	return len(slice)
}

func (slice Players) Less(i, j int) bool {
	return slice[i].Speed > slice[j].Speed //Sort descending
	//return slice[i].Speed < slice[j].Speed;		//Sort ascending
}

func (slice Players) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type Game struct {
	Welcome         string
	Health          int
	CurrentLocation string
}

func (p *Character) Play() {
	DisplayInfo(p.Welcome)
	for {
		DisplayInfo(locationMap[p.CurrentLocation].Description)
		p.ProcessEvents(locationMap[p.CurrentLocation].Events)
		if p.Health <= 0 {
			DisplayInfo("You are dead, game over!!!")
			return
		}
		DisplayInfo("Health:", p.Health)
		DisplayInfo("You can go to these places:")
		for index, loc := range locationMap[p.CurrentLocation].Transitions {
			DisplayInfof("%d - %s\n", index+1, loc)
		}
		i := 0
		for i < 1 || i > len(locationMap[p.CurrentLocation].Transitions) {
			DisplayInfof("%s%d%s\n", "Where do you want to go (0 - to quit), [1...", len(locationMap[p.CurrentLocation].Transitions), "]: ")
			GetUserInput(&i)
		}
		newLoc := i - 1
		p.CurrentLocation = locationMap[p.CurrentLocation].Transitions[newLoc]

	}
}

func (p *Character) ProcessEvents(events []string) {
	for _, evtName := range events {
		p.Health += evts[evtName].ProcessEvent()
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
		if e.Type == "Combat" {
			//Generate opponent
			opp := rand.Intn(len(ennemies))
			// Run combat
			DisplayInfo("Combat Event")
		} else {
			DisplayInfo("\t" + e.Description)
			if e.Evt != "" {
				e.Health = e.Health + evts[e.Evt].ProcessEvent()
			}
		}
		return e.Health
	}
	return 0
}

type Location struct {
	Description string
	Transitions []string
	Events      []string
	Items       []int
}

var evts = map[string]*Event{
	"alienAttack":     {Type: "Combat", Chance: 80, Description: "An alien beams in front of you and shoots you with a ray gun.", Health: -50, Evt: "doctorTreatment"},
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

func init() {
	Out = os.Stdout
	In = os.Stdin
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	p := &Character{Health: 100, Welcome: "Welcome to the Starship Enterprise\n\n", CurrentLocation: "Bridge"}
	p.Play()
}

func runBattle(players Players) {
	sort.Sort(players)

	DisplayInfo(players[0])
	DisplayInfo(players[1])
	DisplayInfo(players)
	round := 1
	numAlive := players.Len()
	playerAction := 0
	for {
		for x := 0; x < players.Len(); x++ {
			players[x].Evasion = 0
		}
		DisplayInfo("Combat round", round, "begins...")
		for x := 0; x < players.Len(); x++ {
			if players[x].Alive != true {
				continue
			}
			playerAction = 0
			if !players[x].Npc {
				DisplayInfo("DO you want to")
				DisplayInfo("\t1 - Run")
				DisplayInfo("\t2 - Evade")
				DisplayInfo("\t3 - Attack")
				GetUserInput(&playerAction)
			}
			if playerAction == 2 {
				players[x].Evasion = rand.Intn(15)
				DisplayInfo("Evasion set to:", players[x].Evasion)
			}
			tgt := selectTarget(players, x)
			if tgt != -1 {
				DisplayInfo("player: ", x, "target: ", tgt)
				attp1 := players[x].Attack() - players[tgt].Evasion
				if attp1 < 0 {
					attp1 = 0
				}
				players[tgt].Health = players[tgt].Health - attp1
				if players[tgt].Health <= 0 {
					players[tgt].Alive = false
					numAlive--
				}
				DisplayInfo(players[x].Name+" attacks and does", attp1, "points of damage with his", Weaps[players[x].Weap].Name, "to the ennemy.")
			}
		}
		if endBattle(players) || playerAction == 1 {
			break
		} else {
			DisplayInfo(players)
			round++
		}
	}
	DisplayInfo(players)
	DisplayInfo("Combat is over...")
	for x := 0; x < players.Len(); x++ {
		if players[x].Alive == true {
			DisplayInfo(players[x].Name + " is still alive!!!")
		}
	}
	DisplayInfo(players)
}

func selectTarget(players []Character, x int) int {
	y := x
	for {
		y = y + 1
		if y >= len(players) {
			y = 0
		}
		if (players[y].Npc != players[x].Npc) && players[y].Alive {
			return y
		}
		if y == x {
			return -1
		}
	}
	return -1
}

func endBattle(players []Character) bool {
	count := make([]int, 2)
	count[0] = 0
	count[1] = 0
	for _, pla := range players {
		if pla.Alive {
			if pla.Npc == false {
				count[0]++
			} else {
				count[1]++
			}
		}
	}
	if count[0] == 0 || count[1] == 0 {
		return true
	} else {
		return false
	}
}

func DisplayInfof(format string, args ...interface{}) {
	fmt.Fprintf(Out, format, args...)
}

func DisplayInfo(args ...interface{}) {
	fmt.Fprintln(Out, args...)
}

func GetUserInput(i *int) {
	fmt.Fscan(In, i)
}

func GetUserStrInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n >>> ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	return text
}
