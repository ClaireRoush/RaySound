// actually name is temporaly

package layout

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	music "raySound/music"
)

func renderHeader(musicManager music.MusicManager) {
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
		if newHeight > 100 && newHeight < 300 {
			navbarHeight = newHeight
		}
	}

	fontSize := 45
	currentSong := musicManager.GetItem(musicManager.Index).Title
	text := "Playing rn: " + currentSong
	textWidth := rl.MeasureText(text, int32(fontSize))
	textX := (int32(rl.GetScreenWidth()) - textWidth) / 2
	textY := int32((navbarHeight - float32(fontSize)) / 2)
	rl.DrawText(text, textX, textY, int32(fontSize), rl.Gray)

}
