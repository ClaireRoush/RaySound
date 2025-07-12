package playMusic

import (
	"fmt"
	"os"

	"github.com/dhowden/tag"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Music struct {
	Title       string
	Artist      string
	Path        string
	Length      float32
	TimePlayed  float32
	IsPlaying   bool
	MusicStream rl.Music
	Cover       rl.Texture2D
	Ext         string
}
type MusicManager struct {
	queue    []*Music
	Index    uint16
	IsPaused bool
}

func (m *MusicManager) AddItem(path string) {
	bytes, ext, err := SaveCoverFromSong(path)
	image := rl.LoadImageFromMemory(ext, bytes, int32(len(bytes)))
	rl.ImageResize(image, 83, 83)
	texture := rl.LoadTextureFromImage(image)
	rl.UnloadImage(image)
	musicStream := rl.LoadMusicStream(path)
	osFile, err := os.Open(path)
	tag, err := tag.ReadFrom(osFile)
	if err != nil {
		fmt.Print(err)
	}
	item := &Music{
		Title:       tag.Title(),
		Artist:      tag.Artist(),
		Path:        path,
		Length:      0,
		TimePlayed:  0,
		IsPlaying:   false,
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
	if len(items) < int(index) && index <= 0 {
		return &Music{}
	}
	return items[index]
}

func (m *MusicManager) PlayMusic(index uint16) {
	songs := m.GetItems()
	rl.StopMusicStream(songs[m.Index].MusicStream)
	m.GetSongRn().IsPlaying = false
	rl.PlayMusicStream(songs[index].MusicStream)
	m.Index = index
	m.GetSongRn().IsPlaying = true
}

func (m *MusicManager) UpdateStream() {
	rl.UpdateMusicStream(m.GetSongRn().MusicStream)
}

func (m *MusicManager) NextSong() {
	if m.Index+1 >= uint16(len(m.GetItems())) {
		m.Index = 0
	} else {
		m.Index++
	}
	m.PlayMusic(m.Index)
}

func (m *MusicManager) PreviousSong() {
	if m.Index <= 0 {
		m.Index = uint16(len(m.GetItems()) - 1)
	} else {
		m.Index--
	}
	m.PlayMusic(m.Index)
}

func (m *MusicManager) PauseSong() {
	song := m.GetSongRn()
	if !m.IsPaused {
		rl.PauseMusicStream(song.MusicStream)
		m.IsPaused = true
	} else {
		rl.ResumeMusicStream(song.MusicStream)
		m.IsPaused = false
	}
}
