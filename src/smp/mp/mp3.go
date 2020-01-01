package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	status   int
	progress int
}

func (player *MP3Player) Play(source string) {
	fmt.Println("Playing MP3 music ", source)

	player.progress = 0

	for player.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		player.progress += 1
	}

	fmt.Println("Finished playing ", source)
}
