package _3_claw_contraption

import (
	"aoc24/vect"
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

var machinePattern = regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)
Button B: X\+(\d+), Y\+(\d+)
Prize: X=(\d+), Y=(\d+)`)

func A() error {

	matches := machinePattern.FindAllStringSubmatch(input, -1)

	var price int

	for _, match := range matches {
		a := vect.Vector{atoi(match[1]), atoi(match[2])}
		b := vect.Vector{atoi(match[3]), atoi(match[4])}
		p := vect.Vector{atoi(match[5]), atoi(match[6])}

		maxA := min(p.X/a.X, p.Y/a.Y)

		var combination vect.Vector

		for mA := 0; mA <= maxA; mA++ {
			tA := a.Mul(mA)
			diff := p.Sub(tA)
			if diff.X < 0 || diff.Y < 0 {
				continue
			}

			if mX, mY := diff.X/b.X, diff.Y/b.Y; mX == mY && diff.X%b.X == 0 && diff.Y%b.Y == 0 {
				combination = vect.Vector{mA, mX}
			}
		}

		price += combination.X*3 + combination.Y
	}

	fmt.Println(price)

	return nil
}

func B() error {

	matches := machinePattern.FindAllStringSubmatch(input, -1)

	var price int

	for _, match := range matches {
		const off = 10000000000000
		a := vect.Vector{atoi(match[1]), atoi(match[2])}
		b := vect.Vector{atoi(match[3]), atoi(match[4])}
		p := vect.Vector{atoi(match[5]) + off, atoi(match[6]) + off}

		// calculate gradients of lines
		mA := float64(a.Y) / float64(a.X)
		mB := float64(b.Y) / float64(b.X)
		mP := float64(p.Y) / float64(p.X)

		// lines must fall on opposite sides of the prize to be solvable
		dmA, dmB := mP-mA, mP-mB
		if math.Signbit(dmA) == math.Signbit(dmB) {
			continue
		}

		// calculate the intersection of the lines, with a crossing the origin, and b crossing the prize
		// 1)            y = mA*x
		// 2)            c = p.Y - mB*p.X
		// 3)            y = mB*x + c
		// 1, 3 -> 4) mA*x = mB*x + c
		// 4, 2 -> 5) mA*x = mB*x + p.Y - mB*p.X
		// 5*) (mA - mB)*x = p.Y - mB*p.X
		// 5**)          x = (p.Y - mB*p.X) / (mA - mB)
		num := float64(p.Y) - (mB * float64(p.X))
		denom := mA - mB
		x := int(math.Round(num / denom))

		// if the intersection is a whole number of a and b, the solution is valid
		if x%a.X == 0 && (p.X-x)%b.X == 0 {
			tA := x / a.X
			tB := (p.X - x) / b.X

			price += 3*tA + tB
		}
	}

	fmt.Println(price)

	return nil
}

func atoi(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
