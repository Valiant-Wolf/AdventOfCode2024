package _6_guard_gallivant

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func A() error {
	world, start := parseMap(input)

	var guard guard
	guard.walk(world, start)

	fmt.Println(guard.countSeen())

	return nil
}

func B() error {
	world, start := parseMap(input)

	var guard checkingGuard
	guard.walk(world, start)

	fmt.Println(guard.countCandidates())

	return nil
}

func parseMap(stringMap string) (world *worldMap, start position) {
	obstacles := make(map[position]bool)
	rows := strings.Split(strings.TrimSpace(stringMap), "\n")
	for y, row := range rows {
		for x, char := range []rune(row) {
			switch char {
			case '^':
				start = position{x, y}
			case '#':
				obstacles[position{x, y}] = true
			}
		}
	}
	world = &worldMap{len(rows[0]), len(rows), obstacles}
	return
}

//region worldMap

// worldMap is a representation of the North Pole prototype suit manufacturing lab.
type worldMap struct {
	width  int
	height int
	// obstacles is the set of positions containing an obstruction.
	obstacles map[position]bool
}

// inBounds returns true if a given position is within the bounds of the worldMap.
func (w worldMap) inBounds(p position) bool {
	return p.X >= 0 && p.X < w.width && p.Y >= 0 && p.Y < w.height
}

// isObstructed returns true if a given position contains an obstruction.
func (w worldMap) isObstructed(p position) bool {
	_, result := w.obstacles[p]
	return result
}

// copy returns a deep copy of a worldMap.
func (w worldMap) copy() worldMap {
	copyObs := make(map[position]bool)
	for k, v := range w.obstacles {
		copyObs[k] = v
	}
	return worldMap{w.width, w.height, copyObs}
}

//endregion

//region guard

// guard is a representation of a patrolling guard in the lab.
type guard struct {
	current position
	facing  direction
	// seen is the set of positions this guard has visited.
	seen map[position]bool
}

// walk moves the guard from a starting position on the map until they exit the map.
// The guard begins facing upwards.
func (g *guard) walk(world *worldMap, start position) {
	g.current = start
	g.facing = up
	g.seen = make(map[position]bool)

	// while the guard is within bounds, step the guard and update the current position
	for exit := false; !exit; g.current, exit = g.step(world) {
		// mark the current position as seen
		g.seen[g.current] = true
	}
}

// step moves the guard once, changing either its position or orientation, returning the new position, and whether the
// guard has stepped out of bounds.
func (g *guard) step(world *worldMap) (pos position, exit bool) {
	newPos := g.facing.move(g.current)

	// if the new position is out of bounds, return with exit
	if !world.inBounds(newPos) {
		return newPos, true
	}

	// if the new position is obstructed, hold position and rotate
	if world.isObstructed(newPos) {
		g.facing = g.facing.rotateCW()
		return g.current, false
	}

	// otherwise return the new position
	return newPos, false
}

// countSeen returns the number of distinct positions the guard has occupied.
func (g *guard) countSeen() int {
	return len(g.seen)
}

//endregion

//region loopingGuard

// loopingGuard is a guard that can detect whether it is caught in a loop.
type loopingGuard struct {
	guard
	// bumps is a collection of positions and initial orientations where this guard has turned.
	bumps map[position]map[direction]bool
}

// walk moves the guard on a map from a starting position and orientation until they loop or exit the map.
// walk returns true if the guard detects it has entered a loop.
func (g *loopingGuard) walk(world *worldMap, start position, dir direction) bool {
	g.current = start
	g.facing = dir
	g.seen = make(map[position]bool)
	g.bumps = make(map[position]map[direction]bool)

	looped := false

	// while the guard is within bounds and not looping, step the guard and update the current position
	for exit := false; !(exit || looped); g.current, exit, looped = g.step(world) {
		// mark the current position as seen
		g.seen[g.current] = true
	}

	return looped
}

// step moves the guard once, changing either its position or orientation, returning the new position, whether the guard
// has stepped out of bounds, and whether the guard has detected it's looping
func (g *loopingGuard) step(world *worldMap) (pos position, exit bool, looped bool) {
	pos = g.current
	newPos := g.facing.move(g.current)

	// if the new position is out of bounds, return with exit
	if !world.inBounds(newPos) {
		pos, exit = newPos, true
		return
	}

	// if the new position is obstructed...
	if world.isObstructed(newPos) {
		// check if we have turned at this location before
		entry, present := g.bumps[g.current]
		if !present {
			g.bumps[g.current] = make(map[direction]bool)
			entry = g.bumps[g.current]
		}

		// if we have turned at this location from this direction before, return with looped
		if entry[g.facing] {
			looped = true
			return
		}

		// record that we turned at this location from this direction
		g.bumps[g.current][g.facing] = true

		// hold position and rotate
		g.facing = g.facing.rotateCW()
		return
	}

	// otherwise move forward
	pos = newPos
	return
}

//endregion

//region checkingGuard

// checkingGuard is a guard that checks for obstruction placements that will cause a loop while it walks.
type checkingGuard struct {
	guard
	// candidates is the set of positions which will result in a loop if an obstruction is placed there.
	candidates map[position]bool
}

// walk moves the guard from a starting position on the map until they exit the map.
// The guard begins facing upwards.
// Before each movement, the guard will check if placing an obstacle directly in front of them will result in them
// entering a loop.
func (g *checkingGuard) walk(world *worldMap, start position) {
	g.current = start
	g.facing = up
	g.seen = make(map[position]bool)

	g.candidates = make(map[position]bool)

	// while the guard is within bounds, step the guard and update the current position
	for exit := false; !exit; g.current, exit = g.step(world) {
		// mark the current position as seen
		g.seen[g.current] = true

		// get the position in front of the guard
		front := g.facing.move(g.current)
		// check if we have already considered this position
		_, skip := g.candidates[front]

		// a position is a candidate for a new obstacle iff:
		// - it isn't already a candidate (saves duplicated work)
		// - it hasn't been crossed already (placing an obstacle would prevent the guard reaching here)
		// - it is in bounds
		// - it doesn't already have an obstacle there
		if !skip && !g.seen[front] && world.inBounds(front) && !world.isObstructed(front) {
			// create a copy of the world with an obstacle placed in front of the guard
			testWorld := world.copy()
			testWorld.obstacles[front] = true

			// create a copy of the guard and check if it enters a loop
			var looper loopingGuard
			loops := looper.walk(&testWorld, g.current, g.facing)

			// if the copy guard loops, the position we're checking is a candidate
			if loops {
				g.candidates[front] = true
			}
		}
	}
}

// returns the number of distinct candidates for looping obstruction placements.
func (g *checkingGuard) countCandidates() int {
	return len(g.candidates)
}

//endregion

//region position

type position struct {
	X int
	Y int
}

func (p position) String() string {
	return fmt.Sprintf("%d, %d", p.X, p.Y)
}

//endregion

//region direction

type direction struct {
	position
}

var (
	up    = direction{position{0, -1}}
	down  = direction{position{0, 1}}
	left  = direction{position{-1, 0}}
	right = direction{position{1, 0}}
)

func (d direction) move(p position) position {
	return position{p.X + d.X, p.Y + d.Y}
}

func (d direction) rotateCW() direction {
	switch d {
	case up:
		d = right
	case right:
		d = down
	case down:
		d = left
	case left:
		d = up
	}
	return d
}

//endregion
