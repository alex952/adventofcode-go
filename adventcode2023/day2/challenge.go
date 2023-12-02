package day2

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

var colors = []string{"red", "blue", "green"}

type CubeGameData struct {
	Id    int
	Cubes map[string]int
}

func ReadCubGameData(line string) *CubeGameData {
	data := CubeGameData{Cubes: map[string]int{}}

	id_regex := regexp.MustCompile(`^Game ([0-9]+)`)
	if id_match := id_regex.FindStringSubmatch(line); id_match != nil {
		data.Id, _ = strconv.Atoi(id_match[1])
	}

	for _, color := range colors {
		color_regex := regexp.MustCompile(fmt.Sprintf("([0-9]+) %s", color))

		color_matches := color_regex.FindAllStringSubmatch(line, -1)
		if color_matches == nil {
			continue
		}

		for _, color_match := range color_matches {
			color_n_match, _ := strconv.Atoi(color_match[1])
			if color_n, exists := data.Cubes[color]; !exists {
				data.Cubes[color] = color_n_match
			} else if color_n_match > color_n {
				data.Cubes[color] = color_n_match
			}
		}
	}

	return &data
}

func RunPart1(data []*CubeGameData) string {
	restrictions := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	ret_sum := 0

	for _, it := range data {
		ok := true
		for color, max := range restrictions {
			if color_n, exists := it.Cubes[color]; exists && color_n > max {
				logrus.WithFields(logrus.Fields{
					"Cubes": it.Cubes,
					"ID":    it.Id,
					"Color": color,
				}).Debug("Game exceeds max for color")
				ok = false
				break
			}
		}

		if ok {
			logrus.WithFields(logrus.Fields{
				"Cubes": it.Cubes,
				"ID":    it.Id,
			}).Debug("Game fulfills restrictions, adding to total")
			ret_sum += it.Id
		}
	}

	return fmt.Sprintf("%d", ret_sum)
}

func RunPart2(data []*CubeGameData) string {
	ret_sum := 0

	for _, it := range data {
		mult := 1
		for _, v := range it.Cubes {
			mult *= v
		}

		logrus.WithFields(logrus.Fields{
			"Cubes": it.Cubes,
			"ID":    it.Id,
		}).Debug(fmt.Sprintf("Total for game is: %d", mult))

		ret_sum += mult
	}

	return fmt.Sprintf("%d", ret_sum)
}

type Day2Runner struct{}

func (runner *Day2Runner) RunChallenge(config shared.AdventOfCodeChallengRunnerConfig) (string, error) {
	f, err := os.Open(config.Filename)
	if err != nil {
		fmt.Println("Cannot open file")
		return "", errors.New("Couldn't run the challenge. Can't open the file")
	}

	scanner := bufio.NewScanner(f)
	cube_data := []*CubeGameData{}

	for scanner.Scan() {
		t := scanner.Text()
		cube_game_data := ReadCubGameData(t)
		logrus.WithFields(logrus.Fields{
			"ID":    cube_game_data.Id,
			"Cubes": cube_game_data.Cubes,
		}).Debug("Cube Game Data Loaded")
		cube_data = append(cube_data, cube_game_data)
	}

	if config.First {
		return RunPart1(cube_data), nil
	} else {
		return RunPart2(cube_data), nil
	}
}
