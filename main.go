package main

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/oto/v2"
)

func main() {
	// Open the sound file
	file, err := os.Open("sound.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the sound file into memory
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Create an audio context
	context, ready, err := oto.NewContext(44100, 2, 8192)
	if err != nil {
		log.Fatal(err)
	}
	<-ready

	// Create a new player using the data as an io.Reader
	player := context.NewPlayer(bytes.NewReader(data))
	defer player.Close()

	// Start playback and wait until it finishes
	player.Play()

	// Block the main goroutine until the sound finishes playing
	select {}
}
