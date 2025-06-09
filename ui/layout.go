package layout

import (
	music "raySound/music"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	navbarHeight = float32(200)
	resize       = false
	navbarColor  = rl.Gray
	songName     = ""
	droppedFiles []string
	scrollOffset = float32(0)
)

func InitBaseLayout(musicManager *music.MusicManager) {
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

	rl.DrawFPS(0, 0)

	songs := musicManager.GetItems()
	padding := 40
	scrollOffset += rl.GetMouseWheelMove() * 20

	contentHeight := float32(len(songs) * padding)
	visibleHeight := float32(rl.GetScreenHeight()) - navbarHeight
	maxScroll := contentHeight - visibleHeight

	if maxScroll > 0 {
		if scrollOffset < 0 {
			scrollOffset = 0
		}
		if scrollOffset > maxScroll {
			scrollOffset = maxScroll
		}
	} else {
		scrollOffset = 0
	}

	for i, song := range songs {
		startY := navbarHeight + 25
		y := int32(startY) + int32(i*padding) - int32(scrollOffset)

		if y+35 < int32(navbarHeight) || y > int32(rl.GetScreenHeight()) {
			continue
		}

		songCard := rl.Rectangle{
			X:      0,
			Y:      float32(y),
			Width:  float32(rl.GetScreenWidth()),
			Height: 35,
		}

		rl.DrawRectangleRec(songCard, rl.Gray)
		rl.DrawText(song.Title, int32(rl.GetScreenWidth()/2), y, 30, rl.DarkGray)

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && rl.CheckCollisionPointRec(rl.GetMousePosition(), songCard) {
			musicManager.PlayMusic(uint16(i))
		}
	}
}
