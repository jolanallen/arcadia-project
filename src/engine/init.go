package engine

import (
	"fmt"
	"main/src/entity"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameEngine struct {
	ScreenWidth  int32
	ScreenHeight int32
	Title        string
}

func (e *Engine) Init() {
	var X int32 = int32(rl.GetScreenWidth())   // création d'une variable X pour stocker la largeur de l'écran
	var Y int32 = int32(rl.GetScreenHeight())  // création d'une variabe Y pour stocker la longeur de l'écran
	e.ScreenHeight = Y 
	e.ScreenWidth = X
	rl.InitWindow(X, Y, "Arcadia") // Initialisation des variables de l'engine
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)

	// Initialisation des composants du jeu
	e.InitEntities()
	e.InitCamera()
	e.InitMusic()
	e.InitMap("textures/map/tilesets/map.json")
	// chargement des texture du jeu 
	e.SpriteLife = rl.LoadTexture("textures/entities/Life/barre_de_vie-removebg-preview.png")
	e.SpriteMoney = rl.LoadTexture("textures/entities/money/Screenshot_from_2024-09-16_12-05-39-removebg-preview.png")
	e.SpriteInventaire = rl.LoadTexture("textures/entities/inventaire/Screenshot_from_2024-09-16_12-20-00-removebg-preview(1).png")
}

func (g *GameEngine) PrintScreenSize() {
	g.ScreenHeight = int32(rl.GetScreenHeight())
	g.ScreenWidth = int32(rl.GetScreenWidth())
	fmt.Println(g.ScreenWidth, "*", g.ScreenHeight)  
}

// ---Init Window--- //
func (g *GameEngine) InitGameEngine(title string) {
	var x = int32(rl.GetScreenWidth()) // récupére la largeur de l'écran et la stock dans x
	var y = int32(rl.GetScreenHeight()) // récupére la hauteur de l'ecran et la stock dans y
	g.ScreenWidth = x
	g.ScreenHeight = y
	g.Title = title
	rl.InitWindow(g.ScreenWidth, g.ScreenHeight, g.Title)
	rl.SetTargetFPS(60)              // initialisation des fps maximum du jeu 
	rl.ToggleFullscreen()
}

func (e *Engine) InitEntities() {            // initialisaion du personnage et des ennemis
	e.Player = entity.Player{
		Position:  rl.Vector2{X: 130, Y: 210},
		Health:    100,
		Money:     0,
		Speed:     1,
		Inventory: []item.Item{},
		Chute:      0.5,
		Psaut:      0,
		Saut:       0,

		IsAlive: true,

		Sprite: e.Player.Sprite,
	}

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "bee guard",
		Position: rl.Vector2{X: 1100, Y: 360},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/map/tilesets/Legacy-Fantasy - High Forest 2.3/Mob/Small Bee/Fly/Fly-Sheet.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Ralouf",
		Position: rl.Vector2{X: 1300, Y: 435},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/map/tilesets/Legacy-Fantasy - High Forest 2.3/Mob/Boar/Walk/Walk-Base-Sheet.png"),
	})
	e.Player.Money = 1
}

func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D( // Camera vide, à changer dans chaque logique de scène
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		2.0,
	)
}

func (e *Engine) InitMusic() {
	rl.InitAudioDevice()

	e.Music = rl.LoadMusicStream("sounds/music/alexander-nakarada-chase(chosic.com).mp3")

	rl.PlayMusicStream(e.Music)
}
