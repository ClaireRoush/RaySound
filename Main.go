package main

import (
	music "raySound/music"
	initMainWindow "raySound/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ScreenWidth = 1280
const ScreenHeight = 920

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(ScreenWidth, ScreenHeight, "meow")
	rl.InitAudioDevice()
	musicManager := music.MusicManager{}
	musicManager.AddItem("./songs/heat abnormal.mp3")
	musicManager.PlayMusic()
	for !rl.WindowShouldClose() {
		musicManager.UpdateStream()
		rl.BeginDrawing()
		initMainWindow.InitBaseLayout(musicManager.GetItem(0).Title)
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
	}
	rl.UnloadMusicStream(musicManager.GetItem(0).MusicStream)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
