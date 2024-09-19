package engine

import (
	"fmt"
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) HomeLogic() {

	//Musique

	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/alexander-nakarada-chase(chosic.com).mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
	if rl.GetMousePosition().X > 1550 && rl.GetMousePosition().X < 1850 && rl.GetMousePosition().Y > 700 && rl.GetMousePosition().Y < 900 {
		e.StartButton.IsHovered = true
		if rl.IsMouseButtonDown(0) {
			e.StateMenu = PLAY
			e.StateEngine = LORE
			e.Timer = rl.GetTime()
			rl.StopMusicStream(e.Music)
		}
	}
	if !(rl.GetMousePosition().X > 1550 && rl.GetMousePosition().X < 1850 && rl.GetMousePosition().Y > 700 && rl.GetMousePosition().Y < 900) {
		e.StartButton.IsHovered = false
	}

	if rl.GetMousePosition().X > 1550 && rl.GetMousePosition().X < 1850 && rl.GetMousePosition().Y > 850 && rl.GetMousePosition().Y < 1050 {
		e.QuitButton.IsHovered = true
		if rl.IsMouseButtonDown(0) {
			e.IsRunning = false
		}
	}
	if !(rl.GetMousePosition().X > 1550 && rl.GetMousePosition().X < 1850 && rl.GetMousePosition().Y > 850 && rl.GetMousePosition().Y < 1050) {
		e.QuitButton.IsHovered = false
	}

	//Menus

	if rl.IsMouseButtonDown(0) {
		e.StateMenu = PLAY
		e.StateEngine = LORE
		e.Timer = rl.GetTime()
		rl.StopMusicStream(e.Music)

	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
}

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyB) {
		e.StateMenu = HOME
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) InGameLogic() {
	if e.Player.Position.X >= 90 {
		if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
			e.Player.Position.X -= e.Player.Speed
		}
	}
	if e.Player.Position.X <= 1500 {
		if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
			e.Player.Position.X += e.Player.Speed
		}
	}
	e.ZoneCollisions()
	// Saut du personnage

	if !e.Player.IsGround {
		rl.WaitTime(0.001)
		e.Player.Position.Y += 4
	}

	if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyUp) {
		if e.Player.IsGround {
			
			e.Player.Position.Y -= 110
			e.Player.IsGround = false
		}
	}

	if rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift) { // sprint du perso
		e.Player.Speed = 3
	} else {
		e.Player.Speed = 1
	}
	if e.Player.Position.Y >= 800 {
		e.StateEngine = GAMEOVER
	}
	if e.Player.Position.X <= 990 && e.Player.Position.X >= 840 && e.Player.Position.Y >= 400 {
		e.Player.IsGround = true
		rl.WaitTime(3)
		e.StateEngine = GAMEOVER
	}
	if e.Player.Position.X >= 1456 {
		rl.WaitTime(2)
		e.StateEngine = WIN

	}

	// Inventory

	if rl.IsKeyPressed(rl.KeyTab) {
		e.StateEngine = INVENTORY
	}

	// Camera
	var ScreenWidth float32
	var ScreenHeight float32
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 490, Y: e.Player.Position.Y + 20} // Bouger la caméra
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}                   // Bouger la
	e.ScreenHeight = int32(ScreenHeight)
	e.ScreenWidth = int32(ScreenWidth)
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X - 400, Y: e.Player.Position.Y - 270} // Bouger la caméra
	e.Camera.Offset = rl.Vector2{X: ScreenWidth, Y: ScreenHeight}                            // Bouger la

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	e.CheckCollisions()

	if e.Player.Health < 1 {
		e.StateEngine = INGAME
	}

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/alexander-nakarada-chase(chosic.com).mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) InventoryLogic() {
	if rl.IsKeyPressed(rl.KeyTab) {
		e.StateEngine = INGAME
	}
}

func (e *Engine) CheckCollisions() {
	e.MonsterCollisions()

}
func (e *Engine) ZoneCollisions() {
	e.Player.IsGround = false
	for _, Colision := range e.ColisionListe {
		if Colision.X > e.Player.Position.X-20 &&
			Colision.X < e.Player.Position.X+20 &&
			Colision.Y > e.Player.Position.Y-39 &&
			Colision.Y < e.Player.Position.Y+39 {
			e.Player.IsGround = true
		}
	}

	// Ajout des colisions sur les zone dite interdit de la map !!!
}

func (e *Engine) FightLogic() {

}

func (e *Engine) MonsterCollisions() {

	for _, monster := range e.Monsters {
		if monster.Position.X > e.Player.Position.X-50 &&
			monster.Position.X < e.Player.Position.X+150 &&
			monster.Position.Y > e.Player.Position.Y-150 &&
			monster.Position.Y < e.Player.Position.Y+150 {

			if monster.Name == "bee guard" {
				e.NormalTalk(monster, "Press E for FIGHT!!")
				if rl.IsKeyPressed(rl.KeyE) {
					e.StateEngine = INFIGHT
					e.Player.CurrentMonster = monster
					fmt.Println("Le combat commence !")
				}

			}
		} else {
			////.....
		}

	}

	for _, Monster2 := range e.Monsters {
		if Monster2.Position.X > e.Player.Position.X-50 &&
			Monster2.Position.X < e.Player.Position.X+50 &&
			Monster2.Position.Y > e.Player.Position.Y-50 &&
			Monster2.Position.Y < e.Player.Position.Y+50 {

			if Monster2.Name == "patate" {
				e.NormalTalk(Monster2, "Press E for FIGHT!!")
				if rl.IsKeyPressed(rl.KeyE) {
					e.StateEngine = INFIGHT
					e.Player.CurrentMonster = Monster2
					fmt.Println("Le combat commence !")
				}
			}
		}
	}
}

func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}

func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyQ) {
		e.StateMenu = HOME
		rl.StopMusicStream(e.Music)
	}

	//Musique
	rl.UpdateMusicStream(e.Music)
}
func (e *Engine) GAMEOver() {
	e.StateMenu = HOME
	e.InitEntities()

}
func (e *Engine) YouWin() {
	e.StateMenu = HOME
	e.InitEntities()
}

func (e *Engine) LoreLogic() {
	if e.Timer+2 <= rl.GetTime() {
		e.loreText = "Knight's quest \n\n\n"
	}
	if e.Timer+4 <= rl.GetTime() {
		e.loreText += "In the village of Oakwood, a legendary Oakwood Acorn has gone missing. \n Dark forces in the nearby forest are suspected. \n\n"
	}
	if e.Timer+6 <= rl.GetTime() {
		e.loreText += "track down the thieves, defeat the porc cerfs and bee swarms guarding the forest,\n and reclaim the treasured artifact. \n Brave knights have protected Oakwood for generations. \n Now, it's your turn. Explore ancient ruins, hidden clearings, \n and treacherous paths. The fate of Oakwood hangs in the balance. \n Will you emerge victorious and restore peace to the village? "
	}
	if e.Timer+10 <= rl.GetTime() {
		e.StateEngine = INGAME
	}
}
