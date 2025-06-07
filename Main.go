package main

import (
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"

	gui "github.com/gen2brain/raylib-go/raygui"
)

const screenWidth = 1280
const screenHeight = 920

var navbarHeight = float32(200)
var resize = false
var navbarColor = rl.Gray
var songName = ""
var droppedFiles []string
var music = rl.LoadMusicStream("")

func drawLayout(mouseY float32, droppedFiles []string) {
	resizeZoneHeight := float32(10)
	rl.DrawRectangle(0, 0, screenWidth, int32(navbarHeight), rl.White)
	rl.DrawRectangle(0, int32(navbarHeight), screenWidth, 10, navbarColor)
	onResizeZone := mouseY > navbarHeight-resizeZoneHeight && mouseY < navbarHeight+resizeZoneHeight
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && onResizeZone {
		navbarColor = rl.DarkGray
		resize = true
	}
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		resize = false
		navbarColor = rl.Gray
	}
	if resize {
		newHeight := mouseY
		if newHeight > 100 && newHeight < 300 {
			navbarHeight = newHeight
		}
	}
	fontSize := 45
	var text = "Nothing here :<"
	if len(droppedFiles) > 0 {
		text = "Listening rn: " + filepath.Base(droppedFiles[0])
	}
	textWidth := rl.MeasureText(text, int32(fontSize))
	textX := (screenWidth - textWidth) / 2
	textY := int32((navbarHeight - float32(fontSize)) / 2)
	rl.DrawText(text, textX, textY, int32(fontSize), rl.Gray)
	rl.DrawFPS(0, 0)
	gui.Button(rl.NewRectangle((screenWidth/2)-25, (navbarHeight)-55, 50, 50), "start")
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "meow")
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.UpdateMusicStream(music)
		drawLayout(float32(rl.GetMouseY()), droppedFiles)
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
	}
	rl.UnloadMusicStream(music)
	rl.CloseWindow()
}
