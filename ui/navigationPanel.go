package layout

import (
	music "raySound/music"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderNavigationPanel(musicManager music.MusicManager) {
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
