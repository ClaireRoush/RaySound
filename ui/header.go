// actually name is temporaly

package layout

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	music "raySound/music"
)

func renderHeader(musicManager *music.MusicManager) {
	mouseY := float32(rl.GetMouseY())
	resizeZoneHeight := float32(10)
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(navbarHeight), rl.White)
	rl.DrawRectangle(0, int32(navbarHeight), int32(rl.GetScreenWidth()), 10, navbarColor)
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
		if newHeight > 300 && newHeight < 300 { //idk how to design it, so there is no resize anymore :<
			navbarHeight = newHeight
		}
	}

	fontSize := 30
	text := musicManager.GetSongRn().Title
	textWidth := rl.MeasureText(text, int32(fontSize))
	textX := (int32(rl.GetScreenWidth()) - textWidth) / 2
	textY := int32(30 + 150)
	rl.DrawText(text, textX, textY, int32(fontSize), rl.Gray)
	drawImage(musicManager)
	drawButtons(musicManager)
}

func drawImage(musicManager *music.MusicManager) {
	rl.DrawTexture(musicManager.GetSongRn().Cover, int32((rl.GetScreenWidth())/2)-75, 30, rl.White)
}

func drawButtons(musicManager *music.MusicManager) {
	nextSong := rl.Rectangle{
		X:      100,
		Y:      30,
		Width:  40,
		Height: 35,
	}
	rl.DrawRectangleRec(nextSong, rl.Gray)

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), nextSong) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		musicManager.NextSong()
	}
	previousSong := rl.Rectangle{
		X:      30,
		Y:      30,
		Width:  40,
		Height: 35,
	}
	rl.DrawRectangleRec(previousSong, rl.Gray)

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), previousSong) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		musicManager.PreviousSong()
	}
	pauseSong := rl.Rectangle{
		X:      250,
		Y:      30,
		Width:  40,
		Height: 35,
	}
	rl.DrawRectangleRec(pauseSong, rl.Gray)
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), pauseSong) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		musicManager.PauseSong()
	}
}
