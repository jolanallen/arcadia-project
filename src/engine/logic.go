package engine

import (
	
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

	if rl.IsKeyPressed(rl.KeyEnter) { // si la touche entrer est pressé
		e.StateMenu = PLAY           // le statut du menu pas en mode PLAY
		e.StateEngine = LORE        // on passe ensuite le statut du jeu en mode lore (histoire du jeu)
		e.Timer = rl.GetTime()      // on lance le timer et on stock dans e.Timer
		rl.StopMusicStream(e.Music)   // on coupe la music 
	}
	if rl.IsKeyPressed(rl.KeyQ) {         // si la touche Q est présser 
		e.IsRunning = false              // le jeu s'arrête permet de quitter le jeu 
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
func (e *Engine) LoreLogic() {             // fonction qui gére la logic d'execution du lore
	if rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
		e.InitEntities()
	}
	if e.Timer+10 <= rl.GetTime() {      // permet de définir la durée aprés la quelle le lore se ferme et le jeu se lance 
		e.StateEngine = INGAME          
	}
}

func (e *Engine) InGameLogic() {   // fonction qui permet la gestion de la logique dans le jeux 
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
			e.Player.IsGround = false                           // on défini le player comme n'étant plus au sol   
		}
			if e.BigJump {
				e.Player.Psaut = -25
				e.Player.IsGround = false                           // on défini le player comme n'étant plus au sol   
			}
	}
	// gestion du saut 
	if e.Player.Psaut < 0 {                    // tant que psaut est inferieur a zero sachant que on démarre a -18
		e.Player.Position.Y += e.Player.Psaut // on ajoute a la position du player en Y psaut
		e.Player.Psaut += 1                  //psaut est incrémenter de 1 a chaque fois 
	}                                       // on a donc -18, -17 -16 -15 -14 -13 ..... etc 
    // arrête du saut 
	if e.Player.IsGround {              // si le player est au sol
		e.Player.Psaut = 0             // au remet a zero la puissance du saut, cela evite d'avoir un saut de plus en  plus grand 
		e.Player.Chute = 0.5          // on remet a 0.5 la force qui fait chuter le player cela evite qui tombe de plus en plus rapidement a chaque chute
	}
	

	if rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift) {    // si la touche shift gauche ou shift droite est pressé
		 if e.SupSpeed {
			e.Player.Speed = 100
		 } else {
			e.Player.Speed = 50
		 }
		e.Player.Speed = 3                                               // alors la variable speed qui correspond a la vitesse du player est defini a 3
	} else {                                                             // sinon 
		e.Player.Speed = 1                                               // la variable speed reste a 1 
	}


	if e.Player.Position.Y >= 800 {                      // si la position du player sur l'axe des Y est supérieur ou égale a 800 
		e.StateEngine = GAMEOVER                         // le statut du jeux passe a GAMEOVER 
	}
	if e.Player.Position.X <= 990 && // si la postion en x du player est comprise entre [840 ; 990] ET 
	e.Player.Position.X >= 840 &&    
	e.Player.Position.Y >= 400 {   // si la position en Y est superieurou egale a 400 
		e.Player.IsGround = true      // le joueur est au sol
		rl.WaitTime(3)                // on attend 3 sec
		e.StateEngine = GAMEOVER     // et le statut du programme passe a gameover et execute donc la fonction liée a gameover 
	}
	if e.Player.Position.X >= 1456 {   // si la position du player en x est superieur ou égale  a 1456
		rl.WaitTime(2)
		e.StateEngine = WIN             // le statut du programe passe a WIN et execute donc le function liée a win 

	}

	// Inventory

	if rl.IsKeyPressed(rl.KeyTab) {   // si la touche TAB est pressée 
		e.StateEngine = INVENTORY     // alors le statut du programme passe a INVENTORY se qui execute la fonction qui permet d'afficher l'inventaire 
	}

	// Camera
	var ScreenWidth float32
	var ScreenHeight float32
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X - 400, Y: e.Player.Position.Y - 270} // Bouger la caméra
	e.Camera.Offset = rl.Vector2{X: ScreenWidth, Y: ScreenHeight}                            // Bouger la

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	e.CheckCollisions()

	if e.Player.Health < 1 {
		e.StateEngine = GAMEOVER
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
}
func (e *Engine) FightLogic() {

}

// Gère les collisions avec les monstres dans le jeu.
func (e *Engine) MonsterCollisions() {
	// Parcourt tous les monstres présents dans le jeu.
	for _, monster := range e.Monsters {
		// Vérifie si la position du monstre est à proximité de celle du joueur (dans un rayon donné).
		if monster.Position.X > e.Player.Position.X-50 &&
			monster.Position.X < e.Player.Position.X+150 &&
			monster.Position.Y > e.Player.Position.Y-150 &&
			monster.Position.Y < e.Player.Position.Y+150 {

			// Si le monstre est nommé "bee guard".
			if monster.Name == "bee guard" {
				// Affiche un message indiquant que le joueur peut combattre.
				e.NormalTalk(monster, "Press E for FIGHT!!")
				// Vérifie si la touche E a été pressée.
				if rl.IsKeyPressed(rl.KeyE) {
					// Change l'état du moteur de jeu pour indiquer que le joueur est en combat.
					e.StateEngine = INFIGHT
					// Définit le monstre actuel du joueur.
					e.Player.CurrentMonster = monster
					// Affiche un message indiquant que le combat commence.
					e.NormalTalk(monster, "le combat commence!!")
				}
			}
		} 
	}

	// Vérifie une autre condition de collision avec les monstres.
	for _, Monster2 := range e.Monsters {
		// Vérifie si la position du deuxième monstre est à proximité du joueur (dans un rayon plus petit).
		if Monster2.Position.X > e.Player.Position.X-50 &&
			Monster2.Position.X < e.Player.Position.X+50 &&
			Monster2.Position.Y > e.Player.Position.Y-50 &&
			Monster2.Position.Y < e.Player.Position.Y+50 {

			// Si le monstre est nommé "Ralouf".
			if Monster2.Name == "Ralouf" {
				// Affiche un message indiquant que le joueur peut combattre.
				e.NormalTalk(Monster2, "Press E for FIGHT!!")

				// Vérifie si la touche E a été pressée.
				if rl.IsKeyPressed(rl.KeyE) {
					// Change l'état du moteur de jeu pour indiquer que le joueur est en combat.
					e.StateEngine = INFIGHT
					// Définit le monstre actuel du joueur.
					e.Player.CurrentMonster = Monster2
					// Affiche un message indiquant que le combat commence.
					e.NormalTalk(Monster2, "Le combat commence !")
				}
			}
		}
	}
}

// Affiche un dialogue entre le joueur et un monstre.
func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	// Appelle la fonction pour rendre le dialogue à l'écran.
	e.RenderDialog(m, sentence)
}

// Gère la logique de pause du jeu.
func (e *Engine) PauseLogic() {
	// Vérifie si le joueur a appuyé sur la touche Échap ou P pour revenir au jeu.
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
	}
	// Vérifie si le joueur a appuyé sur A ou Q pour revenir au menu principal.
	if rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyQ) {
		e.StateMenu = HOME
		// Arrête la musique de fond.
		rl.StopMusicStream(e.Music)
	}

	// Met à jour le flux de musique.
	rl.UpdateMusicStream(e.Music)
}

// Gère la logique de fin de jeu en cas de défaite.
func (e *Engine) GAMEOver() {
	// Réinitialise l'état du moteur de jeu pour revenir à l'état de jeu normal.
	e.StateEngine = INGAME
	// Réinitialise toutes les entités du jeu.
	e.InitEntities()
}

// Gère la logique de fin de jeu en cas de victoire.
func (e *Engine) YouWin() {
	// Réinitialise l'état du moteur de jeu pour revenir à l'état de jeu normal.
	e.StateEngine = INGAME
	// Réinitialise toutes les entités du jeu.
	e.InitEntities()
}

