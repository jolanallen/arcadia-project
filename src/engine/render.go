package engine

import (
	// "crypto/x509"
	"fmt"
	"main/src/entity"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Rendering() {
	rl.ClearBackground(rl.Blue)
}
func (e *Engine) HomeRendering() {
	rl.ClearBackground(rl.DarkGreen) // Définit la couleur de fond pour l'écran d'accueil.
	e.BackgroundFrameCount++          // Incrémente le compteur d'images de fond.

	fmt.Println(e.BackgroundFrameCount) // Affiche la valeur actuelle de BackgroundFrameCount (probablement utilisé pour des tests).

	// Change l'arrière-plan toutes les 6 frames.
	if e.BackgroundFrameCount%6 == 1 {
		if e.BgSourceX == 9000 {
			e.BgSourceX = 0 // Remet à zéro si le fond atteint un certain point.
		} else {
			e.BgSourceX += 600 // Déplace le fond vers la droite.
		}
	}

	// Affiche l'image d'arrière-plan avec un scrolling.
	rl.DrawTexturePro(e.Background, rl.NewRectangle(float32(e.BgSourceX), float32(e.BgSourceY), 600, 338), rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())), rl.NewVector2(0, 0), 0, rl.White)

	// Affiche le titre du jeu à l'écran.
	rl.DrawTexturePro(e.Title, rl.NewRectangle(0, 0, 1472, 832), rl.NewRectangle(450, 300, 1000, 600), rl.NewVector2(0, 0), 0, rl.White)

	// Vérifie si le bouton "Start" est survolé, et affiche l'état correspondant.
	if e.StartButton.IsHovered {
		rl.DrawTexturePro(e.StartButton.HoverTexture, rl.NewRectangle(0, 0, 128, 90), rl.NewRectangle(1550, 700, 300, 200), rl.NewVector2(0, 0), 0, rl.White)
	} else {
		rl.DrawTexturePro(e.StartButton.Texture, rl.NewRectangle(0, 0, 128, 90), rl.NewRectangle(1550, 700, 300, 200), rl.NewVector2(0, 0), 0, rl.White)
	}

	// Affiche le texte "PLAY" pour le bouton de démarrage.
	rl.DrawText("PLAY", 1620, 775, 60, rl.White)

	// Vérifie si le bouton "Quit" est survolé, et affiche l'état correspondant.
	if e.QuitButton.IsHovered {
		rl.DrawTexturePro(e.QuitButton.HoverTexture, rl.NewRectangle(0, 0, 128, 90), rl.NewRectangle(1550, 850, 300, 200), rl.NewVector2(0, 0), 0, rl.White)
	} else {
		rl.DrawTexturePro(e.QuitButton.Texture, rl.NewRectangle(0, 0, 128, 90), rl.NewRectangle(1550, 850, 300, 200), rl.NewVector2(0, 0), 0, rl.White)
	}

	// Affiche le texte "QUIT" pour le bouton de sortie.
	rl.DrawText("QUIT", 1620, 925, 60, rl.White)

	// Affiche le nombre d'images par seconde (FPS).
	rl.DrawText(fmt.Sprint("FPS:", int32(rl.GetFPS())), 1700, 30, 25, rl.Red)
}
func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)         // Efface l'arrière-plan avec une couleur grise.
	rl.BeginMode2D(e.Camera)            // Démarre le rendu de la caméra pour centrer la scène autour du personnage.
	e.RenderMap()                       // Rendu de la carte.
	e.RenderMonsters()                  // Rendu des monstres présents sur la carte.
	e.RenderPlayer()                    // Rendu du personnage joueur.

	rl.EndMode2D()                      // Fin du rendu de la caméra.
	
	// Affiche des informations supplémentaires qui ne sont pas affectées par la caméra.
	rl.DrawTextEx(e.FontMedieval, "press tab for show Inventory", rl.Vector2{X: 820, Y: 10}, 30, 2, rl.White)
	rl.DrawTextEx(e.FontMedieval, "Money:"+strconv.Itoa(e.Player.Money)+" /100", rl.Vector2{X: 5, Y: 100}, 40, 2, rl.Gold)
	rl.DrawText("Press [P] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("Press [P] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-490, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprint("fps:", int32(rl.GetFPS())), 1700, 30, 40, rl.Red)

	// Rendu des éléments comme les barres de vie, inventaire, etc.
	rl.DrawTexturePro(e.SpriteLife, rl.NewRectangle(0, 0, 435, 100), rl.NewRectangle(0, 0, 435, 100), rl.NewVector2(0, 0), 0, rl.White)
	rl.DrawTexturePro(e.SpriteInventaire, rl.NewRectangle(100, 0, 510, 451), rl.NewRectangle(50, 800, 130, 100), rl.NewVector2(0, 0), 0, rl.White)

	// Vérifie la position du joueur et affiche "Game Over" ou "Win" selon certaines conditions.
	if e.Player.Position.Y >= 420 {
		rl.DrawTexturePro(e.GameOver, rl.NewRectangle(0, 0, 1280, 1280), rl.NewRectangle(300, -50, 1500, 1500), rl.NewVector2(0, 0), 0, rl.White)
	}
	if e.Player.Position.X >= 1450 {
		rl.DrawTexturePro(e.Win, rl.NewRectangle(0, 0, 300, 300), rl.NewRectangle(120, 100, 1500, 1500), rl.NewVector2(0, 0), 0, rl.White)
	}
}

func (e *Engine) PauseRendering() {
	rl.DrawTexturePro(e.Background, rl.NewRectangle(float32(e.BgSourceX), float32(e.BgSourceY), 600, 338), rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())), rl.NewVector2(0, 0), 0, rl.White)
	rl.DrawText("PAUSE", int32(rl.GetScreenWidth())/2-rl.MeasureText("PAUSE", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.White)
	rl.DrawText("[P] or [Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.White)
	rl.DrawText("[Q]/[A] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.White)

	rl.EndDrawing()
}

func (e *Engine) InventoryRendering() {

	rl.DrawTexturePro(
		e.InventoryUI,
		rl.NewRectangle(2, 0, 564, 441),
		rl.NewRectangle(650, 350, 566, 441),
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)
}

func (e *Engine) RenderPlayer() {

	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 100, 90),
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 90, 90),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

}

func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		if monster.Name == "Ralouf" {
			rl.DrawTexturePro(
				monster.Sprite,
				rl.NewRectangle(0, 0, 40, 35),
				rl.NewRectangle(monster.Position.X, monster.Position.Y, 40, 35),
				rl.Vector2{X: 0, Y: 0},
				0,
				rl.White,
			)
		}
		if monster.Name == "bee guard" {
			rl.DrawTexturePro(
				monster.Sprite,
				rl.NewRectangle(0, 0, 60, 70),
				rl.NewRectangle(monster.Position.X, monster.Position.Y, 60, 70),
				rl.Vector2{X: 0, Y: 0},
				0,
				rl.White,
			)
		}
	}
}

func (e *Engine) FightRendering() {
	e.StartedFightCountFrames++

	rl.DrawTexturePro(e.FondFight, rl.NewRectangle(0, 0, 640, 460), rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight() + 50)), rl.NewVector2(0, 0), 0, rl.White)
	rl.ClearBackground(rl.White)
	if e.StartedFightCountFrames < 200 {
		rl.DrawTexturePro(e.StartedFight, rl.NewRectangle(0, 0, 450, 450), rl.NewRectangle(160, 0, 1590, 900), rl.NewVector2(0, 0), 0, rl.White)
	}


	rl.DrawTexturePro(e.Player.CurrentMonster.Sprite, rl.NewRectangle(0, 0, 50, 39), rl.NewRectangle(700, 800, 200, 200), rl.NewVector2(0, 0), 0, rl.White)
	
	rl.DrawTexturePro(e.Player.Sprite, rl.NewRectangle(0, 0, 100, 100), rl.NewRectangle(100, 900, 311, 311), rl.NewVector2(0, 0), 0, rl.White)

}

func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X),
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}

func (e *Engine) LoreRendering() {
	rl.ClearBackground(rl.Black)
	rl.DrawText("Press [P] for Pass", int32(rl.GetScreenWidth())/2-rl.MeasureText("Press [P] for Pass", 20)/2, int32(rl.GetScreenHeight())/2-550, 15, rl.RayWhite)

	if e.Timer+2 <= rl.GetTime() {
		rl.DrawTextPro(e.FontFreshman, "Knight's Quest", rl.NewVector2(650, 100), rl.NewVector2(0, 0), 0, 60, 10, rl.DarkGreen)
	}

	if e.Timer+4 <= rl.GetTime() {
		rl.DrawTextPro(e.FontFreshman, "Knight's Quest", rl.NewVector2(650, 100), rl.NewVector2(0, 0), 0, 60, 10, rl.DarkGreen)
		rl.DrawTextPro(e.FontFreshman, "In the village of Oakwood, a legendary Oakwood Acorn has gone missing. \n Dark forces in the nearby forest are suspected.", rl.NewVector2(50, 200), rl.NewVector2(0, 0), 0, 30, 5, rl.DarkGreen)
	}
	
	if e.Timer+6 <= rl.GetTime() {
		rl.DrawTextPro(e.FontFreshman, "Knight's Quest", rl.NewVector2(650, 100), rl.NewVector2(0, 0), 0, 60, 10, rl.DarkGreen)
		rl.DrawTextPro(e.FontFreshman, "In the village of Oakwood, a legendary Oakwood Acorn has gone missing. \n Dark forces in the nearby forest are suspected.", rl.NewVector2(50, 200), rl.NewVector2(0, 0), 0, 30, 5, rl.DarkGreen)
		rl.DrawTextPro(e.FontFreshman, "track down the thieves, defeat the porc cerfs and bee swarms guarding the forest,\n and reclaim the treasured artifact. \n Brave knights have protected Oakwood for generations. \n Now, it's your turn. Explore ancient ruins, hidden clearings, \n and treacherous paths. The fate of Oakwood hangs in the balance. \n Will you emerge victorious and restore peace to the village? ", rl.NewVector2(50, 350), rl.NewVector2(0, 0), 0, 30, 5, rl.DarkGreen)
	}
	
	if e.Timer+10 <= rl.GetTime() {
		e.StateEngine = INGAME
	}

	rl.EndDrawing()
}
