package mp

import "fmt"
import "time"

type MP3Player struct {
	stat     int
	profress int
}

func (p *MP3Player) Play(source string) {
	fmt.Println("Playing MP3 music:", source)
	p.profress = 0
	for p.profress < 100 {
		time.Sleep(100 * time.Millisecond) //   正    fmt.Print(".")
		p.profress += 10
	}
	fmt.Println("\nFinished playing", source)
}
