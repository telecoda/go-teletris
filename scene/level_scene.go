package scene

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"

	"github.com/telecoda/go-teletris/domain"
	"github.com/telecoda/go-teletris/scene/config"
	"github.com/telecoda/go-teletris/scene/io"
	"github.com/telecoda/gomo-simra/simra"
	"github.com/telecoda/gomo-simra/simra/peer"
	"golang.org/x/mobile/exp/sprite"
)

// LevelScene represents a scene object for LevelScene
type LevelScene struct {
	Game          *domain.Game
	background    simra.Sprite
	blockImages   map[domain.BlockColour]*image.RGBA
	blockTextures map[domain.BlockColour]*sprite.SubTex
	boardSprites  [][]*simra.Sprite
	ctrlup        simra.Sprite
	ctrldown      simra.Sprite
	ctrlleft      simra.Sprite
	ctrlright     simra.Sprite
	// buttonState represents which ctrl is pressed (or no ctrl pressed)
	buttonState    int
	buttonReplaced bool

	// images
	backgroundImage image.Image
}

const (
	ctrlNop = iota
	ctrlUp
	ctrlDown
	ctrlLeft
	ctrlRight
)

const (
	ctrlMarginLeft      = 10
	ctrlMarginBottom    = 100
	ctrlMarginBetween   = 10
	buttonMarginRight   = 20
	buttonMarginBottom  = 20
	buttonMarginBetween = 10
)

// Initialize initializes LevelScene scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (l *LevelScene) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// add global touch listener to catch touch end event
	simra.GetInstance().AddTouchListener(l)

	// TODO: when goes to next scene, remove global touch listener
	// simra.GetInstance().RemoveTouchListener(LevelScene)

	// initialize sprites
	l.initSprites()
	l.buttonReplaced = false
	simra.LogDebug("[OUT]")
}

// OnTouchBegin is called when LevelScene scene is Touched.
func (l *LevelScene) OnTouchBegin(x, y float32) {
	// nop
}

// OnTouchMove is called when LevelScene scene is Touched and moved.
func (l *LevelScene) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when LevelScene scene is Touched and it is released.
func (l *LevelScene) OnTouchEnd(x, y float32) {
	l.buttonState = ctrlNop
}

func (l *LevelScene) initSprites() {
	l.initBackground()
	l.initBlockTextures()
	l.initBackgroundTexture()
	l.initBoardBlocks()
	l.initctrlDown()
	l.initctrlUp()
	l.initctrlLeft()
	l.initctrlRight()
}

func (l *LevelScene) initBackground() {
	// add background sprite
	l.background.W = float32(config.ScreenWidth)
	l.background.H = float32(config.ScreenHeight)

	// put center of screen
	l.background.X = config.ScreenWidth / 2
	l.background.Y = config.ScreenHeight / 2

	simra.GetInstance().AddSprite("space_background.png",
		image.Rect(0, 0, int(l.background.W), int(l.background.H)),
		&l.background)

}

func (l *LevelScene) initBlockTextures() {

	l.blockTextures = make(map[domain.BlockColour]*sprite.SubTex, 0)
	l.blockImages = make(map[domain.BlockColour]*image.RGBA, 0)
	rect := image.Rect(0, 0, domain.BlockPixels, domain.BlockPixels)

	for i, name := range domain.SpriteNames {

		blockImage, _, err := io.LoadImage("assets/" + name)
		if err != nil {
			panic(fmt.Sprintf("Error loading image: %s\n", err))
		} else {

			// Save texture for using with Sprites
			tex := peer.GetGLPeer().LoadTextureFromImage(blockImage, rect)
			l.blockTextures[i] = &tex
			// Create and save imageRGBA for offscreen rendering
			bounds := blockImage.Bounds()
			blockRGBA := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
			draw.Draw(blockRGBA, blockRGBA.Bounds(), blockImage, bounds.Min, draw.Src)
			l.blockImages[i] = blockRGBA
		}
	}
}

func (l *LevelScene) initBackgroundTexture() {

	// load image from file
	sourceImage, _, err := io.LoadImage("assets/space_background.png")
	if err != nil {
		panic(fmt.Sprintf("Error loading image: %s\n", err))
	} else {
		gridImage := DrawGrid(sourceImage, 20, 20)
		// draw grey blocks on background
		borderImage := l.drawBorderBlocks(gridImage)
		rect := image.Rect(0, 0, borderImage.Bounds().Dx(), borderImage.Bounds().Dy())
		tex := peer.GetGLPeer().LoadTextureFromImage(borderImage, rect)

		// update background image
		peer.GetSpriteContainer().ReplaceTexture(&l.background.Sprite, tex)

	}
}

func (l *LevelScene) drawBorderBlocks(sourceImage image.Image) image.Image {

	blocks := *l.Game.GetBlocks()
	boardWidth := len(blocks)

	greyBlock := l.blockImages[domain.Grey]
	bounds := sourceImage.Bounds()
	borderImage := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(borderImage, bounds, sourceImage, bounds.Min, draw.Src)

	for x := 0; x < boardWidth; x++ {
		boardHeight := len(blocks[x])
		for y := 0; y < boardHeight; y++ {

			block := blocks[x][y]
			point := image.Point{X: 0, Y: 0}
			// This is a border block so render is on background image
			if block.Colour == domain.Grey {
				log.Printf("Grey block x: %d y: %d\n", x, y)
				xCoord := (x * domain.BlockPixels) + domain.BoardOffsetX
				yCoord := (y * domain.BlockPixels) + domain.BoardOffsetY
				rect := image.Rect(xCoord, yCoord, xCoord+domain.BlockPixels, yCoord+domain.BlockPixels)
				draw.Draw(borderImage, rect, greyBlock, point, draw.Src)
			}
		}
	}

	return borderImage
}

func DrawGrid(sourceImage image.Image, tileWidth int, tileHeight int) image.Image {

	log.Println("Drawing grid over image.")

	lineWidth := 1
	// convert sourceImage to RGBA image
	bounds := sourceImage.Bounds()
	gridImage := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(gridImage, gridImage.Bounds(), sourceImage, bounds.Min, draw.Src)

	lineColour := color.RGBA{0, 0, 0, 255}

	// draw horizontal lines
	for y := 0; y < bounds.Dy(); y += tileHeight {

		lineBounds := image.Rect(0, y, bounds.Dx(), y+lineWidth)
		//lineBounds := &image.Rectangle{Min: {X: 0, Y: 0}, Max: {X: 160, Y: 5}}
		draw.Draw(gridImage, lineBounds, &image.Uniform{lineColour}, image.ZP, draw.Src)

	}

	// draw vertical lines
	for x := 0; x < bounds.Dx(); x += tileWidth {

		lineBounds := image.Rect(x, 0, x+lineWidth, bounds.Dy())
		//lineBounds := &image.Rectangle{Min: {X: 0, Y: 0}, Max: {X: 160, Y: 5}}
		draw.Draw(gridImage, lineBounds, &image.Uniform{lineColour}, image.ZP, draw.Src)

	}

	return gridImage
}

func (l *LevelScene) initBoardBlocks() {

	blocks := *l.Game.GetBlocks()
	boardWidth := len(blocks)
	l.boardSprites = make([][]*simra.Sprite, boardWidth)

	for x := 0; x < boardWidth; x++ {
		boardHeight := len(blocks[x])
		l.boardSprites[x] = make([]*simra.Sprite, boardHeight)
		for y := 0; y < boardHeight; y++ {

			// add background sprite
			block := blocks[x][y]

			// skip border blocks
			if block.Colour == domain.Grey {
				continue
			}
			blockSprite := new(simra.Sprite)
			blockSprite.W = float32(domain.BlockPixels)
			blockSprite.H = float32(domain.BlockPixels)

			// put center of screen
			blockSprite.X = float32(domain.BlockPixels*x + domain.BlockPixels/2 + domain.BoardOffsetX)
			blockSprite.Y = float32(domain.BlockPixels*y + domain.BlockPixels/2 + domain.BoardOffsetY)

			// lookup block sprite
			blockImage := domain.SpriteNames[block.Colour]

			simra.GetInstance().AddSprite(blockImage,
				image.Rect(0, 0, int(domain.BlockPixels), int(domain.BlockPixels)),
				blockSprite)

			l.boardSprites[x][y] = blockSprite

		}
	}
}

func (l *LevelScene) updateBoardBlocks() {

	game := *l.Game
	blocks := *game.GetBlocks()
	boardWidth := len(blocks)

	// update sprite textures based on block colours
	// changed from LoadTexture() to ReplaceTexture() so we don't reload
	// textures from file 60 times a second...
	for x := 0; x < boardWidth; x++ {
		boardHeight := len(blocks[x])
		for y := 0; y < boardHeight; y++ {
			block := blocks[x][y]
			// skip border blocks
			if block.Colour == domain.Grey {
				continue
			}
			boardSprite := l.boardSprites[x][y]

			tex := l.blockTextures[block.Colour]
			peer.GetSpriteContainer().ReplaceTexture(&boardSprite.Sprite, *tex)

		}
	}

	// overlay players blocks onto board
	player := game.Player
	playerBlocks := player.GetShapeBlocks()
	for _, playerBlock := range playerBlocks {
		blockX := playerBlock.X + player.X
		blockY := playerBlock.Y + player.Y
		boardSprite := l.boardSprites[blockX][blockY]
		tex := l.blockTextures[playerBlock.Colour]
		peer.GetSpriteContainer().ReplaceTexture(&boardSprite.Sprite, *tex)
	}
}

// ctrlUp
type ctrlUpTouchListener struct {
	parent *LevelScene
}

func (l *ctrlUpTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("[IN] ctrlUp Begin!")

	ctrl := l.parent
	ctrl.buttonState = ctrlUp

	simra.LogDebug("[OUT]")
}

func (l *ctrlUpTouchListener) OnTouchMove(x, y float32) {
	simra.LogDebug("[IN] ctrlUp Move!")

	ctrl := l.parent
	ctrl.buttonState = ctrlUp

	simra.LogDebug("[OUT]")
}

func (l *ctrlUpTouchListener) OnTouchEnd(x, y float32) {
	simra.LogDebug("[IN] ctrlUp End")

	ctrl := l.parent
	ctrl.buttonState = ctrlNop

	simra.LogDebug("[OUT]")
}

func (l *LevelScene) initctrlUp() {
	// set size of ctrlUp
	l.ctrlup.W = float32(120)
	l.ctrlup.H = float32(120)

	// put ctrlUp on middle bottom of screen
	l.ctrlup.X = (config.ScreenWidth / 2)
	l.ctrlup.Y = ctrlMarginBottom + l.ctrldown.H + ctrlMarginBetween + (l.ctrlup.H / 2)

	// add sprite to glpeer
	simra.GetInstance().AddSprite("arrow.png",
		image.Rect(0, 0, int(l.ctrlup.W), int(l.ctrlup.H)),
		&l.ctrlup)

	// add touch listener for sprite
	ctrlup := &ctrlUpTouchListener{}
	l.ctrlup.AddTouchListener(ctrlup)
	ctrlup.parent = l
}

// ctrlDown
type ctrlDownTouchListener struct {
	parent *LevelScene
}

func (l *ctrlDownTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("[IN] ctrlDown Begin!")

	ctrl := l.parent
	ctrl.buttonState = ctrlDown

	simra.LogDebug("[OUT]")
}

func (l *ctrlDownTouchListener) OnTouchMove(x, y float32) {
	simra.LogDebug("[IN] ctrlDown Move!")

	ctrl := l.parent
	ctrl.buttonState = ctrlDown

	simra.LogDebug("[OUT]")
}

func (l *ctrlDownTouchListener) OnTouchEnd(x, y float32) {
	simra.LogDebug("[IN] ctrlDown End")

	ctrl := l.parent
	ctrl.buttonState = ctrlNop

	simra.LogDebug("[OUT]")
}

func (l *LevelScene) initctrlDown() {
	// set size of ctrlDown
	l.ctrldown.W = float32(120)
	l.ctrldown.H = float32(120)

	// put ctrlDown on middle bottom of screen
	l.ctrldown.X = (config.ScreenWidth / 2)
	l.ctrldown.Y = ctrlMarginBottom + (l.ctrldown.H / 2)

	// rotate arrow to indicate down control
	l.ctrldown.R = math.Pi

	// add sprite to glpeer
	simra.GetInstance().AddSprite("arrow.png",
		image.Rect(0, 0, int(l.ctrldown.W), int(l.ctrldown.H)),
		&l.ctrldown)

	// add touch listener for sprite
	ctrldown := &ctrlDownTouchListener{}
	l.ctrldown.AddTouchListener(ctrldown)
	ctrldown.parent = l
}

// ctrlLeft
type ctrlLeftTouchListener struct {
	parent *LevelScene
}

func (LevelScene *ctrlLeftTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("[IN] ctrlLeft Begin!")

	ctrl := LevelScene.parent
	ctrl.buttonState = ctrlLeft

	simra.LogDebug("[OUT]")
}

func (l *ctrlLeftTouchListener) OnTouchMove(x, y float32) {
	simra.LogDebug("[IN] ctrlLeft Move!")

	ctrl := l.parent
	ctrl.buttonState = ctrlLeft

	simra.LogDebug("[OUT]")
}

func (l *ctrlLeftTouchListener) OnTouchEnd(x, y float32) {
	simra.LogDebug("[IN] ctrlLeft End")

	ctrl := l.parent
	ctrl.buttonState = ctrlNop

	simra.LogDebug("[OUT]")
}

func (l *LevelScene) initctrlLeft() {
	// set size of ctrlLeft
	l.ctrlleft.W = float32(120)
	l.ctrlleft.H = float32(120)

	// put ctrlLeft on left bottom
	l.ctrlleft.X = (config.ScreenWidth / 2) - l.ctrlleft.W - ctrlMarginBetween
	l.ctrlleft.Y = ctrlMarginBottom + l.ctrlleft.H

	// rotate arrow to indicate left control
	l.ctrlleft.R = math.Pi / 2

	// add sprite to glpeer
	simra.GetInstance().AddSprite("arrow.png",
		image.Rect(0, 0, int(l.ctrlleft.W), int(l.ctrlleft.H)),
		&l.ctrlleft)

	// add touch listener for sprite
	ctrlleft := &ctrlLeftTouchListener{}
	l.ctrlleft.AddTouchListener(ctrlleft)
	ctrlleft.parent = l
}

// ctrlRight
type ctrlRightTouchListener struct {
	parent *LevelScene
}

func (l *ctrlRightTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("[IN] ctrlRight Begin!")

	ctrl := l.parent
	ctrl.buttonState = ctrlRight

	simra.LogDebug("[OUT]")
}

func (l *ctrlRightTouchListener) OnTouchMove(x, y float32) {
	simra.LogDebug("[IN] ctrlRight Move!")

	ctrl := l.parent
	ctrl.buttonState = ctrlRight

	simra.LogDebug("[OUT]")
}

func (l *ctrlRightTouchListener) OnTouchEnd(x, y float32) {
	simra.LogDebug("[IN] ctrlRight End")

	ctrl := l.parent
	ctrl.buttonState = ctrlNop

	simra.LogDebug("[OUT]")
}

func (l *LevelScene) initctrlRight() {
	// set size of ctrlRight
	l.ctrlright.W = float32(120)
	l.ctrlright.H = float32(120)

	// put ctrlRight on left bottom
	l.ctrlright.X = (config.ScreenWidth / 2) + l.ctrlright.W + ctrlMarginBetween
	l.ctrlright.Y = ctrlMarginBottom + l.ctrlright.H

	// rotate arrow to indicate right control
	l.ctrlright.R = -math.Pi / 2

	// add sprite to glpeer
	simra.GetInstance().AddSprite("arrow.png",
		image.Rect(0, 0, int(l.ctrlright.W), int(l.ctrlright.H)),
		&l.ctrlright)

	// add touch listener for sprite
	ctrlright := &ctrlRightTouchListener{}
	l.ctrlright.AddTouchListener(ctrlright)
	ctrlright.parent = l
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (l *LevelScene) Drive() {

	l.updateBoardBlocks()

	switch l.buttonState {
	case ctrlUp:
		l.Game.Rotate()
	case ctrlDown:
		l.Game.MoveDown()
	case ctrlLeft:
		l.Game.MoveLeft()
	case ctrlRight:
		l.Game.MoveRight()
	}
	// stop button repeats
	//l.buttonState = ctrlNop
	if l.Game.GetState() == domain.GameOver {
		simra.GetInstance().SetScene(&GameOverScene{Game: l.Game})
	}
}
