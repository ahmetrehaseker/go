package rover

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateNewGrid(t *testing.T) {
	assert.Equal(t,
		NewGrid(4, 5).Field,
		[][]int{
			[]int{0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0}})
}

func TestShouldSnapRoverToGrid(t *testing.T) {
	grid := NewGrid(4, 5)
	rover := New(Coordinates{1, 2}, North)

	grid.Snap(rover)
	assert.Equal(t, grid.At(rover.Coords), ROVER)
}

func TestShouldOverflowWidth(t *testing.T) {
	grid := NewGrid(4, 4)
	newPosition, _ := grid.OverflowPosition(Coordinates{3, 3}, 1, 0)
	assert.Equal(t, newPosition, Coordinates{0, 3})
}

func TestShouldOverflowHeight(t *testing.T) {
	grid := NewGrid(4, 4)
	newPosition, _ := grid.OverflowPosition(Coordinates{3, 3}, 0, 1)
	assert.Equal(t, newPosition, Coordinates{3, 0})
}

func TestShouldReturnTrueWhenNoCollisionImminent(t *testing.T) {
	grid := NewGrid(4, 4)
	grid.Insert(Coordinates{1, 0}, NOTHING)
	newPosition, ok := grid.OverflowPosition(Coordinates{0, 0}, 1, 0)
	assert.Equal(t, ok, true)
	assert.Equal(t, newPosition, Coordinates{1, 0})
}

func TestShouldNotMoveWhenColliding(t *testing.T) {
	grid := NewGrid(4, 4)
	rover := New(Coordinates{0, 0}, North)
	grid.Snap(rover)
	rover.Grid.Insert(Coordinates{0, 1}, OBSTACLE)

	rover.Advance()
	assert.Equal(t, rover.Coords, Coordinates{0, 0})
}

func TestShouldReturnFalseWhenCollisionWithObstacleImminent(t *testing.T) {
	grid := NewGrid(4, 4)
	grid.Insert(Coordinates{1, 0}, OBSTACLE)
	newPosition, ok := grid.OverflowPosition(Coordinates{0, 0}, 1, 0)
	assert.Equal(t, ok, false)
	assert.Equal(t, newPosition, Coordinates{0, 0})
}
