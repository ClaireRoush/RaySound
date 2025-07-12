package main

import (
	music "raySound/music"
	ui "raySound/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ScreenWidth = 1050
const ScreenHeight = 600

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "meow")
	rl.InitAudioDevice()
	musicManager := music.MusicManager{}
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/edited_secret_weapon.flac")
	musicManager.AddItem("./songs/heat_abnormal_with_another_image.mp3")
	for !rl.WindowShouldClose() {
		musicManager.UpdateStream()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{R: 30, G: 30, B: 46, A: 255})
		ui.InitBaseLayout(&musicManager)
		rl.EndDrawing()
	}
	rl.CloseAudioDevice()
	music.UnloadMusicStreams(&musicManager)
	rl.CloseWindow()
}
