package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Load() {
	// Chargement des textures du personnage
	e.Player.Sprite = rl.LoadTexture("textures/entities/soldier/Soldier-Idle.png")
	e.StartButton.Texture = rl.LoadTexture("textures/img/ButtonGame.png")
	e.StartButton.HoverTexture = rl.LoadTexture("textures/img/ButtonHover.png")
	e.Background = rl.LoadTexture("textures/img/bganim2.png")
	e.QuitButton.Texture = rl.LoadTexture("textures/img/ButtonGame.png")
	e.QuitButton.HoverTexture = rl.LoadTexture("textures/img/ButtonHover.png")
	e.Title = rl.LoadTexture("textures/img/Title.png")

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
