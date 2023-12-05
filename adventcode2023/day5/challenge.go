package day5

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"alex952.com/advent2023/adventcode2023/shared"
)

type MapItem struct {
	dest   int
	source int
	length int
}

type Map []MapItem
type TranslationMaps map[string]Map

type PlantData struct {
	seeds []int
	maps  TranslationMaps
}

var MAP_NAMES []string = []string{
	"seed-to-soil map",
	"soil-to-fertilizer map",
	"fertilizer-to-water map",
	"water-to-light map",
	"light-to-temperature map",
	"temperature-to-humidity map",
	"humidity-to-location map",
}

func (m Map) GetTranslatedValue(value int) int {
	for _, it := range m {
		if value >= it.source && value < it.source+it.length {
			return value - it.source + it.dest
		}
	}

	return value
}

func (t TranslationMaps) GetTranslatedValue(value int) int {
	x := value
	for _, map_name := range MAP_NAMES {
		x = t[map_name].GetTranslatedValue(x)
	}

	return x
}

func ReadInput(scanner *bufio.Scanner) *PlantData {
	data := PlantData{seeds: make([]int, 0), maps: make(map[string]Map)}
	scanner.Scan()
	first_line := scanner.Text()

	seeds := strings.Split(strings.TrimSpace(strings.Split(first_line, ":")[1]), " ")

	for _, s := range seeds {
		s_int, _ := strconv.Atoi(s)
		data.seeds = append(data.seeds, s_int)
	}

	name := ""
	for scanner.Scan() {
		t := scanner.Text()

		if l := strings.TrimSpace(t); len(l) == 0 {
			name = ""
		} else {
			if sep_idx := strings.Index(l, ":"); sep_idx > 0 {
				name = l[:sep_idx]
				data.maps[name] = make(Map, 0)
				continue
			} else {
				nums := strings.Split(strings.TrimSpace(l), " ")
				dest, _ := strconv.Atoi(nums[0])
				src, _ := strconv.Atoi(nums[1])
				length, _ := strconv.Atoi(nums[2])
				data.maps[name] = append(data.maps[name], MapItem{source: src, dest: dest, length: length})
			}
		}
	}

	return &data
}

func RunPart1(data *PlantData) string {
	var m *int = nil
	for _, seed := range data.seeds {
		x := data.maps.GetTranslatedValue(seed)

		if m == nil || x < *m {
			m = &x
		}
	}

	return fmt.Sprintf("%d", *m)
}

func RunPart2(data *PlantData) string {
	var m *int = nil

	for i := 0; i < len(data.seeds); i += 2 {
		seed_range := data.seeds[i : i+2]
		fmt.Println(seed_range)
		for x := 0; x < seed_range[1]; x += 1 {
			seed := seed_range[0] + x
			seed_translated := data.maps.GetTranslatedValue(seed)

			if m == nil || seed_translated < *m {
				m = &seed_translated
			}
		}

	}

	return fmt.Sprintf("%d", *m)
}

type Day5Runner struct{}

func (runner *Day5Runner) RunChallenge(config shared.AdventOfCodeChallengRunnerConfig) (string, error) {
	f, err := os.Open(config.Filename)
	if err != nil {
		fmt.Println("Cannot open file")
		return "", errors.New("Couldn't run the challenge. Can't open the file")
	}

	scanner := bufio.NewScanner(f)
	input_data := ReadInput(scanner)

	fmt.Println(input_data)

	if config.First {
		return RunPart1(input_data), nil
	} else {
		return RunPart2(input_data), nil
	}
}
