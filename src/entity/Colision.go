package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Colision struct{
	Name     string
	Position rl.Vector2
	IsContact bool
	

}