package goaudiosuite

import (
	"time"

	"gopkg.in/hraban/opus.v2"
)

// PCMToOpus PCMデータをOpusにエンコードします
//
// 引数:
//
//	pcmData: 入力PCMデータのint16スライス
//	sampleRate: PCMデータのサンプリングレート
//	frameSizeMs: 各フレームの時間（ミリ秒単位）
//	opusOutput: エンコードされたOpusデータを送信するためのチャネル
//
// 戻り値:
//
//	エンコード中に発生したエラー（ある場合）
func PCMToOpus(pcmData []int16, sampleRate int, channels int, frameSizeMs int, opusOutput chan []byte) error {
	var opusEncoder *opus.Encoder
	var opusBuffer []byte
	var frame []int16
	var opusSize int
	var end int
	var err error

	frameSize := sampleRate * frameSizeMs / 1000

	if opusEncoder, err = opus.NewEncoder(sampleRate, channels, opus.AppVoIP); err == nil {
		opusBuffer = make([]byte, frameSize*2)
		for i := 0; i < len(pcmData); i += frameSize {
			end = i + frameSize
			if i+frameSize > len(pcmData) {
				break
			}
			frame = pcmData[i:end]
			opusSize, _ = opusEncoder.Encode(frame, opusBuffer)
			opusOutput <- opusBuffer[:opusSize]
			time.Sleep(time.Duration(frameSizeMs) * time.Millisecond) // コードによっては不必要かも
		}
		close(opusOutput) // ここも同様に、コードによっては不必要かも
	} else {
		return err
	}
	return nil
}
