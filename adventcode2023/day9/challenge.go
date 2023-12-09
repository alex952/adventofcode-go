package day9

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"alex952.com/advent2023/adventcode2023/shared"
	"github.com/sirupsen/logrus"
)

type Day9Runner struct{}

type InstabilitySensor []int

func ReadData(s *bufio.Scanner) []InstabilitySensor {
	data := make([]InstabilitySensor, 0)

	num_regex := regexp.MustCompile(`([-\d]+)`)

	for s.Scan() {
		d := make(InstabilitySensor, 0)
		t := s.Text()
		matches := num_regex.FindAllString(t, -1)

		for _, m := range matches {
			v, _ := strconv.Atoi(m)
			d = append(d, v)
		}

		data = append(data, d)
	}

	logrus.WithField("Data", data).Debug("Read data")

	return data
}

func checkFinished(x []int, isFinished func(int) bool) bool {
	for _, v := range x {
		if !isFinished(v) {
			return false
		}
	}
	return true
}

func findNextValue(i InstabilitySensor, first bool) int {
	new_i := make(InstabilitySensor, 0)

	if checkFinished(i, func(x int) bool { return x == 0 }) {
		return 0
	}

	for idx := 1; idx < len(i); idx += 1 {
		new_i = append(new_i, i[idx]-i[idx-1])
	}

	if first {
		return i[len(i)-1] + findNextValue(new_i, first)
	} else {
		return i[0] - findNextValue(new_i, first)
	}

}

func Run(data []InstabilitySensor, first bool) string {
	sum := 0

	for _, d := range data {
		next_value := findNextValue(d, first)
		logrus.WithFields(logrus.Fields{
			"Sequence": d,
			"Next":     next_value,
		}).Debug("Found Next Value")

		sum += next_value
	}

	return fmt.Sprintf("%d", sum)
}

func (runner *Day9Runner) RunChallenge(config shared.AdventOfCodeChallengRunnerConfig) (string, error) {
	f, err := os.Open(config.Filename)
	if err != nil {
		fmt.Println("Cannot open file")
		return "", errors.New("couldn't run the challenge. Can't open the file")
	}

	scanner := bufio.NewScanner(f)

	data := ReadData(scanner)

	return Run(data, config.First), nil
}
