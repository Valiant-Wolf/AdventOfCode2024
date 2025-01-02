package challenges

import (
	_1_historian_hysteria "aoc24/challenges/01_historian_hysteria"
	_2_red_nosed_reports "aoc24/challenges/02_red_nosed_reports"
)

var Challenges = make(map[string]func() error)

func init() {
	Challenges["01a"] = _1_historian_hysteria.A
	Challenges["01b"] = _1_historian_hysteria.B
	Challenges["02a"] = _2_red_nosed_reports.A
	Challenges["02b"] = _2_red_nosed_reports.B
}
