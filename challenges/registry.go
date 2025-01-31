package challenges

import (
	day01 "aoc24/challenges/01_historian_hysteria"
	day02 "aoc24/challenges/02_red_nosed_reports"
	day03 "aoc24/challenges/03_mull_it_over"
	day04 "aoc24/challenges/04_ceres_search"
	day05 "aoc24/challenges/05_print_queue"
	day06 "aoc24/challenges/06_guard_gallivant"
	day07 "aoc24/challenges/07_bridge_repair"
	day08 "aoc24/challenges/08_resonant_collinearity"
	day09 "aoc24/challenges/09_disk_fragmenter"
	day10 "aoc24/challenges/10_hoof_it"
	day11 "aoc24/challenges/11_plutonian_pebbles"
	day12 "aoc24/challenges/12_garden_groups"
	day13 "aoc24/challenges/13_claw_contraption"
)

var Challenges = make(map[string]func() error)

func init() {
	Challenges["01a"] = day01.A
	Challenges["01b"] = day01.B
	Challenges["02a"] = day02.A
	Challenges["02b"] = day02.B
	Challenges["03a"] = day03.A
	Challenges["03b"] = day03.B
	Challenges["04a"] = day04.A
	Challenges["04b"] = day04.B
	Challenges["05a"] = day05.A
	Challenges["05b"] = day05.B
	Challenges["06a"] = day06.A
	Challenges["06b"] = day06.B
	Challenges["07a"] = day07.A
	Challenges["07b"] = day07.B
	Challenges["08a"] = day08.A
	Challenges["08b"] = day08.B
	Challenges["09a"] = day09.A
	Challenges["09b"] = day09.B
	Challenges["10a"] = day10.A
	Challenges["10b"] = day10.B
	Challenges["11a"] = day11.A
	Challenges["11b"] = day11.B
	Challenges["12a"] = day12.A
	Challenges["12b"] = day12.B
	Challenges["13a"] = day13.A
	Challenges["13b"] = day13.B
}
