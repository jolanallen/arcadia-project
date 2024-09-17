package engine

import (
	"fmt"
	"main/src/entity"


	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Rendering() {
	rl.ClearBackground(rl.Blue)
}

func (e *Engine) HomeRendering() {
	rl.ClearBackground(rl.DarkGreen)
	e.BackgroundFrameCount++
	fmt.Println(e.BackgroundFrameCount)

	if e.BackgroundFrameCount%6 == 1 {
		if e.BgSourceX == 9000 {
			e.BgSourceX = 0
		} else {
			e.BgSourceX += 600
		}
	}

	rl.DrawTexturePro(e.Background, rl.NewRectangle(float32(e.BgSourceX), float32(e.BgSourceY), 600, 338), rl.NewRectangle(0, 0, 1920, 1080), rl.NewVector2(0, 0), 0, rl.White)
	rl.DrawText("Home Menu", int32(rl.GetScreenWidth())/2-rl.MeasureText("Home Menu", 50)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("KNIGHT'S QUEST", int32(rl.GetScreenWidth())/2-rl.MeasureText("KNIGHT'S QUEST", 100)/2, int32(rl.GetScreenHeight())/2- 50, 100, rl.Black)

	if e.StartButton.IsHovered {
		rl.DrawTexturePro(e.StartButton.HoverTexture, rl.NewRectangle(0, 0, 128, 90), rl.NewRectangle(1550, 700, 300, 200), rl.NewVector2(0, 0), 0, rl.White)
	} else {rl.DrawTexturePro(e.StartButton.Texture, rl.NewRectangle(0, 0, 128, 90), rl.NewRectangle(1550, 700, 300, 200), rl.NewVector2(0, 0), 0, rl.White)
	}
	rl.DrawText("PLAY", 1620, 775, 60, rl.White)

	if e.QuitButton.IsHovered {
		rl.DrawTexturePro(e.QuitButton.HoverTexture, rl.NewRectangle(0, 0, 128, 90), rl.NewRectangle(1550, 850, 300, 200), rl.NewVector2(0, 0), 0, rl.White)
	} else {
		rl.DrawTexturePro(e.QuitButton.Texture, rl.NewRectangle(0, 0, 128, 90), rl.NewRectangle(1550, 850, 300, 200), rl.NewVector2(0, 0), 0, rl.White)
	}
	rl.DrawText("QUIT", 1620, 925, 60, rl.White)
	rl.DrawText(fmt.Sprint("FPS:", int32(rl.GetFPS())), 1600, 50, 50, rl.Black)
}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)
	rl.BeginMode2D(e.Camera) // On commence le rendu camera
	e.RenderMap()
	e.RenderMonsters()
	e.RenderPlayer()

	rl.EndMode2D() // On finit le rendu camera
	vrm := rl.LoadFont("ressource/font/MedievalSharp/MedievalSharp-Regular.ttf")
	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawTextEx(vrm, "Inventory", rl.Vector2{X: 1500, Y: 1000}, 40, 2, rl.Black)                             // rajouter le tableau en faut faire une boucle le tableau est dans init.go
	rl.DrawText("Press [P] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("Press [P] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-490, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprint("FPS:", int32(rl.GetFPS())), 1600, 50, 50, rl.Black)
	
	rl.DrawTexturePro(
        e.SpriteLife, 
        rl.NewRectangle(0, 0, 435, 100),
        rl.NewRectangle(0, 0, 435, 100),
        rl.NewVector2(0, 0),
        0,
        rl.White)

		rl.DrawTexturePro(
			e.SpriteMoney, 
			rl.NewRectangle(20, 0, 487, 95),
			rl.NewRectangle(0, 120, 487, 95),
			rl.NewVector2(0, 0),
			0,
			rl.White)

			rl.DrawTexturePro(
				e.SpriteInventaire, 
				rl.NewRectangle(100, 0, 510, 451),
				rl.NewRectangle(50, 800, 130, 100),
				rl.NewVector2(0, 0),
				0,
				rl.White)
}

func (e *Engine) PauseRendering() {
	rl.ClearBackground(rl.LightGray)

	rl.DrawText("PAUSE", int32(rl.GetScreenWidth())/2-rl.MeasureText("PAUSE", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Q]/[A] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

	rl.EndDrawing()
}

func (e *Engine) RenderPlayer() {

	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

}

func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		if monster.Name == "Ralouf"{
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
