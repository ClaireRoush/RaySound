package playMusic

import (
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Music struct {
	Title       string
	artist      string
	path        string
	length      float32
	isPlaying   bool
	MusicStream rl.Music
}
type MusicManager struct {
	queue []Music
}

func (m *MusicManager) AddItem(path string) {
	musicStream := rl.LoadMusicStream(path)
	item := Music{
		Title:       filepath.Base(path),
		artist:      "meowArtist",
		path:        path,
		length:      3200.0,
		isPlaying:   false,
		MusicStream: musicStream,
	}
	m.queue = append(m.queue, item)
}

func (m *MusicManager) GetItem(index uint16) Music {
	if int(index) >= len(m.queue) {
		return Music{}
	}
	return m.queue[index]
}

func (m *MusicManager) PlayMusic() {
	rl.PlayMusicStream(m.queue[0].MusicStream)
	m.queue[0].isPlaying = true
}

func (m *MusicManager) UpdateStream() {
	rl.UpdateMusicStream(m.queue[0].MusicStream)
}
