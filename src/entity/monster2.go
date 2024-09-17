package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Monster2 struct {
	Name     string
	Position rl.Vector2
	Health   int
	Damage   int
	Loot     []item.Item
	Worth    int //valeur en argent quand tué

	IsAlive bool

	Sprite rl.Texture2D
}

func (m *Monster2) Attack(p *Player) {
	p.Health -= 1
}



func (m *Monster2) ToString() {
	fmt.Printf("Je suis un monstre avec %d points de vie\n", m.Health)
}
