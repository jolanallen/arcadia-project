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
	INFIGHT engine = iota 
)

type Engine struct {
	Background           rl.Texture2D
	BgSourceX            int
	BgSourceY            int
	BackgroundFrameCount int
	QuitButton  entity.Button
	StartButton entity.Button
	Fight Fight

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

type Fight struct {

	StartedFight rl.Texture2D
	Background rl.Texture2D
	CurrentMonster entity.Monster

}