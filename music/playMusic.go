package playMusic

import (
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Music struct {
	Title       string
	artist      string
	Path        string
	length      float32
	isPlaying   bool
	MusicStream rl.Music
}
type MusicManager struct {
	queue []Music
	Index uint16
}

func (m *MusicManager) AddItem(path string) {
	item := Music{
		Title:       filepath.Base(path),
		artist:      "meowArtist",
		Path:        path,
		length:      3200.0,
		isPlaying:   false,
		MusicStream: rl.LoadMusicStream(path),
	}
	m.queue = append(m.queue, item)
}

func (m *MusicManager) GetItem(index uint16) Music {
	if int(index) >= len(m.queue) {
		return Music{}
	}
	return m.queue[index]
}

func (m *MusicManager) GetItems() []Music {
	return m.queue
}

func (m *MusicManager) PlayMusic(index uint16) {
	if int(m.Index) < len(m.queue) && m.queue[m.Index].isPlaying {
		rl.StopMusicStream(m.queue[m.Index].MusicStream)
		m.queue[m.Index].isPlaying = false
	}
	if int(index) >= len(m.queue) {
		println("Invalid index:", index)
		return
	}
	m.queue[index].MusicStream = rl.LoadMusicStream(m.queue[index].Path)
	rl.PlayMusicStream(m.queue[index].MusicStream)
	m.queue[index].isPlaying = true
	m.Index = index
}

func (m *MusicManager) UpdateStream() {
	rl.UpdateMusicStream(m.queue[m.Index].MusicStream)
}
