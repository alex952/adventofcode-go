package day3

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

type PartsData struct {
	Symbols    map[int]byte
	Numbers    [][]int
	Line       string
	LineNumber int
}

func ReadPartsData(line_number int, line string) *PartsData {
	data := PartsData{Symbols: make(map[int]byte), Numbers: make([][]int, 0)}
	data.Line = line
	data.LineNumber = line_number

	symbol_regex := regexp.MustCompile(`[^a-z^A-Z^\d^.]`)
	if symbols_match := symbol_regex.FindAllStringIndex(line, -1); symbols_match != nil {
		for _, indices := range symbols_match {
			data.Symbols[indices[0]] = line[indices[0]]
		}
	}

	numbers_regex := regexp.MustCompile(`[\d]+`)
	if numbers_match := numbers_regex.FindAllStringIndex(line, -1); numbers_match != nil {
		data.Numbers = numbers_match
	}

	return &data
}

func RunPart1(data map[int]*PartsData) string {
	sum := 0

	for index := 0; index < len(data); index = index + 1 {
		targets := make([]*PartsData, 0)
		prev := index - 1
		next := index + 1

		if prev >= 0 {
			targets = append(targets, data[prev])
		}
		targets = append(targets, data[index])
		if next < len(data) {
			targets = append(targets, data[next])
		}

		d := data[index]
		for _, n := range d.Numbers {
			for _, t := range targets {
			targets:
				for i := n[0] - 1; i < n[1]+1; i = i + 1 {
					if s, exists := t.Symbols[i]; exists {
						number, _ := strconv.Atoi(d.Line[n[0]:n[1]])
						sum += number

						logrus.WithFields(logrus.Fields{
							"Line":   d.Line,
							"Number": number,
							"Symbol": string(s),
						}).Debug("Number Adjacent to Symbol")
						break targets
					}
				}
			}
		}

	}

	return fmt.Sprintf("%d", sum)
}

type SymbolCoordinates struct {
	Line  int
	Index int
}

type NumbersList []int

type SymbolMatches map[SymbolCoordinates]NumbersList

func (matches SymbolMatches) Ratio() int {
	sum := 0
	for _, n := range matches {
		if len(n) == 2 {
			sum += n[0] * n[1]
		}
	}

	return sum
}

func RunPart2(data map[int]*PartsData) string {
	sum := 0
	gears := make(SymbolMatches)

	for index := 0; index < len(data); index = index + 1 {
		targets := make([]*PartsData, 0)
		prev := index - 1
		next := index + 1

		if prev >= 0 {
			targets = append(targets, data[prev])
		}
		targets = append(targets, data[index])
		if next < len(data) {
			targets = append(targets, data[next])
		}

		d := data[index]
		for _, n := range d.Numbers {
			for _, t := range targets {
			targets:
				for i := n[0] - 1; i < n[1]+1; i = i + 1 {
					if s, exists := t.Symbols[i]; exists {
						number, _ := strconv.Atoi(d.Line[n[0]:n[1]])
						sum += number

						logrus.WithFields(logrus.Fields{
							"Line":   d.Line,
							"Number": number,
							"Symbol": string(s),
						}).Debug("Number Adjacent to Symbol")

						coor := SymbolCoordinates{Line: t.LineNumber, Index: i}

						if _, exists := gears[coor]; !exists {
							gears[coor] = make(NumbersList, 0)
						}

						gears[coor] = append(gears[coor], number)

						break targets
					}
				}
			}
		}

	}

	sum = gears.Ratio()

	return fmt.Sprintf("%d", sum)
}

func (p *PartsData) PrintInfo() string {
	s := []byte("Numbers found: ")
	for _, n_indices := range p.Numbers {
		s = fmt.Appendf(s, "%s ", p.Line[n_indices[0]:n_indices[1]])
	}

	return string(s)
}

type Day3Runner struct{}

func (runner *Day3Runner) RunChallenge(config shared.AdventOfCodeChallengRunnerConfig) (string, error) {
	f, err := os.Open(config.Filename)
	if err != nil {
		fmt.Println("Cannot open file")
		return "", errors.New("Couldn't run the challenge. Can't open the file")
	}

	scanner := bufio.NewScanner(f)
	parts_data := make(map[int]*PartsData)

	line := 0
	for scanner.Scan() {
		t := scanner.Text()
		line_parts_data := ReadPartsData(line, t)

		parts_data[line] = line_parts_data

		logrus.WithFields(logrus.Fields{
			"Line":           line,
			"NumbersIndices": line_parts_data.Numbers,
			"SymbolsIndices": line_parts_data.Symbols,
		}).Debug("Read line data")

		fmt.Println(line_parts_data.PrintInfo())

		line += 1
	}

	if config.First {
		return RunPart1(parts_data), nil
	} else {
		return RunPart2(parts_data), nil
	}
}
