package fight

import (
	"main/src/engine"
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
	// rl "github.com/gen2brain/raylib-go/raylib"
)

type fight int

const (
	PLAYER_TURN  fight = iota
	MONSTER_TURN fight = iota
)

func Fight(player entity.Player, monster entity.Monster, e engine.Engine) {

	rl.ClearBackground(rl.Red)
	rl.DrawTexturePro(e.Fight.StartedFight, rl.NewRectangle(0, 0, 840, 452), rl.NewRectangle(0, 0, 1920, 1080), rl.NewVector2(0, 0), 0, rl.White)

	// Check si le joueur ou le monstre est vaincu. Si c'est le cas, on sort de la boucle
	if player.Health <= 0 {
		player.IsAlive = false
	} else if monster.Health <= 0 {
		player.Inventory = append(player.Inventory, monster.Loot...)
		player.Money += monster.Worth
	}

	player.Attack(&monster)
	monster.Attack(&player)
}
