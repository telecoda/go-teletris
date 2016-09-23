package domain

var GameTitle = "Teletris"

// Board constants
const (
	BoardWidth      = 12
	BoardHeight     = 22
	BoardOffsetX    = 32
	BoardOffsetY    = 32
	BlockPixels     = 40
	NextBlockPixels = 20
	NextOffsetY     = 35
)

var ArrowPixels = 64
var InfoPanelWidth = (BoardWidth - 1) * BlockPixels

// score constants
const (
	DigitsWidth       = 30
	DigitsHeight      = 40
	AudioButtonWidth  = 40
	AudioButtonHeight = 40
	MaxScoreDigits    = 6
	MaxLevelDigits    = 2
	ScorePerRow       = 5
)

// speed
const (
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
	Pink
	Blue
	Yellow
	Green
	Purple
	Grey
)

var SpriteNames map[BlockColour]string = map[BlockColour]string{
	Empty:  "empty_block.png",
	Red:    "red_block_gopher.png",
	Pink:   "pink_block_gopher.png",
	Blue:   "blue_block_gopher.png",
	Yellow: "yellow_block_gopher.png",
	Green:  "green_block_gopher.png",
	Purple: "purple_block_gopher.png",
	Grey:   "grey_block.png",
}

type ShapeType int

const (
	Square = iota
	Bar
	LeftL
	RightL
	LeftStep
	RightStep
	T
)

type GameState int

const (
	Menu GameState = iota
	Playing
	Suspended
	GameOver
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
