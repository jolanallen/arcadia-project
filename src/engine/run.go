package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *Engine) Run() {
	rl.SetTargetFPS(60)

	for engine.IsRunning {

		rl.BeginDrawing()

		switch engine.StateMenu {
		case HOME:
			engine.HomeRendering()
			engine.HomeLogic()

		case SETTINGS:
			engine.SettingsLogic()

		case PLAY:
			switch engine.StateEngine {
			case INGAME:
				engine.InGameRendering()
				engine.InGameLogic()

			case PAUSE:
				engine.PauseRendering()
				engine.PauseLogic()

			case INVENTORY:
				engine.InventoryRendering()
				engine.InventoryLogic()
			case INFIGHT: 
				engine.FightRendering()
				engine.FightLogic() 

			case GAMEOVER:
				engine.GAMEOver()

			case WIN:
				engine.YouWin()

			case LORE:
				engine.LoreRendering()
				engine.LoreLogic()

			}
		}

		rl.EndDrawing()
	}
}
