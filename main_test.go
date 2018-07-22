package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaceGoodInput(t *testing.T) {
	s := "PLACE 1,2,NORTH"

	pos, err := place(strings.Split(s, " ")[1])

	assert.Nil(t, err)
	assert.NotNil(t, pos)
	assert.Equal(t, pos.x, 1)
	assert.Equal(t, pos.y, 2)
	assert.Equal(t, pos.dir, NORTH)
}

func TestPlaceBadInputXNoInt(t *testing.T) {
	s := "PLACE banana,2,NORTH"

	pos, err := place(strings.Split(s, " ")[1])

	assert.NotNil(t, err)
	assert.Nil(t, pos)
}

func TestPlaceBadInputYNoInt(t *testing.T) {
	s := "PLACE 1,banana,NORTH"

	pos, err := place(strings.Split(s, " ")[1])

	assert.NotNil(t, err)
	assert.Nil(t, pos)
}

func TestPlaceBadInputNoDir(t *testing.T) {
	s := "PLACE 1,2,banana"

	pos, err := place(strings.Split(s, " ")[1])

	assert.NotNil(t, err)
	assert.Nil(t, pos)
}

func TestPlaceBadInput(t *testing.T) {
	s := "PLACE 000"

	pos, err := place(strings.Split(s, " ")[1])

	assert.NotNil(t, err)
	assert.Nil(t, pos)
}

func TestRunGoodInput(t *testing.T) {
	s := strings.NewReader("PLACE 1,0,NORTH")
	testState := &Position{}
	err := run(s, testState)
	assert.Nil(t, err)
}

func TestRunWrongCommand(t *testing.T) {
	s := strings.NewReader("Banana")
	testState := &Position{}
	err := run(s, testState)

	assert.NotNil(t, err)
}
