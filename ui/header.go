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
	textY := int32(30 + 155)
	rl.DrawText(text, textX, textY, int32(fontSize), rl.Gray)
	drawImage(musicManager)
	drawButtons(musicManager)
}

func drawImage(musicManager *music.MusicManager) {
	rl.DrawTexture(musicManager.GetSongRn().Cover, int32((rl.GetScreenWidth())/2)-75, 30, rl.White)
}

func drawButtons(musicManager *music.MusicManager) {
	nextSong := rl.Rectangle{
		X:      (float32(rl.GetScreenWidth())/2 - 20) - 50,
		Y:      220,
		Width:  40,
		Height: 35,
	}
	rl.DrawRectangleRec(nextSong, rl.Gray)

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), nextSong) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		musicManager.NextSong()
	}
	previousSong := rl.Rectangle{
		X:      (float32(rl.GetScreenWidth())/2 - 20) + 50,
		Y:      220,
		Width:  40,
		Height: 35,
	}
	rl.DrawRectangleRec(previousSong, rl.Gray)

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), previousSong) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		musicManager.PreviousSong()
	}
	pauseSong := rl.Rectangle{
		X:      (float32(rl.GetScreenWidth())/2 - 20) - 0,
		Y:      220,
		Width:  40,
		Height: 35,
	}
	rl.DrawRectangleRec(pauseSong, rl.Gray)
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), pauseSong) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		musicManager.PauseSong()
	}
	drawSlider(*musicManager)
}

func drawSlider(m music.MusicManager) {
	rectX := float32(120)
	rectMaxX := float32(rl.GetScreenWidth()) - 120
	tMax := rl.GetMusicTimeLength(m.GetSongRn().MusicStream)
	tRn := rl.GetMusicTimePlayed(m.GetSongRn().MusicStream)
	secondsToPosition(rectX, rectMaxX, tRn, tMax)
	circlePos := secondsToPosition(rectX, rectMaxX, tRn, tMax)
	sliderLine := rl.Rectangle{
		X:      rectX,
		Y:      navbarHeight - 30,
		Width:  rectMaxX - rectX,
		Height: 5,
	}
	sliderPointer := rl.Rectangle{
		X:      circlePos - 5,
		Y:      navbarHeight - 35,
		Width:  10,
		Height: 15,
	}
	rl.DrawRectangleRec(sliderLine, rl.Gray)
	rl.DrawRectangleRec(sliderPointer, rl.Gray)
	mouse := rl.GetMousePosition()
	if rl.CheckCollisionPointRec(mouse, sliderLine) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		newTime := positionToSeconds(float32(rl.GetMouseX()), rectX, int32(rectMaxX), tMax)
		rl.SeekMusicStream(m.GetSongRn().MusicStream, newTime)
	}
}

func positionToSeconds(circlePos, rectX float32, rectMaxX int32, tMax float32) float32 {
	return (circlePos - rectX) / (float32(rectMaxX) - rectX) * tMax
}

func secondsToPosition(rectX float32, rectMaxX float32, tRn float32, tMax float32) float32 {
	circlePos := rectX + (tRn/tMax)*(float32(rectMaxX)-rectX)
	return circlePos
}
