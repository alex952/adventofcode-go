package day6

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"alex952.com/advent2023/adventcode2023/shared"
	"github.com/sirupsen/logrus"
)

type RaceData struct {
	Time     int
	Distance int
}

func (r RaceData) NumberOfWaysToWin() int {
	n := 0
	for i := 0; i <= r.Time; i += 1 {
		if target_d := i * (r.Time - i); target_d > r.Distance {
			n += 1
		} else if n > 0 && target_d <= r.Distance {
			break
		}
	}

	return n
}

type Races []RaceData

func ReadRacesPart1(scanner *bufio.Scanner) Races {
	scanner.Scan()
	times_st := scanner.Text()
	scanner.Scan()
	distances_st := scanner.Text()

	numbers_regex := regexp.MustCompile(`[0-9]+`)
	times := numbers_regex.FindAllString(times_st, -1)
	distances := numbers_regex.FindAllString(distances_st, -1)

	ret := make(Races, 0)

	for i := 0; i < len(times); i += 1 {
		d_int, _ := strconv.Atoi(distances[i])
		t_int, _ := strconv.Atoi(times[i])
		race := RaceData{Time: t_int, Distance: d_int}
		logrus.WithFields(logrus.Fields{
			"Time":     race.Time,
			"Distance": race.Distance,
		}).Debug("Race Data Loaded")
		ret = append(ret, race)
	}

	return ret
}

func ReadRacesPart2(scanner *bufio.Scanner) Races {
	scanner.Scan()
	times_st := scanner.Text()
	times_st = strings.Replace(strings.Split(times_st, ":")[1], " ", "", -1)
	scanner.Scan()
	distances_st := scanner.Text()
	distances_st = strings.Replace(strings.Split(distances_st, ":")[1], " ", "", -1)

	numbers_regex := regexp.MustCompile(`[0-9]+`)
	times := numbers_regex.FindAllString(times_st, -1)
	distances := numbers_regex.FindAllString(distances_st, -1)

	ret := make(Races, 0)

	for i := 0; i < len(times); i += 1 {
		d_int, _ := strconv.Atoi(distances[i])
		t_int, _ := strconv.Atoi(times[i])
		race := RaceData{Time: t_int, Distance: d_int}
		logrus.WithFields(logrus.Fields{
			"Time":     race.Time,
			"Distance": race.Distance,
		}).Debug("Race Data Loaded")
		ret = append(ret, race)
	}

	return ret
}

func RunCalc(races_data Races) string {
	ret := 1
	for _, r := range races_data {
		if wins := r.NumberOfWaysToWin(); wins > 0 {
			ret *= wins
		}
	}

	return fmt.Sprintf("%d", ret)
}

type Day6Runner struct{}

func (runner *Day6Runner) RunChallenge(config shared.AdventOfCodeChallengRunnerConfig) (string, error) {
	f, err := os.Open(config.Filename)
	if err != nil {
		fmt.Println("Cannot open file")
		return "", errors.New("Couldn't run the challenge. Can't open the file")
	}

	scanner := bufio.NewScanner(f)

	if config.First {
		races := ReadRacesPart1(scanner)
		return RunCalc(races), nil
	} else {
		races := ReadRacesPart2(scanner)
		return RunCalc(races), nil
	}
}
