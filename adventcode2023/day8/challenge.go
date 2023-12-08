package day8

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"alex952.com/advent2023/adventcode2023/shared"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/maps"
)

type Direction string

const (
	Left  Direction = "L"
	Right Direction = "R"
)

type Node string

type Instruction struct {
	left  Node
	right Node
}

type Map map[Node]Instruction

func (i Instruction) Navigate(d Direction) Node {
	if d == Left {
		return i.left
	} else {
		return i.right
	}
}

func ReadMap(s *bufio.Scanner) (Map, []Direction) {
	m := Map{}
	s.Scan()

	// Directions
	directions := strings.Split(s.Text(), "")
	d := make([]Direction, 0)
	for _, x := range directions {
		if x == string(Left) {
			d = append(d, Left)
		} else {
			d = append(d, Right)
		}
	}

	logrus.WithField("Directions", d).Debug("Directions Loaded")

	r := regexp.MustCompile(`^([A-Z1-9]{3}).*=.*\(([A-Z1-9]{3}).*([A-Z1-9]{3})\)$`)

	for s.Scan() {
		t := s.Text()
		if matches := r.FindAllStringSubmatch(t, -1); len(matches) > 0 {
			logrus.WithFields(logrus.Fields{
				"Line":    t,
				"Matches": matches,
			}).Debug("Direction line matches")

			i := Instruction{left: Node(matches[0][2]), right: Node(matches[0][3])}
			m[Node(matches[0][1])] = i
		} else {
			logrus.WithFields(logrus.Fields{
				"Line":    t,
				"Matches": matches,
			}).Debug("Direction line matches")
		}
	}

	logrus.WithField("Map", m).Debug("Map Loaded")

	return m, d
}

func findSteps(m Map, d []Direction, initial Node, finishedFunc func(Node) bool) int {
	steps := 0
	next := initial
	node := m[next]

	for {
		for _, dir := range d {
			next := node.Navigate(dir)

			steps += 1
			if finishedFunc(next) {
				return steps
			}
			node = m[next]
		}
	}
}

func finishedPart1(n Node) bool {
	if n == "ZZZ" {
		return true
	} else {
		return false
	}
}

func finishedPart2(n Node) bool {
	if n[len(n)-1] == 'Z' {
		return true
	} else {
		return false
	}
}
func RunPart1(m Map, d []Direction) string {
	steps := findSteps(m, d, Node("AAA"), finishedPart1)
	return fmt.Sprintf("%d", steps)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func RunPart2(m Map, d []Direction) string {
	steps := make([]int, 0)
	initial_nodes := make([]Node, 0)

	for _, k := range maps.Keys(m) {
		if k[len(k)-1] == 'A' {
			initial_nodes = append(initial_nodes, k)
		}
	}

	for _, ini_node := range initial_nodes {
		logrus.WithField("Node", ini_node).Debug("Finding steps for node")
		steps = append(steps, findSteps(m, d, ini_node, finishedPart2))
		logrus.WithField("Steps", steps).Debug("Steps found")
	}

	lcm := LCM(steps[0], steps[1], steps[2:]...)

	return fmt.Sprintf("%d", lcm)
}

type Day8Runner struct{}

func (runner *Day8Runner) RunChallenge(config shared.AdventOfCodeChallengRunnerConfig) (string, error) {
	f, err := os.Open(config.Filename)
	if err != nil {
		fmt.Println("Cannot open file")
		return "", errors.New("couldn't run the challenge. Can't open the file")
	}

	scanner := bufio.NewScanner(f)

	m, d := ReadMap(scanner)

	if config.First {
		return RunPart1(m, d), nil
	} else {
		return RunPart2(m, d), nil
	}

	return "", nil
}
