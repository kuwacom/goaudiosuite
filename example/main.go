package main

import (
	"encoding/binary"
	"fmt"
	"os"

	"local.packages/goaudiosuite"
)

func main() {
	// PCMデータを生成
	sineWavData := goaudiosuite.GenerateSineWavePCM(48000, 10, 100)
	fmt.Printf("PCM Data: %d\n", len(sineWavData))

	// PCMデータをファイルに保存
	err := savePCMToFile("output.pcm", sineWavData)
	if err != nil {
		fmt.Printf("Error saving PCM data: %v\n", err)
		return
	}
	fmt.Printf("PCM data saved to output.pcm\n")

	testVoiceCh := make(chan []uint8)

	fmt.Printf("Start Encode Opus\n")
	go func() {
		err := goaudiosuite.PCMToOpus(sineWavData, 48000, 1, 20, testVoiceCh)
		if err != nil {
			fmt.Printf("Error encoding Opus: %v\n", err)
		}
	}()

	// Opusデータを結合するためのスライス
	var opusDataAll []uint8
	for opusData := range testVoiceCh {
		fmt.Printf("Opus Data: %d\n", len(opusData))
		opusDataAll = append(opusDataAll, opusData...)
	}

	// Opusデータをファイルに保存
	err = saveOpusToFile("output.opus", opusDataAll)
	if err != nil {
		fmt.Printf("Error saving Opus data: %v\n", err)
		return
	}
	fmt.Printf("Opus data saved to output.opus\n")

	fmt.Printf("Done\n")
}

// PCMデータをファイルに保存する関数
func savePCMToFile(filename string, data []int16) error {
	// PCMデータを[]byteに変換
	byteData := make([]byte, len(data)*2)
	for i, sample := range data {
		binary.LittleEndian.PutUint16(byteData[i*2:], uint16(sample))
	}

	// ファイルを作成し、データを書き込む
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(byteData)
	if err != nil {
		return err
	}

	return nil
}

// Opusデータをファイルに保存する関数
func saveOpusToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
