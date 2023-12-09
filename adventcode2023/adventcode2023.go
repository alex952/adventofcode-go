package adventcode2023

import (
	"errors"

	"alex952.com/advent2023/adventcode2023/day1"
	"alex952.com/advent2023/adventcode2023/day2"
	"alex952.com/advent2023/adventcode2023/day3"
	"alex952.com/advent2023/adventcode2023/day4"
	"alex952.com/advent2023/adventcode2023/day5"
	"alex952.com/advent2023/adventcode2023/day6"
	"alex952.com/advent2023/adventcode2023/day7"
	"alex952.com/advent2023/adventcode2023/day8"
	"alex952.com/advent2023/adventcode2023/day9"
	"alex952.com/advent2023/adventcode2023/shared"
)

type AdventOfCodeChallengRunner interface {
	RunChallenge(config shared.AdventOfCodeChallengRunnerConfig) (string, error)
}

var DayChallengesRunners = map[int]AdventOfCodeChallengRunner{
	1: &day1.Day1Runner{},
	2: &day2.Day2Runner{},
	3: &day3.Day3Runner{},
	4: &day4.Day4Runner{},
	5: &day5.Day5Runner{},
	6: &day6.Day6Runner{},
	7: &day7.Day7Runner{},
	8: &day8.Day8Runner{},
	9: &day9.Day9Runner{},
}

func MakeChallengeRunner(day int) (AdventOfCodeChallengRunner, error) {
	if runner, exists := DayChallengesRunners[day]; exists {
		return runner, nil
	} else {
		return nil, errors.New("Day not available!")
	}
}
