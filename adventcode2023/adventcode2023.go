package adventcode2023

import (
	"errors"

	"alex952.com/advent2023/adventcode2023/day1"
	"alex952.com/advent2023/adventcode2023/day2"
	"alex952.com/advent2023/adventcode2023/shared"
)

type AdventOfCodeChallengRunner interface {
	RunChallenge(config shared.AdventOfCodeChallengRunnerConfig) (string, error)
}

var DayChallengesRunners = map[int]AdventOfCodeChallengRunner{
	1: &day1.Day1Runner{},
	2: &day2.Day2Runner{},
}

func MakeChallengeRunner(day int) (AdventOfCodeChallengRunner, error) {
	if runner, exists := DayChallengesRunners[day]; exists {
		return runner, nil
	} else {
		return nil, errors.New("Day not available!")
	}
}
