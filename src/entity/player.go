package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {     // definition des varaiable lie au boutton

	Position  rl.Vector2
	Health    int
	Money     int
	Speed     float32
	Inventory []item.Item
	IsGround   bool
	Chute     float32
	Psaut     float32
	Saut      float32

	IsAlive bool

	Sprite rl.Texture2D

	CurrentMonster Monster
}

func (p *Player) Attack(m *Monster) {
	m.Health -= 1
}

func (p *Player) ToString() {
	fmt.Printf(`
    Joueur:
        Vie: %d,
        Argent: %d,
        Inventaire: %+v
    
    \n`, p.Health, p.Money, p.Inventory)
}

func (p *Player) UpdateInventory() {

}

func (p *Player) UpdateMoney() {
	
}
