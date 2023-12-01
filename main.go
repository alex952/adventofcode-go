package main

import (
	"flag"
	"fmt"

	"alex952.com/advent2023/adventcode2023"
	"alex952.com/advent2023/adventcode2023/shared"
	"github.com/sirupsen/logrus"
)

func main() {
	available_keys := []byte{}

	for k := range adventcode2023.DayChallengesRunners {
		available_keys = fmt.Appendf(available_keys, "%d,", k)
	}

	day := flag.Int("day", 1, fmt.Sprintf("Day to run (available days [%s])", string(available_keys[0:len(available_keys)-1])))
	first := flag.Bool("first", false, "Whether to run first challenge or second")
	debug := flag.Bool("debug", false, "Enable debug logging")
	input := flag.String("input", "input.txt", "Challenge input")

	flag.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	config := shared.AdventOfCodeChallengRunnerConfig{Filename: *input, First: *first}

	runner, err := adventcode2023.MakeChallengeRunner(*day)
	if err != nil {
		fmt.Println(err)
		flag.PrintDefaults()
		return
	}
	result, err := runner.RunChallenge(config)

	if err != nil {
		fmt.Printf("Error!: %s\n", err)
		return
	}

	fmt.Printf("Result!: %s\n", result)
}
