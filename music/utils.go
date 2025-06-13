package playMusic

import (
	"fmt"
	"os"

	"github.com/dhowden/tag"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func SaveCoverFromSong(path string) ([]byte, string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, "", fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	tags, err := tag.ReadFrom(file)
	if err != nil {
		return nil, "", fmt.Errorf("read tags: %w", err)
	}

	cover := tags.Picture()
	if cover == nil {
		return nil, "", fmt.Errorf("err %s", cover.MIMEType)
	}

	var ext string
	switch cover.MIMEType {
	case "image/jpeg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	default:
		return nil, "", fmt.Errorf("err %s", cover.MIMEType)
	}

	return cover.Data, ext, nil
}

func UnloadMusicStreams(m *MusicManager) {
	for index := range len(m.GetItems()) {
		rl.UnloadMusicStream(m.GetItem(uint32(index)).MusicStream)
	}
}
