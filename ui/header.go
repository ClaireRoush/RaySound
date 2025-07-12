package layout

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	music "raySound/music"
)

func RenderAnotherHeader(m *music.MusicManager) {
	screenHeight := int32(rl.GetScreenHeight())
	screenWidth := int32(rl.GetScreenWidth())
	panelBorderTop := screenHeight - 150
	currentPlayed := m.GetSongRn()
	rl.DrawRectangle(0, panelBorderTop, screenWidth, 3, rl.Color{R: 88, G: 91, B: 112, A: 255})
	rl.DrawText(currentPlayed.Title, 35+95, panelBorderTop+15, 25, rl.Color{R: 205, G: 214, B: 244, A: 255})
	rl.DrawText(currentPlayed.Artist, 35+95, panelBorderTop+40, 25, rl.Color{R: 166, G: 173, B: 200, A: 255})
	drawSlider(*m, float32(panelBorderTop))
	drawButtons(m, panelBorderTop)
	drawImage(m, panelBorderTop)
	remain := convertToMinAndSeconds(int32(rl.GetMusicTimePlayed(m.GetSongRn().MusicStream)))
	duration := convertToMinAndSeconds(int32(rl.GetMusicTimeLength(m.GetSongRn().MusicStream)))
	totalTime := remain + " / " + duration
	rl.DrawText(totalTime, 35+95, panelBorderTop+65, 25, rl.Color{R: 166, G: 173, B: 200, A: 255})
}

func convertToMinAndSeconds(seconds int32) string {
	totalMinutes := seconds / 60
	totalSeconds := seconds % 60
	return fmt.Sprintf("%02d:%02d", totalMinutes, totalSeconds)
}

func drawImage(m *music.MusicManager, panelBorderTop int32) {
	rl.DrawTexture(m.GetSongRn().Cover, 35, panelBorderTop+15, rl.White)
}

func drawButtons(musicManager *music.MusicManager, panelBorderTop int32) {
	y := float32(panelBorderTop + 60)
	width := float32(35)
	height := float32(35)
	nextSong := rl.Rectangle{
		X:      (float32(rl.GetScreenWidth())/2 - 20) - 50,
		Y:      y,
		Width:  width,
		Height: height,
	}
	rl.DrawRectangleRec(nextSong, rl.Gray)

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), nextSong) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		musicManager.NextSong()
	}
	previousSong := rl.Rectangle{
		X:      (float32(rl.GetScreenWidth())/2 - 20) + 50,
		Y:      y,
		Width:  width,
		Height: height,
	}
	rl.DrawRectangleRec(previousSong, rl.Gray)

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), previousSong) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		musicManager.PreviousSong()
	}
	pauseSong := rl.Rectangle{
		X:      (float32(rl.GetScreenWidth())/2 - 20) - 0,
		Y:      y,
		Width:  width,
		Height: height,
	}
	rl.DrawRectangleRec(pauseSong, rl.Gray)
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), pauseSong) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		musicManager.PauseSong()
	}
}

func drawSlider(m music.MusicManager, panelBorderTop float32) {
	rectX := float32(35)
	rectMaxX := float32(rl.GetScreenWidth()) - 35
	musicStream := m.GetSongRn().MusicStream
	tMax := rl.GetMusicTimeLength(musicStream)
	tRn := rl.GetMusicTimePlayed(musicStream)
	sliderPos := secondsToPosition(rectX, rectMaxX, tRn, tMax, m)
	sliderLine := rl.Rectangle{
		X:      rectX,
		Y:      panelBorderTop + 110,
		Width:  rectMaxX - rectX,
		Height: 20,
	}
	sliderPointer := rl.Rectangle{
		X:      rectX,
		Y:      panelBorderTop + 110,
		Width:  sliderPos - 35,
		Height: 20,
	}
	rl.DrawRectangleRec(sliderLine, rl.Color{R: 88, G: 91, B: 112, A: 255})
	rl.DrawRectangleRec(sliderPointer, rl.Color{R: 180, G: 190, B: 254, A: 255})
	mouse := rl.GetMousePosition()
	if rl.CheckCollisionPointRec(mouse, sliderLine) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		newTime := positionToSeconds(float32(rl.GetMouseX()), rectX, int32(rectMaxX), tMax)
		rl.SeekMusicStream(m.GetSongRn().MusicStream, newTime)
	}
}

func positionToSeconds(circlePos, rectX float32, rectMaxX int32, tMax float32) float32 {
	return (circlePos - rectX) / (float32(rectMaxX) - rectX) * tMax
}

func secondsToPosition(rectX float32, rectMaxX float32, tRn float32, tMax float32, m music.MusicManager) float32 {
	circlePos := rectX + (tRn/tMax)*(float32(rectMaxX)-rectX)
	fmt.Println(circlePos, rl.GetMusicTimePlayed(m.GetSongRn().MusicStream))
	return circlePos
}
