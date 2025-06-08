package main

import (
	initMainWindow "raySound/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ScreenWidth = 1280
const ScreenHeight = 920

var DroppedFiles []string

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(ScreenWidth, ScreenHeight, "meow")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		initMainWindow.InitBaseLayout(DroppedFiles)
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
