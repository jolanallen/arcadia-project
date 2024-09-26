package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {    // definition des varaiable lie au boutton

	Texture rl.Texture2D
	HoverTexture rl.Texture2D
	IsHovered bool

}