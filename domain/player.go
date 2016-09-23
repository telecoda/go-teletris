package domain

import (
	"fmt"
	"math/rand"
)

type Player struct {
	Score     int
	Level     int
	TotalRows int
	state     PlayerState
	X, Y      int
	shape     *Shape
	nextShape *Shape
}

func NewPlayer() *Player {
	player := &Player{
		Level:     1,
		Score:     0,
		TotalRows: 0,
		state:     Alive,
		X:         BoardWidth / 2,
		Y:         BoardHeight - 3,
		shape:     nil,
		nextShape: nil,
	}

	return player
}

func (p *Player) GetShapeBlocks() []*Block {
	if p.shape != nil {
		return p.shape.GetBlocks()
	}
	return nil
}

func (p *Player) GetNextShapeBlocks() []*Block {
	if p.nextShape != nil {
		return p.nextShape.GetBlocks()
	}
	return nil
}

func (p *Player) MoveDown() {
	p.Y -= 1
}

func (p *Player) MoveLeft() {
	p.X -= 1
}

func (p *Player) MoveRight() {
	p.X += 1
}

func (p *Player) Rotate() {
	p.shape.Rotate()
}

func (p *Player) RotateBack() {
	p.shape.RotateBack()
}

func (p *Player) setNextRandomShape() {
	// copy next shape
	p.shape = p.nextShape
	// random colour, not empty or grey
	colour := BlockColour(rand.Intn(Purple) + 1)
	shapeType := ShapeType(rand.Intn(T) + 1)
	switch shapeType {
	case Square:
		p.nextShape = SquareShape(colour)
	case Bar:
		p.nextShape = BarShape(colour)
	case LeftL:
		p.nextShape = LeftLShape(colour)
	case RightL:
		p.nextShape = RightLShape(colour)
	case LeftStep:
		p.nextShape = LeftStepShape(colour)
	case RightStep:
		p.nextShape = RightStepShape(colour)
	case T:
		p.nextShape = TShape(colour)
	default:
		err := fmt.Sprintf("Unexpected shape type: %d", shapeType)
		panic(err)
	}

	// position at top middle of board
	p.X = BoardWidth / 2
	p.Y = BoardHeight - 3
}
