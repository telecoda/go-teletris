package scene

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"

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
	playerSprites []*simra.Sprite
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

// const (
// 	ctrlNop = iota
// 	ctrlUp
// 	ctrlDown
// 	ctrlLeft
// 	ctrlRight
// )

// const (
// 	ctrlMarginLeft      = 10
// 	ctrlMarginBottom    = 100
// 	ctrlMarginBetween   = 10
// 	buttonMarginRight   = 20
// 	buttonMarginBottom  = 20
// 	buttonMarginBetween = 10
// )

// Initialize initializes LevelScene scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (l *LevelScene) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// add global touch listener to catch touch end event
	//simra.GetInstance().AddTouchListener(l)

	// TODO: when goes to next scene, remove global touch listener
	// simra.GetInstance().RemoveTouchListener(LevelScene)

	// initialize sprites
	l.initSprites()
	l.buttonReplaced = false
	simra.LogDebug("[OUT]")
}

func (l *LevelScene) initSprites() {
	l.initBackground()
	l.initBlockTextures()
	l.initBackgroundImage()
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

	// add left touch listener for background
	touchListener := &touchListener{}
	touchListener.parent = l
	l.background.AddTouchListener(touchListener)

}

func (l *LevelScene) initBlockTextures() {

	l.blockTextures = make(map[domain.BlockColour]*sprite.SubTex, 0)
	l.blockImages = make(map[domain.BlockColour]*image.RGBA, 0)
	rect := image.Rect(0, 0, domain.BlockPixels, domain.BlockPixels)

	for i, name := range domain.SpriteNames {

		blockImage, _, err := io.LoadImage(name)
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

func (l *LevelScene) initBackgroundImage() {

	// load image from file
	sourceImage, _, err := io.LoadImage("space_background.png")
	if err != nil {
		panic(fmt.Sprintf("Error loading image: %s\n", err))
	} else {
		gridImage := DrawGrid(sourceImage, 20, 20)
		// draw grey blocks on background
		targetImage := l.drawBlocks(gridImage)
		rect := image.Rect(0, 0, targetImage.Bounds().Dx(), targetImage.Bounds().Dy())
		tex := peer.GetGLPeer().LoadTextureFromImage(targetImage, rect)

		// update background image
		peer.GetSpriteContainer().ReplaceTexture(&l.background.Sprite, tex)

		// save image for reuse
		l.backgroundImage = targetImage
	}
}

func (l *LevelScene) redrawBackgroundImage() {

	l.initBackgroundImage()
	rect := image.Rect(0, 0, l.backgroundImage.Bounds().Dx(), l.backgroundImage.Bounds().Dy())
	//clearImage := image.NewNRGBA(rect)
	targetImage := l.drawBlocks(l.backgroundImage)
	//rect := image.Rect(0, 0, targetImage.Bounds().Dx(), targetImage.Bounds().Dy())
	tex := peer.GetGLPeer().LoadTextureFromImage(targetImage, rect)

	// update background image
	peer.GetSpriteContainer().ReplaceTexture(&l.background.Sprite, tex)

}

func (l *LevelScene) drawBlocks(sourceImage image.Image) image.Image {

	gameBlocks := l.Game.GetBlocks()
	blocks := *gameBlocks
	boardWidth := len(blocks)

	point := image.Point{X: 0, Y: 0}
	bounds := sourceImage.Bounds()
	maxY := bounds.Dy()
	targetImage := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(targetImage, bounds, sourceImage, bounds.Min, draw.Src)

	for x := 0; x < boardWidth; x++ {
		boardHeight := len(blocks[x])
		for y := 0; y < boardHeight; y++ {

			block := blocks[x][y]
			if block.Colour == domain.Empty {
				continue
			}
			blockImage := l.blockImages[block.Colour]
			xCoord := (x * domain.BlockPixels) + domain.BoardOffsetX
			yCoord := maxY - ((y + 2) * domain.BlockPixels) + domain.BoardOffsetY - domain.BlockPixels/2 - 4

			rect := image.Rect(xCoord, yCoord, xCoord+domain.BlockPixels, yCoord+domain.BlockPixels)
			draw.Draw(targetImage, rect, blockImage, point, draw.Src)
		}
	}

	return targetImage
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

func (l *LevelScene) initPlayerSprites() {

	// playerBlocks are the only moving sprites
	game := l.Game
	player := game.Player
	playerBlocks := player.GetShapeBlocks()
	l.playerSprites = make([]*simra.Sprite, len(playerBlocks))
	for i, _ := range playerBlocks {
		playerSprite := new(simra.Sprite)

		playerBlockX := playerBlocks[i].X + player.X
		playerBlockY := playerBlocks[i].Y + player.Y
		playerSprite.W = float32(domain.BlockPixels)
		playerSprite.H = float32(domain.BlockPixels)

		// put center of screen
		playerSprite.X = float32(domain.BlockPixels*playerBlockX + domain.BlockPixels/2 + domain.BoardOffsetX)
		playerSprite.Y = float32(domain.BlockPixels*playerBlockY + domain.BlockPixels/2 + domain.BoardOffsetY)

		// lookup blockImage for sprite colour
		blockImage := domain.SpriteNames[playerBlocks[i].Colour]
		simra.GetInstance().AddSprite(blockImage,
			image.Rect(0, 0, int(domain.BlockPixels), int(domain.BlockPixels)),
			playerSprite)

		l.playerSprites[i] = playerSprite
	}

}

func (l *LevelScene) removePlayerSprites() {
	for i, _ := range l.playerSprites {
		simra.GetInstance().RemoveSprite(l.playerSprites[i])
	}
	l.playerSprites = nil
}

func (l *LevelScene) updatePlayerSprites() {
	game := l.Game

	// init Player sprites if they do not exist
	if l.playerSprites == nil {
		l.initPlayerSprites()
	}

	player := game.Player
	playerBlocks := player.GetShapeBlocks()

	for i, _ := range playerBlocks {
		playerSprite := l.playerSprites[i]

		playerBlockX := playerBlocks[i].X + player.X
		playerBlockY := playerBlocks[i].Y + player.Y
		playerSprite.W = float32(domain.BlockPixels)
		playerSprite.H = float32(domain.BlockPixels)

		// put center of screen
		playerSprite.X = float32(domain.BlockPixels*playerBlockX + domain.BlockPixels/2 + domain.BoardOffsetX)
		playerSprite.Y = float32(domain.BlockPixels*playerBlockY + domain.BlockPixels/2 + domain.BoardOffsetY)

	}

}

// touchListener
type touchListener struct {
	parent                       *LevelScene
	touching                     bool
	touchBeginX, touchBeginY     float32
	touchCurrentX, touchCurrentY float32
	touchEndX, touchEndY         float32
}

func (t *touchListener) OnTouchBegin(x, y float32) {
	t.touchBeginX = x
	t.touchBeginY = y
	t.touchCurrentX = x
	t.touchCurrentY = y
	t.touching = true
}

func (t *touchListener) OnTouchMove(x, y float32) {

	if !t.touching {
		// update values
		t.touchBeginX = x
		t.touchBeginY = y
		t.touchCurrentX = x
		t.touchCurrentY = y
		t.touching = true
	}

	t.touchCurrentX = x
	t.touchCurrentY = y

	xMovement := t.touchBeginX - t.touchCurrentX
	yMovement := t.touchBeginY - t.touchCurrentY

	// check if touch is near edge of block
	moveTolerance := float32(domain.BlockPixels)

	if xMovement >= moveTolerance {
		t.parent.Game.MoveLeft()
		// reset touch
		t.touching = false
		return
	}

	if xMovement <= -moveTolerance {
		t.parent.Game.MoveRight()
		// reset touch
		t.touching = false
		return
	}

	if yMovement >= moveTolerance {
		t.parent.Game.MoveDown()
		// reset touch
		t.touching = false
		return
	}

	if yMovement <= -moveTolerance {
		t.parent.Game.Rotate()
		// reset touch
		t.touching = false
		return
	}
}

func (t *touchListener) OnTouchEnd(x, y float32) {
	t.touching = false
	t.touchEndX = x
	t.touchEndY = y
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (l *LevelScene) Drive() {

	if l.Game.IsBoardDirty() {
		// redraw board
		l.removePlayerSprites()
		l.redrawBackgroundImage()
		l.Game.CleanBoard()
	}
	l.updatePlayerSprites()

	if l.Game.GetState() == domain.GameOver {
		simra.GetInstance().SetScene(&GameOverScene{Game: l.Game})
	}
}
