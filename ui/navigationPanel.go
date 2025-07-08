package layout

import (
	music "raySound/music"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawScrollWindow(m *music.MusicManager) {
	songs := m.GetItems()
	padding := 40
	scrollOffset += rl.GetMouseWheelMove() * 20

	startY := navbarHeight + 15

	contentHeight := float32(len(songs)*padding) + 25
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
		rl.DrawText(song.Title, 40, y, 30, rl.DarkGray)
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && rl.CheckCollisionPointRec(rl.GetMousePosition(), songCard) {
			m.PlayMusic(uint16(i))
		}
	}
	rl.EndScissorMode()
}
