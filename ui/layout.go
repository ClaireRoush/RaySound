package layout

import (
	music "raySound/music"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	navbarHeight = float32(300)
	resize       = false
	navbarColor  = rl.Gray
	songName     = ""
	droppedFiles []string
	scrollOffset = float32(0)
)

func InitBaseLayout(musicManager *music.MusicManager) {
	// DrawScrollWindow(musicManager)
	// renderHeader(musicManager)
	RenderAnotherHeader(musicManager)
}
