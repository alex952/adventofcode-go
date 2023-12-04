package day4

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"alex952.com/advent2023/adventcode2023/shared"
	"github.com/sirupsen/logrus"
)

type ScratchCard struct {
	Id             int
	WinningNumbers []int
	Numbers        []int
	OurNumbers     []int
}

func ReadCard(line string) *ScratchCard {
	scratch := ScratchCard{}
	rex := regexp.MustCompile(`Card (?P<id>[0-9]+):(?P<win>(?:[0-9]+)+)| (?P<ours>(?:[0-9]+)+)`)
	index_bar := strings.Index(line, "|")
	submatches := rex.FindAllStringIndex(line, -1)

	scratch.Id, _ = strconv.Atoi(strings.Trim(line[submatches[0][0]:submatches[0][1]], " "))
	scratch.Numbers = make([]int, 0)
	scratch.OurNumbers = make([]int, 0)
	scratch.WinningNumbers = make([]int, 0)

	for idx := 1; idx < len(submatches); idx = idx + 1 {
		if m := submatches[idx]; m[0] < index_bar {
			n, _ := strconv.Atoi(strings.Trim(line[m[0]:m[1]], " "))
			scratch.Numbers = append(scratch.Numbers, n)
		} else {
			n, _ := strconv.Atoi(strings.Trim(line[m[0]:m[1]], " "))
			scratch.OurNumbers = append(scratch.OurNumbers, n)

			if slices.Contains(scratch.Numbers, n) {
				scratch.WinningNumbers = append(scratch.WinningNumbers, n)
			}
		}

	}

	return &scratch
}

func (s *ScratchCard) CalculateWinningScore() float64 {
	winning := 0.0
	if len(s.WinningNumbers) > 0 {
		winning = math.Pow(2, float64(len(s.WinningNumbers)-1))
	}
	logrus.WithFields(logrus.Fields{
		"Total Winning Numbers": len(s.WinningNumbers),
		"Total Winning":         winning,
	}).Debug("Win for Card %d calculated", s.Id)

	return winning
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func RunPart1(data []*ScratchCard) string {
	sum := 0.0
	for _, v := range data {
		sum += v.CalculateWinningScore()
	}

	return fmt.Sprintf("%d", int(sum))
}
func RunPart2(data []*ScratchCard) string {
	new_slice := make([]*ScratchCard, len(data))
	copy(new_slice, data)

	for idx := 0; idx < len(new_slice); idx = idx + 1 {
		if len(new_slice[idx].WinningNumbers) > 0 {
			new_slice = append(new_slice, data[new_slice[idx].Id:min(len(data), new_slice[idx].Id+len(new_slice[idx].WinningNumbers))]...)
		}
	}

	return fmt.Sprintf("%d", len(new_slice))
}

type Day4Runner struct{}

func (runner *Day4Runner) RunChallenge(config shared.AdventOfCodeChallengRunnerConfig) (string, error) {
	f, err := os.Open(config.Filename)
	if err != nil {
		fmt.Println("Cannot open file")
		return "", errors.New("Couldn't run the challenge. Can't open the file")
	}

	scanner := bufio.NewScanner(f)
	scratch_cards := make([]*ScratchCard, 0)

	for scanner.Scan() {
		t := scanner.Text()
		ReadCard(t)
		scratch_card := ReadCard(t)

		logrus.WithFields(logrus.Fields{
			"Card": scratch_card,
		}).Debug("Loaded card")

		scratch_cards = append(scratch_cards, scratch_card)
	}

	if config.First {
		return RunPart1(scratch_cards), nil
	} else {
		return RunPart2(scratch_cards), nil
	}
}
