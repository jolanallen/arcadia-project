package engine

import (
	"main/src/entity"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Rendering() {
	rl.ClearBackground(rl.Blue)
}

func (e *Engine) HomeRendering() {
	rl.ClearBackground(rl.DarkGreen)

	rl.DrawTexturePro(e.background, rl.NewRectangle(0, 0, 1368, 768), rl.NewRectangle(0, 0, 1920, 1080), rl.NewVector2(0, 0), 0, rl.White)

	rl.DrawText("Home Menu", int32(rl.GetScreenWidth())/2-rl.MeasureText("Home Menu", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	if e.StartButton.IsHovered {
		rl.DrawTexturePro(e.StartButton.HoverTexture, rl.NewRectangle(0, 0, 128, 90), rl.NewRectangle(1550, 700, 300, 200), rl.NewVector2(0, 0), 0, rl.White)
	} else {
		rl.DrawTexturePro(e.StartButton.Texture, rl.NewRectangle(0, 0, 128, 90), rl.NewRectangle(1550, 700, 300, 200), rl.NewVector2(0, 0), 0, rl.White)
	}
	rl.DrawText("PLAY", 1620, 775, 60, rl.White)
	rl.DrawTexturePro(e.QuitButton.Texture, rl.NewRectangle(0, 0, 128, 90), rl.NewRectangle(1550, 850, 300, 200), rl.NewVector2(0, 0), 0, rl.White)

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
	rl.DrawTextEx(vrm, "Money:"+strconv.Itoa(e.Player.Money)+" /100", rl.Vector2{X: 5, Y: 50}, 40, 2, rl.Gold) // init.go
	rl.DrawTextEx(vrm, "Heal:"+strconv.Itoa(e.Player.Health)+" / 100", rl.Vector2{X: 5, Y: 5}, 40, 2, rl.Red)
	rl.DrawText("Press [P] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("Press [P] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-490, 20, rl.RayWhite)
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
		rl.DrawTexturePro(
			monster.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(monster.Position.X, monster.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
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
