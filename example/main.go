package main

import (
	"fmt"

	"local.packages/goaudiosuite"
)

func main() {
	fmt.Printf("Hello")
	sineWavData := goaudiosuite.GenerateSineWavePCM(48000, 10, 100)
	fmt.Printf("PCM Data: %d \n", len(sineWavData))
	testVoiceCh := make(chan []uint8)

	fmt.Printf("Start Encode Opus")
	go func() {
		err := goaudiosuite.PCMToOpus(sineWavData, 48000, 1, 20, testVoiceCh)
		if err != nil {
			fmt.Printf("Error encoding Opus: %v\n", err)
			close(testVoiceCh)
		}
	}()

	for opusData := range testVoiceCh {
		fmt.Printf("Opus Data: %d \n", len(opusData))
	}
	fmt.Printf("Done\n")
}
