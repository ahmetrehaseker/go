package rover

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRover(t *testing.T) {
	rover := New(Coordinates{0, 0}, South)
	assert.Equal(t, rover.Coords, Coordinates{0, 0})
	assert.Equal(t, rover.Facing, South)
}

func TestRoverShouldBeFacingInitialDirection(t *testing.T) {
	rover := New(Coordinates{0, 0}, South)
	assert.Equal(t, rover.Facing, South)
}

func TestShouldAdvanceNorth(t *testing.T) {
	rover := New(Coordinates{0, 0}, North)
	rover.Advance()
	assert.Equal(t, rover.Coords, Coordinates{0, 1})
}

func TestShouldAdvanceEast(t *testing.T) {
	rover := New(Coordinates{0, 0}, East)
	rover.Advance()
	assert.Equal(t, rover.Coords, Coordinates{1, 0})
}

func TestShouldAdvanceWest(t *testing.T) {
	rover := New(Coordinates{1, 0}, West)
	rover.Advance()
	assert.Equal(t, rover.Coords, Coordinates{0, 0})
}

func TestShouldAdvanceSouth(t *testing.T) {
	rover := New(Coordinates{0, 1}, South)
	rover.Advance()
	assert.Equal(t, rover.Coords, Coordinates{0, 0})
}

func TestShouldRetreatNorth(t *testing.T) {
	rover := New(Coordinates{0, 1}, North)
	rover.Retreat()
	assert.Equal(t, rover.Coords, Coordinates{0, 0})
}

func TestShouldRetreatEast(t *testing.T) {
	rover := New(Coordinates{1, 0}, East)
	rover.Retreat()
	assert.Equal(t, rover.Coords, Coordinates{0, 0})
}

func TestShouldRetreatWest(t *testing.T) {
	rover := New(Coordinates{0, 0}, West)
	rover.Retreat()
	assert.Equal(t, rover.Coords, Coordinates{1, 0})
}

func TestShouldRetreatSouth(t *testing.T) {
	rover := New(Coordinates{0, 0}, South)
	rover.Retreat()
	assert.Equal(t, rover.Coords, Coordinates{0, 1})
}

func TestShouldTurnRight(t *testing.T) {
	rover := New(Coordinates{0, 0}, North)
	rover.TurnRight()
	assert.Equal(t, rover.Facing, East)
	rover.TurnRight()
	assert.Equal(t, rover.Facing, South)
	rover.TurnRight()
	assert.Equal(t, rover.Facing, West)
	rover.TurnRight()
	assert.Equal(t, rover.Facing, North)
}

func TestShouldTurnLeft(t *testing.T) {
	rover := New(Coordinates{0, 0}, North)
	rover.TurnLeft()
	assert.Equal(t, rover.Facing, West)
	rover.TurnLeft()
	assert.Equal(t, rover.Facing, South)
	rover.TurnLeft()
	assert.Equal(t, rover.Facing, East)
	rover.TurnLeft()
	assert.Equal(t, rover.Facing, North)
}

func TestShouldStaySnappedToGrid(t *testing.T) {
	rover := New(Coordinates{0, 0}, North)
	rover.Advance()
	assert.Equal(t, rover.Grid.At(Coordinates{0, 0}), NOTHING)
	assert.Equal(t, rover.Grid.At(rover.Coords), ROVER)
}

func TestShouldAdvanceWithOverflow(t *testing.T) {
	rover := New(Coordinates{99, 99}, North)
	rover.Advance()
	assert.Equal(t, rover.Coords, Coordinates{99, 0})
}

func TestShouldRetreatWithOverflow(t *testing.T) {
	rover := New(Coordinates{99, 99}, East)
	rover.Advance()
	assert.Equal(t, rover.Coords, Coordinates{0, 99})
}

func TestShouldReportObstacle(t *testing.T) {
	rover := New(Coordinates{0, 0}, North)
	rover.Grid.Insert(Coordinates{1, 1}, OBSTACLE)
	rover.Advance()
	rover.TurnRight()
	rover.Advance()
	assert.Equal(t, rover.Coords, Coordinates{0, 1})
}
