package main

import (
	music "raySound/music"
	ui "raySound/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ScreenWidth = 800
const ScreenHeight = 600

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "meow")
	rl.InitAudioDevice()
	musicManager := music.MusicManager{}
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/heat_abnormal_with_another_image.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/heat_abnormal_with_another_image.mp3")
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	musicManager.AddItem("./songs/heat_abnormal_with_another_image.mp3")
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		musicManager.UpdateStream()
		ui.InitBaseLayout(&musicManager)
		rl.EndDrawing()
	}
	rl.CloseAudioDevice()
	music.UnloadMusicStreams(&musicManager)
	rl.CloseWindow()
}
