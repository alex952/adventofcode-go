package day7

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"alex952.com/advent2023/adventcode2023/shared"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/maps"
)

type Card struct {
	S string
}

func (c Card) Symbol() string {
	return c.S

}
func (c *Card) Value() int {
	switch c.S {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	default:
		v, _ := strconv.Atoi(c.S)
		return v
	}
}

type Hand struct {
	Cards []Card
	Bid   int
}
type CardPart2 struct {
	S string
}

func (c CardPart2) Symbol() string {
	return c.S

}
func (c *CardPart2) Value() int {
	switch c.S {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 1
	case "T":
		return 10
	default:
		v, _ := strconv.Atoi(c.S)
		return v
	}
}

type HandPart2 struct {
	Cards []CardPart2
	Bid   int
}

func (h HandPart2) Value() ([]int, map[int]int) {
	values := make(map[int]int)
	aces := 0

	for _, h := range h.Cards {
		if h.Symbol() == "J" {
			aces += 1
			continue
		}

		if n, exists := values[h.Value()]; exists {
			values[h.Value()] = n + 1
		} else {
			values[h.Value()] = 1
		}
	}

	keys := maps.Keys(values)

	sort.SliceStable(keys, func(i, j int) bool {
		if values[keys[i]] == values[keys[j]] {
			if keys[i] > keys[j] {
				return true
			} else {
				return false
			}
		}
		return values[keys[i]] > values[keys[j]]
	})

	if len(keys) > 0 {
		values[keys[0]] += aces
	} else {
		values['1'] = aces
		keys = append(keys, '1')

	}

	logrus.WithField("Hand", values).WithField("Order", keys).Debug("The hand has this structure")

	return keys, values

}

// Returns a map of value of card to number of cards
func (h Hand) Value() ([]int, map[int]int) {
	values := make(map[int]int)

	for _, h := range h.Cards {
		if n, exists := values[h.Value()]; exists {
			values[h.Value()] = n + 1
		} else {
			values[h.Value()] = 1
		}
	}

	keys := maps.Keys(values)

	sort.SliceStable(keys, func(i, j int) bool {
		if values[keys[i]] == values[keys[j]] {
			if keys[i] > keys[j] {
				return true
			} else {
				return false
			}
		}
		return values[keys[i]] > values[keys[j]]
	})

	logrus.WithField("Hand", values).WithField("Order", keys).Debug("The hand has this structure")

	return keys, values

}

func compareHandsPart2(left HandPart2, right HandPart2) bool {
	left_keys, left_value := left.Value()
	right_keys, right_value := right.Value()

	for i := 0; i < len(left_keys); i = i + 1 {
		if left_value[left_keys[i]] != right_value[right_keys[i]] {
			return left_value[left_keys[i]] > right_value[right_keys[i]]
		}
	}

	for i := 0; i < len(left.Cards); i = i + 1 {
		if left.Cards[i].Value() != right.Cards[i].Value() {
			return left.Cards[i].Value() > right.Cards[i].Value()
		}
	}

	return true
}

func compareHands(left Hand, right Hand) bool {
	left_keys, left_value := left.Value()
	right_keys, right_value := right.Value()

	for i := 0; i < len(left_keys); i = i + 1 {
		if left_value[left_keys[i]] != right_value[right_keys[i]] {
			return left_value[left_keys[i]] > right_value[right_keys[i]]
		}
	}

	for i := 0; i < len(left.Cards); i = i + 1 {
		if left.Cards[i].Value() != right.Cards[i].Value() {
			return left.Cards[i].Value() > right.Cards[i].Value()
		}
	}

	return true
}

func RunPart1(hands []Hand) string {
	sort.SliceStable(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j])
	})

	sum := 0

	for idx, h := range hands {
		logrus.WithField("Ordered hands", h).Debug("Ordered")
		sum = sum + ((len(hands) - idx) * h.Bid)
	}

	return fmt.Sprintf("%d", sum)
}

func RunPart2(hands []HandPart2) string {
	sort.SliceStable(hands, func(i, j int) bool {
		return compareHandsPart2(hands[i], hands[j])
	})

	sum := 0

	for idx, h := range hands {
		logrus.WithField("Ordered hands", h).Debug("Ordered")
		sum = sum + ((len(hands) - idx) * h.Bid)
	}

	return fmt.Sprintf("%d", sum)
}

func ReadHands(line string) Hand {
	c := make([]Card, 0)

	parts := strings.Split(line, " ")
	symbols := strings.Split(parts[0], "")

	for _, s := range symbols {
		c = append(c, Card{S: s})
	}

	bid, _ := strconv.Atoi(parts[1])

	return Hand{Cards: c, Bid: bid}
}

func ReadHandsPart2(line string) HandPart2 {
	c := make([]CardPart2, 0)

	parts := strings.Split(line, " ")
	symbols := strings.Split(parts[0], "")

	for _, s := range symbols {
		c = append(c, CardPart2{S: s})
	}

	bid, _ := strconv.Atoi(parts[1])

	return HandPart2{Cards: c, Bid: bid}
}

type Day7Runner struct{}

func (runner *Day7Runner) RunChallenge(config shared.AdventOfCodeChallengRunnerConfig) (string, error) {
	f, err := os.Open(config.Filename)
	if err != nil {
		fmt.Println("Cannot open file")
		return "", errors.New("couldn't run the challenge. Can't open the file")
	}

	scanner := bufio.NewScanner(f)

	if config.First {
		hands := make([]Hand, 0)
		for scanner.Scan() {
			t := scanner.Text()

			hands = append(hands, ReadHands(t))
		}

		return RunPart1(hands), nil
	} else {
		hands := make([]HandPart2, 0)
		for scanner.Scan() {
			t := scanner.Text()

			hands = append(hands, ReadHandsPart2(t))
		}

		return RunPart2(hands), nil

	}
}
