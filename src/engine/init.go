package engine

import (
    "fmt"
    "main/src/entity"
    "main/src/item"

    rl "github.com/gen2brain/raylib-go/raylib"
)

const (
    ScreenWidth  = 1800
    ScreenHeight = 1080
)

type GameEngine struct {
    ScreenWidth  int32
    ScreenHeight int32
    Title        string
}

func (e *Engine) Init() {
    rl.InitWindow(ScreenWidth, ScreenHeight, "Arcadia")

    // Initialisation des variables de l'engine
    e.IsRunning = true
    e.Sprites = make(map[string]rl.Texture2D)

    // Initialisation des composants du jeu
    e.InitEntities()
    e.InitCamera()
    e.InitMusic()
    e.InitMap("textures/map/tilesets/map.json")

    e.SpriteLife = rl.LoadTexture("textures/entities/Life/barre_de_vie-removebg-preview.png")
    e.SpriteMoney = rl.LoadTexture("textures/entities/money/Screenshot_from_2024-09-16_12-05-39-removebg-preview.png")
    e.SpriteInventaire = rl.LoadTexture("textures/entities/inventaire/Screenshot_from_2024-09-16_12-20-00-removebg-preview(1).png")
}

func (g *GameEngine) PrintScreenSize() {
    fmt.Println(g.ScreenWidth, "*", g.ScreenHeight)
}

// ---Init Window--- //
func (g *GameEngine) InitGameEngine(x int32, y int32, title string) {
    g.ScreenWidth = x
    g.ScreenHeight = y
    g.Title = title
    rl.InitWindow(g.ScreenWidth, g.ScreenHeight, g.Title)
    rl.SetTargetFPS(60)
    rl.ToggleFullscreen()
}

func (e *Engine) InitEntities() {
    e.Player = entity.Player{
        Position:  rl.Vector2{X: 300, Y: 300},
        Health:    100,
        Money:     1000,
        Speed:     1,
        Inventory: []item.Item{},

        IsAlive: true,

        Sprite: e.Player.Sprite,
    }

    e.Monsters = append(e.Monsters, entity.Monster{
        Name:     "claude",
        Position: rl.Vector2{X: 400, Y: 320},
        Health:   20,
        Damage:   5,
        Loot:     []item.Item{},
        Worth:    12,

        IsAlive: true,
        Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
    })

    e.Player.Money = 0
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

    e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")

    rl.PlayMusicStream(e.Music)
}
