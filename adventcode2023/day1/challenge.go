package day1

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

var regexNumbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var numbersMap = map[string]byte{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

type CalibrationInput struct {
	calibrationValues []byte
}

func findFirstMatch(input []byte) []int {
	var first_match []int

	for _, value_to_find := range regexNumbers {
		re := regexp.MustCompile(value_to_find)
		var match []int

		if matches := re.FindAllIndex(input, -1); len(matches) == 0 {
			match = make([]int, 0)
		} else {
			match = matches[0]
		}

		if len(first_match) == 0 {
			first_match = match
		} else if len(match) == 2 && match[0] < first_match[0] {
			first_match = match
		}
	}

	return first_match
}

func findLastMatch(input []byte) []int {
	var last_match []int

	for _, value_to_find := range regexNumbers {
		re := regexp.MustCompile(value_to_find)
		var match []int

		if matches := re.FindAllIndex(input, -1); len(matches) == 0 {
			match = make([]int, 0)
		} else {
			match = matches[len(matches)-1]
		}

		if len(last_match) == 0 {
			last_match = match
		} else if len(match) == 2 && match[0] > last_match[0] {
			last_match = match
		}
	}

	return last_match
}

func strToByte(input []byte) byte {
	return numbersMap[string(input)]
}

func (input *CalibrationInput) firstPartFirstInt() byte {
	for i := 0; i < len(input.calibrationValues); i = i + 1 {
		if _, err := strconv.Atoi(string(input.calibrationValues[i])); err == nil {
			return input.calibrationValues[i]
		}

	}

	return '0'
}
func (input *CalibrationInput) firstInt() byte {
	first_match := findFirstMatch(input.calibrationValues)

	for i := 0; i < len(input.calibrationValues); i = i + 1 {
		if _, err := strconv.Atoi(string(input.calibrationValues[i])); err == nil {
			if len(first_match) == 2 && first_match[0] < i {
				return strToByte(input.calibrationValues[first_match[0]:first_match[1]])

			} else {
				return input.calibrationValues[i]
			}
		}

	}

	if len(first_match) == 2 {
		return strToByte(input.calibrationValues[first_match[0]:first_match[1]])
	}

	return '0'
}

func (input *CalibrationInput) firstPartLastInt() byte {
	for i := len(input.calibrationValues) - 1; i >= 0; i = i - 1 {
		if _, err := strconv.Atoi(string(input.calibrationValues[i])); err == nil {
			return input.calibrationValues[i]
		}

	}

	return '0'
}
func (input *CalibrationInput) lastInt() byte {
	last_match := findLastMatch(input.calibrationValues)

	for i := len(input.calibrationValues) - 1; i >= 0; i = i - 1 {
		if _, err := strconv.Atoi(string(input.calibrationValues[i])); err == nil {
			if len(last_match) == 2 && last_match[0] > i {
				return strToByte(input.calibrationValues[last_match[0]:last_match[1]])

			} else {
				return input.calibrationValues[i]
			}
		}

	}

	if len(last_match) == 2 {
		return strToByte(input.calibrationValues[last_match[0]:last_match[1]])
	}

	return '0'
}

func (input *CalibrationInput) FirstPartCalibrationNumber() int {
	r_slice := []byte{
		input.firstPartFirstInt(),
		input.firstPartLastInt(),
	}

	if calibration, err := strconv.Atoi(string(r_slice)); err != nil {
		return 0
	} else {
		return calibration
	}
}

func (input *CalibrationInput) CalibrationNumber() int {
	r_slice := []byte{
		input.firstInt(),
		input.lastInt(),
	}

	if calibration, err := strconv.Atoi(string(r_slice)); err != nil {
		return 0
	} else {
		return calibration
	}
}

func FirstPartTotalCalibration(inputs []*CalibrationInput) int {
	sum := 0
	for _, input := range inputs {
		sum += input.FirstPartCalibrationNumber()
	}

	return sum
}

func TotalCalibration(inputs []*CalibrationInput) int {
	sum := 0
	for _, input := range inputs {
		sum += input.CalibrationNumber()
	}

	return sum
}

func ReadCalibrationLine(line string) *CalibrationInput {
	return &CalibrationInput{calibrationValues: []byte(line)}
}

type Day1Runner struct {
}

func (d *Day1Runner) RunChallenge(config shared.AdventOfCodeChallengRunnerConfig) (string, error) {
	f, err := os.Open(config.Filename)
	if err != nil {
		fmt.Println("Cannot open file")
		return "", errors.New("Couldn't run the challenge. Can't open the file")
	}

	calibration_data := []*CalibrationInput{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
		calibration_data = append(calibration_data, ReadCalibrationLine(t))
	}

	for _, e := range calibration_data {
		if config.First {
			logrus.Debug(e.FirstPartCalibrationNumber())
		} else {
			logrus.Debug(e.CalibrationNumber())
		}
	}

	if config.First {
		return fmt.Sprintf("Total calibration number is %d\n", FirstPartTotalCalibration(calibration_data)), nil
	} else {
		return fmt.Sprintf("Total calibration number is %d\n", TotalCalibration(calibration_data)), nil
	}
}
