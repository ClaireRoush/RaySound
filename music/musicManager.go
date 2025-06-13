package playMusic

import (
	"fmt"
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
	Cover       rl.Texture2D
	Ext         string
}
type MusicManager struct {
	queue []*Music
	Index uint16
}

func (m *MusicManager) AddItem(path string) {
	bytes, ext, err := SaveCoverFromSong(path)
	image := rl.LoadImageFromMemory(ext, bytes, int32(len(bytes)))
	rl.ImageResize(image, 150, 150)
	texture := rl.LoadTextureFromImage(image)
	musicStream := rl.LoadMusicStream(path)
	if err != nil {
		fmt.Print(err)
	}
	item := &Music{
		Title:       filepath.Base(path),
		artist:      "meowArtist",
		Path:        path,
		length:      3200.0,
		isPlaying:   false,
		MusicStream: musicStream,
		Cover:       texture,
		Ext:         ext,
	}
	m.queue = append(m.queue, item)
}

func (m *MusicManager) GetSongRn() *Music {
	return m.queue[m.Index]
}

func (m *MusicManager) GetItems() []*Music {
	return m.queue
}

func (m *MusicManager) GetItem(index uint32) *Music {
	items := m.GetItems()
	if len(items) < int(index) && index < 0 {
		return &Music{}
	}
	return items[index]
}

func (m *MusicManager) PlayMusic(index uint16) {
	songs := m.GetItems()
	rl.PlayMusicStream(songs[index].MusicStream)
	m.Index = index
}

func (m *MusicManager) UpdateStream() {
	rl.UpdateMusicStream(m.GetSongRn().MusicStream)
}
