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

	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.StateEngine = LORE
		e.Timer = rl.GetTime()
		rl.StopMusicStream(e.Music)
	}
	if rl.IsKeyPressed(rl.KeyQ) {
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
func (e *Engine) LoreLogic() {
	if rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
		e.InitEntities()
	}
	if e.Timer+10 <= rl.GetTime() {
		e.StateEngine = INGAME
	}
}

func (e *Engine) InGameLogic() {
	// colisions a droit de la map sur l'axe des x 
	if e.Player.Position.X >= 90 {
		if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
			e.Player.Position.X -= e.Player.Speed
		}
	}
	// collisons a gauche de la map sur l'axe des x
	if e.Player.Position.X <= 1500 {
		if rl.IsKeyDown(rl.KeyE) || rl.IsKeyDown(rl.KeyRight) {
			e.Player.Position.X += e.Player.Speed
		}
	}
	e.ZoneCollisions()


	                                            // gravité appliquer au si le player n'est pas au sol
	if !e.Player.IsGround {                    // si le personnage n'est pas au sol donc en l'air
		e.Player.Position.Y += e.Player.Chute // on ajoute la valeur de la variable chute
		e.Player.Chute += 0.7                // on ajoute a la variable chute +0.7 tanque que le joueur est en l'air
	}
	if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyUp) { // si la touche espaces ou fleche du haut est pressé
		if e.Player.IsGround {                                    // et si le joueur est au sol
			e.Player.Psaut = -18                                 // on defini la variable Psaut "puissance saut" a -18 
		}
	}
	// gestion du saut 
	if e.Player.Psaut < 0 {                    // tant que psaut est inferieur a zero sachant que on démarre a -18
		e.Player.Position.Y += e.Player.Psaut // on ajoute a la position du player en Y psaut
		e.Player.Psaut += 1                //psaut est incrémenter de 1 a chaque fois 
	}                                     // on a donc -18, -17 -16 -15 -14 -13 ..... etc 
    // arrête du saut 
	if e.Player.IsGround {              // si le player est au sol
		e.Player.Psaut = 0             // au remet a zero la puissance du saut, cela evite d'avoir un saut de plus en  plus grand 
		e.Player.Chute = 0.5          // on remet a 0.5 la force qui fait chuter le player cela evite qui tombe de plus en plus rapidement a chaque chute
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

			if Monster2.Name == "Ralouf" {
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
	e.StateEngine = INGAME
	e.InitEntities()

}
func (e *Engine) YouWin() {
	e.StateEngine = INGAME
	e.InitEntities()
}
