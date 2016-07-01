package domain

var GameTitle = "Teletris"

// Board constants
var (
	BoardWidth   = 12
	BoardHeight  = 22
	BoardOffsetX = 32
	BoardOffsetY = 32
	BlockPixels  = 32
)

var ArrowPixels = 64
var InfoPanelWidth = (BoardWidth - 1) * BlockPixels

// speed
var (
	BlockStartSpeed    = 500 // Speed blocks fall in milliseconds
	KeyRepeat          = 150 // Key repeat in milliseconds
	RowsPerLevel       = 5   // increase level every X rows
	LevelSpeedIncrease = 50
)

//Block colours
type BlockColour int

const (
	Empty = iota
	Red
	Green
	Blue
	Yellow
	Magenta
	Cyn
	Grey
)

type ShapeType int

const (
	Square = iota
	Bar
	LeftL
	RightL
	LeftStep
	RightStep
)

type GameState int

const (
	Menu GameState = iota
	Playing
	Paused
)

type PlayerState int

const (
	Alive = iota
	Dead
)

type GameEvent int

const (
	BlockDownEvent = iota
	RowsCompleteEvent
	LevelUpEvent
)

type Alignment int

const (
	Centre = iota
	Left
	Right
	Top
	Middle
	Bottom
)
