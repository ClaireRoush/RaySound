package main

import (
	"fmt"
	"os"
	music "raySound/music"
	ui "raySound/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ScreenWidth = 1280
const ScreenHeight = 920

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(ScreenWidth, ScreenHeight, "meow")
	rl.InitAudioDevice()
	musicManager := music.MusicManager{}
	musicManager.AddItem("./songs/edited_heat_abnormal.mp3")
	img := rl.LoadImageFromMemory(".png", []byte(musicManager.GetItem(musicManager.Index).Cover), int32(len(musicManager.GetItem(musicManager.Index).Cover)))

	os.WriteFile("cover_test.png", []byte(musicManager.GetItem(musicManager.Index).Cover), 0644)
	fmt.Print([]byte(musicManager.GetItem(musicManager.Index).Cover))
	texture := rl.LoadTextureFromImage(img)
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		musicManager.UpdateStream()
		rl.BeginDrawing()
		rl.DrawTexture(texture, 120, 120, rl.White)
		ui.InitBaseLayout(&musicManager)
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
	}
	rl.UnloadMusicStream(musicManager.GetItem(0).MusicStream)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
