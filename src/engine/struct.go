package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type menu int

const (
	HOME     menu = iota
	SETTINGS menu = iota
	PLAY     menu = iota
)

type engine int

const (
	INGAME   engine = iota
	PAUSE    engine = iota
	GAMEOVER engine = iota
	WIN      engine = iota
)

type Engine struct {
	Title                rl.Texture2D
	Background           rl.Texture2D
	BgSourceX            int
	BgSourceY            int
	BackgroundFrameCount int
	QuitButton           entity.Button
	StartButton          entity.Button
	ScreenWidth          int32
	ScreenHeight         int32
	ColisionListe        []rl.Rectangle
	GameOver             rl.Texture2D
	WIn                  rl.Texture2D

	Player   entity.Player
	Monsters []entity.Monster

	Music       rl.Music
	MusicVolume float32

	Sprites          map[string]rl.Texture2D
	SpriteLife       rl.Texture2D
	SpriteMoney      rl.Texture2D
	SpriteInventaire rl.Texture2D
	

	Camera rl.Camera2D

	MapJSON MapJSON

	IsRunning   bool
	StateMenu   menu
	StateEngine engine
}
