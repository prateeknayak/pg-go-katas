package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type InputCommand string

const (
	PLACE  InputCommand = "PLACE"
	MOVE   InputCommand = "MOVE"
	LEFT   InputCommand = "LEFT"
	RIGHT  InputCommand = "RIGHT"
	REPORT InputCommand = "REPORT"
)

type Direction string

const (
	NORTH Direction = "NORTH"
	SOUTH Direction = "SOUTH"
	EAST  Direction = "EAST"
	WEST  Direction = "WEST"
)

type Position struct {
	x   int
	y   int
	dir Direction
}

var globalState Position

func main() {

	fmt.Println("starting toy robot")
	globalState = Position{
		x:   0,
		y:   0,
		dir: NORTH,
	}
	err := run(os.Stdin, &globalState)
	if err != nil {
		fmt.Printf("unable to run the toy robot")
		os.Exit(1)
	}

}

func run(r io.Reader, state *Position) error {
	var count int
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		cmd := InputCommand(strings.ToUpper(input[0]))

		switch cmd {
		case PLACE:
			err := place(state, input[1])
			if err != nil {
				return fmt.Errorf("please provide correct input, %v", err)
			}
			count++
		case MOVE:
			if count != 0 {
				move(state)
			}
		case LEFT:
			if count != 0 {
				left(state)
			}
		case RIGHT:
			if count != 0 {
				right(state)
			}
		case REPORT:
			if count != 0 {
				str := report(state)
				fmt.Println(str)
			}
		default:
			return fmt.Errorf("unknown command specified")
		}
	}
	return nil
}

func place(state *Position, args string) error {
	s := strings.Split(args, ",")

	if len(s) < 3 {
		return fmt.Errorf("invalid input given")
	}

	x, err := strconv.Atoi(s[0])
	if err != nil {
		return fmt.Errorf("invalid x co-ordinate: %v", err)
	}

	y, err := strconv.Atoi(s[1])
	if err != nil {
		return fmt.Errorf("invalid y co-ordinate: %v", err)
	}
	dir := Direction(strings.ToUpper(s[2]))
	switch dir {
	case NORTH:
	case SOUTH:
	case EAST:
	case WEST:
		state.dir = dir
	default:
		return fmt.Errorf("invalid direction provided")
	}

	state.x = x
	state.y = y

	return nil
}

func left(state *Position) error {
	switch state.dir {
	case NORTH:
		state.dir = WEST
		return nil
	case SOUTH:
		state.dir = EAST
		return nil
	case EAST:
		state.dir = NORTH
		return nil
	case WEST:
		state.dir = SOUTH
		return nil
	}
	return fmt.Errorf("could not determinse direction")
}

func right(state *Position) error {
	switch state.dir {
	case NORTH:
		state.dir = EAST
		return nil
	case SOUTH:
		state.dir = WEST
		return nil
	case EAST:
		state.dir = NORTH
		return nil
	case WEST:
		state.dir = SOUTH
		return nil
	}
	return fmt.Errorf("could not determinse direction")
}

func move(state *Position) error {
	switch state.dir {
	case NORTH:
		state.y += 1
		return nil
	case SOUTH:
		state.y += 1
		return nil
	case EAST:
		state.x += 1
		return nil
	case WEST:
		state.x += 1
		return nil
	}
	return fmt.Errorf("could not determinse direction")
}

func report(state *Position) string {
	return fmt.Sprintf("Output: %d,%d,%s", state.x, state.y, state.dir)
}

//
//type Command struct {
//	cmd      string
//	position []string
//}
//
//type Game struct {
//	scanner *bufio.Scanner
//}
//type CommandReader interface {
//	read() (*Command, error)
//}
//
//type CommandValidator interface {
//	validate() error
//}
//
//type RobotMover interface {
//	move() error
//}
//
//type PositionReporter interface {
//	report() (string, error)
//}
//
//type Runner interface {
//	CommandReader
//	CommandValidator
//	RobotMover
//	PositionReporter
//}
//
//func moveRobot(cr Runner) error {
//	c, err := cr.read()
//	if err != nil {
//		return fmt.Errorf("cannot read command; %v", err)
//	}
//	_ = c
//	return nil
//}
//
///*
//	Reader
//*/
//
//func (g *Game) read() (*Command, error) {
//	var c string
//	for g.scanner.Scan() {
//		c = g.scanner.Text()
//		if strings.Contains(c, "PLACE") {
//			fmt.Println(c)
//		}
//	}
//	return nil, nil
//}
