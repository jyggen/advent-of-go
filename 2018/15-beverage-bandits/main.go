package main

import (
	"fmt"
	"github.com/beefsack/go-astar"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	round, elves, goblins := playGame(rows, 3)
	remainingHp := 0

	for _, c := range elves {
		if !c.IsDead() {
			remainingHp += c.hp
		}
	}

	for _, c := range goblins {
		if !c.IsDead() {
			remainingHp += c.hp
		}
	}

	return strconv.Itoa(round * remainingHp), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	ap := 4
PartTwoLoop:
	for {
		round, elves, goblins := playGame(rows, ap)
		remainingHp := 0

		for _, c := range elves {
			if !c.IsDead() {
				remainingHp += c.hp
			} else {
				ap++
				continue PartTwoLoop
			}
		}

		for _, c := range goblins {
			if !c.IsDead() {
				remainingHp += c.hp
			}
		}

		return strconv.Itoa(round * remainingHp), nil
	}
}

func playGame(rows []string, elvesAp int) (int, []*Creature, []*Creature) {
	colLen := len(rows[0])
	battlefield := make([]*Tile, len(rows)*colLen)
	elves := make([]*Creature, 0)
	goblins := make([]*Creature, 0)
	input := 0

	for i, row := range rows {
		row := strings.Split(row, "")

		for j, column := range row {
			input++
			position := i*colLen + j

			switch column {
			case "#":
				battlefield[position] = &Tile{
					position: position,
					colLen:   colLen,
					occupied: true,
					tiles:    &battlefield,
				}
				break
			case ".":
				battlefield[position] = &Tile{
					position: position,
					colLen:   colLen,
					occupied: false,
					tiles:    &battlefield,
				}
				break
			case "E":
				battlefield[position] = &Tile{
					position: position,
					colLen:   colLen,
					occupied: true,
					tiles:    &battlefield,
				}
				elves = append(elves, &Creature{
					ap:      elvesAp,
					enemies: &goblins,
					hp:      200,
					id:      len(elves),
					kind:    "elf",
					tile:    battlefield[position],
				})
				break
			case "G":
				battlefield[position] = &Tile{
					position: position,
					colLen:   colLen,
					occupied: true,
					tiles:    &battlefield,
				}
				goblins = append(goblins, &Creature{
					ap:      3,
					enemies: &elves,
					hp:      200,
					id:      len(goblins),
					kind:    "goblin",
					tile:    battlefield[position],
				})
				break
			}
		}
	}

	creatures := make([]*Creature, len(elves)+len(goblins))

	copy(creatures, append(elves, goblins...))
	victory := false
	lastUnit := true
	round := 1

GameLoop:
	for {
		//Draw(battlefield, elves, goblins)

		if round > 47 {
			//break
		}

		//fmt.Printf("Starting round %d!\n", round)

		sort.Slice(creatures, func(i, j int) bool {
			return creatures[i].tile.position < creatures[j].tile.position
		})

		for _, c := range creatures {
			if c.IsDead() {
				continue
			}

			if victory {
				lastUnit = false
			}

			//fmt.Printf("\tIt's %s #%d's turn! This %s is at %s and has %d HP.\n", c.kind, c.id, c.kind, c.tile.Coordinates(), c.hp)

			if !c.HasEnemiesAlive() {
				//fmt.Printf("\t\tThis %s has no enemies alive - victory!\n", c.kind)
				victory = true
				break GameLoop
			}

			c.Move()
		}

		round++
	}

	if lastUnit {
		round--
	}

	//fmt.Println(round)
	//Draw(battlefield, elves, goblins)
	//fmt.Println()

	return round, elves, goblins
}

func calculateDistance(from int, to int, colLen int) int {
	if from > to {
		to, from = from, to
	}

	colStart := 0

	for to > colStart {
		colStart += colLen
	}

	colStart -= colLen
	y := 0

	for from < colStart {
		y++
		from += colLen
	}

	if from > to {
		to, from = from, to
	}

	return y + to - from
}

type Creature struct {
	ap      int
	enemies *[]*Creature
	hp      int
	id      int
	kind    string
	tile    *Tile
}

func (c *Creature) Attack(enemies []*Creature) {
	var closestEnemy *Creature

	for _, e := range enemies {
		//fmt.Printf("\t\t%s #%d is at %d with %d HP!\n", e.kind, e.id, e.tile.position, e.hp)
		if closestEnemy == nil || e.hp < closestEnemy.hp || (e.hp == closestEnemy.hp && e.tile.position < closestEnemy.tile.position) {
			closestEnemy = e
		}
	}

	if closestEnemy == nil {
		return
	}

	//fmt.Printf("\t\tAttacking %s #%d!\n", closestEnemy.kind, closestEnemy.id)

	closestEnemy.hp -= c.ap

	if closestEnemy.IsDead() {
		closestEnemy.tile.occupied = false
	}
}

func (c *Creature) HasEnemiesAlive() bool {
	for _, e := range *c.enemies {
		if !e.IsDead() {
			return true
		}
	}

	return false
}

func (c *Creature) IsDead() bool {
	return c.hp <= 0
}

func (c *Creature) IsNextTo(c2 *Creature) bool {
	for _, i := range []int{
		c.tile.position - c.tile.colLen,
		c.tile.position - 1,
		c.tile.position + 1,
		c.tile.position + c.tile.colLen,
	} {
		if i == c2.tile.position {
			return true
		}
	}

	return false
}

func (c *Creature) Move() error {
	options := make([]*Creature, 0)
	optionsLen := 0
	possibilities := make([][]*Tile, 0)

	for _, e := range *c.enemies {
		if e.IsDead() {
			continue
		}

		if c.IsNextTo(e) {
			//fmt.Printf("\t\tWe're next to %s #%d, could attack!\n", e.kind, e.id)
			options = append(options, e)
			optionsLen++
			continue
		}

		if optionsLen > 0 {
			continue
		}

		for _, et := range e.tile.PathNeighbors() {
			//fmt.Printf("\t\tCould attack %s #%d if we get to %s!\n", e.kind, e.id, et.(*Tile).Coordinates())

			for _, ct := range c.tile.PathNeighbors() {
				possibilities = append(possibilities, []*Tile{ct.(*Tile), et.(*Tile)})
			}
		}
	}

	if optionsLen > 0 {
		c.Attack(options)
		return nil
	}

	if len(possibilities) == 0 {
		//fmt.Print("\t\tNo move options. Skip turn :(\n")
		return nil
	}

	bestPath := make([]astar.Pather, 0)
	bestPathDistance := float64(-1)

	for _, p := range possibilities {
		path, distance, found := astar.Path(p[0], p[1])

		if !found {
			//fmt.Printf("\t\tNo path to %s from %s, skip!\n", p[1].Coordinates(), p[0].Coordinates())
			continue
		}

		/*coords := make([]string, len(path))

		for i, yolo := range path {
			coords[i] = yolo.(*Tile).Coordinates()
		}*/

		//fmt.Printf("\t\t%s is %f steps away!\n", strings.Join(coords, ", "), distance)

		if bestPathDistance == -1 || distance < bestPathDistance || (distance == bestPathDistance && path[0].(*Tile).position < bestPath[0].(*Tile).position) || (distance == bestPathDistance && path[0].(*Tile).position == bestPath[0].(*Tile).position && path[len(path)-1].(*Tile).position < bestPath[len(bestPath)-1].(*Tile).position) {
			bestPath = path
			bestPathDistance = distance
		}
	}

	if bestPathDistance == -1 {
		//fmt.Print("\t\tSeems like we can't go anywhere :(!\n")
		return nil
	}

	if bestPathDistance > -1 {
		newTile := bestPath[len(bestPath)-1].(*Tile)
		//fmt.Printf("\t\tMoved from %s to %s!\n", c.tile.Coordinates(), newTile.Coordinates())

		c.tile.occupied, newTile.occupied = false, true
		c.tile = newTile
		bestPathDistance--
	}

	if bestPathDistance != -1 {
		return nil
	}

	for _, e := range *c.enemies {
		if e.IsDead() {
			continue
		}

		if c.IsNextTo(e) {
			//fmt.Printf("\t\tWe can reach %s #%d!\n", e.kind, e.id)
			options = append(options, e)
		}
	}

	if len(options) == 0 {
		return nil
	}

	c.Attack(options)

	return nil
}

type Tile struct {
	position int
	colLen   int
	occupied bool
	tiles    *[]*Tile
}

func (t *Tile) Coordinates() string {
	x := t.position
	y := 0

	for x > t.colLen {
		y++
		x -= t.colLen
	}

	return fmt.Sprintf("%dx%d", y, x)
}

func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := make([]astar.Pather, 0)

	for _, i := range []int{t.position - t.colLen, t.position - 1, t.position + 1, t.position + t.colLen} {
		if i < 0 || i >= len(*t.tiles) || (*t.tiles)[i].occupied {
			continue
		}

		neighbors = append(neighbors, (*t.tiles)[i])
	}

	return neighbors
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	return float64(calculateDistance(t.position, to.(*Tile).position, t.colLen))
}

func Draw(tiles []*Tile, elves []*Creature, goblins []*Creature) {
	output := make([][]string, 0)

	for _, t := range tiles {
		x := t.position
		y := 0

		for x >= t.colLen {
			y++
			x -= t.colLen
		}

		if y == 0 {
			output = append(output, make([]string, t.colLen))
		}

		if t.occupied {
			output[y][x] = "#"
		} else {
			output[y][x] = "."
		}
	}

	for _, e := range elves {
		if e.IsDead() {
			continue
		}

		x := e.tile.position
		y := 0

		for x >= e.tile.colLen {
			y++
			x -= e.tile.colLen
		}

		output[y][x] = "E"
	}

	for _, g := range goblins {
		if g.IsDead() {
			continue
		}

		x := g.tile.position
		y := 0

		for x >= g.tile.colLen {
			y++
			x -= g.tile.colLen
		}

		output[y][x] = "G"
	}

	for _, cols := range output {
		for _, char := range cols {
			fmt.Print(char)
		}

		fmt.Println()
	}

	fmt.Println()
}
