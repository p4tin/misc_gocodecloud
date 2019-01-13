package combat

import (
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

func init() {
	Out = os.Stdout
	In = os.Stdin
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	DisplayInfo("Welcome to combat 1.0")
	//Player
	p1 := new(Character)
	p1.Name = "Paul"
	p1.Speed = 1 + rand.Intn(100)
	p1.Health = 100
	p1.Alive = true
	p1.Weap = 1

	//Random NPC
	p2 := new(Character)
	*p2 = *ennemies[1+rand.Intn(2)]
	p2.Npc = true
	p2.Speed = 1 + rand.Intn(100)

	//p3 := new(Character)
	//*p3 = *ennemies[1+rand.Intn(2)]
	//p3.Npc = true
	//p3.Speed = 1 + rand.Intn(100)

	players := Players{*p1, *p2 /*, *p3 */}

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
