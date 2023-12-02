package day2

import (
	"testing"
)

func TestReadGameDataLine(t *testing.T) {
	line := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	data := ReadCubGameData(line)

	if data.Id != 3 {
		t.Fail()
	}

	if red, exists := data.Cubes["red"]; !exists || red != 20 {
		t.Fail()
	}

	if green, exists := data.Cubes["green"]; !exists || green != 13 {
		t.Fail()
	}

	if blue, exists := data.Cubes["blue"]; !exists || blue != 6 {
		t.Fail()
	}
}

func TestRunPart1NoGame(t *testing.T) {
	line := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	data := []*CubeGameData{ReadCubGameData(line)}

	if res := RunPart1(data); res != "0" {
		t.Fail()
	}
}

func TestRunPart1TwoGames(t *testing.T) {
	line := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	line2 := "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"
	line3 := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	data := []*CubeGameData{
		ReadCubGameData(line),
		ReadCubGameData(line2),
		ReadCubGameData(line3),
	}

	if res := RunPart1(data); res != "3" {
		t.Fail()
	}
}

func TestRunPart2(t *testing.T) {
	line := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	line2 := "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"
	data := []*CubeGameData{
		ReadCubGameData(line),
		ReadCubGameData(line2),
	}

	if res := RunPart2(data); res != "60" {
		t.Fail()
	}
}
