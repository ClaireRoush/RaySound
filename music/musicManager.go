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
	Length      float32
	TimePlayed  float32
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
	rl.UnloadImage(image)
	musicStream := rl.LoadMusicStream(path)
	if err != nil {
		fmt.Print(err)
	}
	item := &Music{
		Title:       filepath.Base(path),
		artist:      "meowArtist",
		Path:        path,
		Length:      0,
		TimePlayed:  0,
		isPlaying:   true,
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
	rl.PlayMusicStream(songs[index].MusicStream)
	m.Index = index
}

func (m *MusicManager) UpdateStream() {
	rl.UpdateMusicStream(m.GetSongRn().MusicStream)
	time := rl.GetMusicTimePlayed(m.GetSongRn().MusicStream)
	fmt.Println(time, rl.GetMusicTimeLength(m.GetSongRn().MusicStream))
}

func (m *MusicManager) NextSong() {
	items := len(m.GetItems())
	if m.Index+1 >= uint16(len(m.GetItems())) {
		m.Index = 0
	} else {
		m.Index++
	}
	m.PlayMusic(m.Index)
	fmt.Println(items, m.Index)
}

func (m *MusicManager) PreviousSong() {
	items := len(m.GetItems())
	if m.Index <= 0 {
		m.Index = uint16(items - 1)
	} else {
		m.Index--
	}
	fmt.Println(items, m.Index)
	m.PlayMusic(m.Index)
}

func (m *MusicManager) PauseSong() {
	song := m.GetSongRn()
	if song.isPlaying {
		rl.PauseMusicStream(song.MusicStream)
		song.isPlaying = false
	} else {
		rl.ResumeMusicStream(song.MusicStream)
		song.isPlaying = true
	}
}
