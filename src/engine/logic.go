package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) HomeLogic() {

	//Musique

	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
	if rl.GetMousePosition().X > 1550 && rl.GetMousePosition().X < 1850 && rl.GetMousePosition().Y > 700 && rl.GetMousePosition().Y < 900 {
		e.StartButton.IsHovered = true
		if rl.IsMouseButtonDown(0) {
			e.StateMenu = PLAY
			e.StateEngine = INGAME
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
		e.StateEngine = INGAME
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
	// Mouvement
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed
	}
	// Saut du personnage

	const jump float32 = 12.0
	const poid float32 = 1
	const sol float32 = 409.0 // hauteur sol

	if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyUp) {
		if !e.Player.Jumping {
			e.Player.Jumping = true
			e.Player.Chute = -jump // saute avec une vitesse de -12 sur l'axe y
		}
	}

	// gestion de la chute
	if e.Player.Jumping {
		e.Player.Position.Y += e.Player.Chute
		e.Player.Chute += poid //  le poids pour faire redescendre le personnage

		if e.Player.Position.Y >= sol { //// si la postioon du personnage sur l'axe des y est supérieur ou égal a celle du sol
			e.Player.Position.Y = sol //// Rester au sol
			e.Player.Jumping = false  // permet que le personnage ne suate pas a l'infini
		}
	}

	if rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift) { // sprint du perso
		e.Player.Speed = 3
	} else {
		e.Player.Speed = 1
	}

	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X +490, Y: e.Player.Position.Y + 20} // Bouger la caméra
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}                   // Bouger la

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	e.CheckCollisions()

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-07-Simon_s-In-There-Somewhere.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) CheckCollisions() {

	e.MonsterCollisions()
	e.ZoneCollisions()

}
func (e *Engine) ZoneCollisions() {
	// Ajout des colisions sur les zone dite interdit de la map !!!
}

func (e *Engine) MonsterCollisions() {

	for _, monster := range e.Monsters {
		if monster.Position.X > e.Player.Position.X-50 &&
			monster.Position.X < e.Player.Position.X+50 &&
			monster.Position.Y > e.Player.Position.Y-50 &&
			monster.Position.Y < e.Player.Position.Y+50 {

			if monster.Name == "bee guard" {
				e.NormalTalk(monster, "Press E for FIGHT!!")
				if rl.IsKeyPressed(rl.KeyE) {
				} else {
					///: ....
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
					//lancer un combat en attendant juste dire combat refuse
				}
			}
		} else {
			///.....
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
