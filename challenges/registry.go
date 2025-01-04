package challenges

import (
	day01 "aoc24/challenges/01_historian_hysteria"
	day02 "aoc24/challenges/02_red_nosed_reports"
	day03 "aoc24/challenges/03_mull_it_over"
)

var Challenges = make(map[string]func() error)

func init() {
	Challenges["01a"] = day01.A
	Challenges["01b"] = day01.B
	Challenges["02a"] = day02.A
	Challenges["02b"] = day02.B
	Challenges["03a"] = day03.A
	Challenges["03b"] = day03.B
}
