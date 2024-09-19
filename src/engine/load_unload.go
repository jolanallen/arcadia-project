package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Load() {
	// Chargement des textures du personnage
	e.Player.Sprite = rl.LoadTexture("textures/map/tilesets/Legacy-Fantasy - High Forest 2.3/Character/Run/Run.gif")
	e.StartButton.Texture = rl.LoadTexture("textures/img/ButtonGame.png")
	e.StartButton.HoverTexture = rl.LoadTexture("textures/img/ButtonHover.png")
	e.Background = rl.LoadTexture("textures/img/bganim2.png")
	e.QuitButton.Texture = rl.LoadTexture("textures/img/ButtonGame.png")
	e.QuitButton.HoverTexture = rl.LoadTexture("textures/img/ButtonHover.png")
	e.Title = rl.LoadTexture("textures/img/Title.png")
	e.InventoryUI = rl.LoadTexture("textures/entities/inventaire/Screenshot_from_2024-09-19_11-05-01-removebg-preview.png")
	e.GameOver = rl.LoadTexture("textures/img/GameOver.png")
	e.Win = rl.LoadTexture("textures/img/WIn.png")
	e.StartedFight = rl.LoadTexture("textures/img/StartedFight.png")
	e.FondFight = rl.LoadTexture("textures/img/FondFight.jpg")
	e.FontMedieval = rl.LoadFont("ressource/font/MedievalSharp/MedievalSharp-Regular.ttf")
	e.FontFreshman = rl.LoadFont("ressource/font/freshman/Freshman.ttf")
}


func (e *Engine) Unload() {
	// On libère les textures chargées, le joueur, la map, les monstres, etc...
	rl.UnloadTexture(e.Player.Sprite)

	for _, sprite := range e.Sprites {
		rl.UnloadTexture(sprite)
	}

	for _, monster := range e.Monsters {
		rl.UnloadTexture(monster.Sprite)
	}
}
