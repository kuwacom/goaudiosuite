package goaudiosuite

import (
	"encoding/binary"
	"errors"
)

// WAVToPCM WAVデータをPCMデータに変換します。
//
// 引数:
//
//	wavData: 入力WAVデータのバイトスライス
//	inputRate: 入力サンプリングレート 0の場合はWAVヘッダーから読み取ります
//	outputRate: 出力サンプリングレート 0の場合は入力サンプリングレートと同じに設定されます
//
// 戻り値:
//
//	PCMデータのint16スライスとエラー（発生した場合）
func WAVToPCM(wavData []byte, inputRate int, outputRate int) ([]int16, error) {
	const wavHeaderSize = 44
	if len(wavData) < wavHeaderSize {
		return nil, errors.New("wav data too short")
	}

	if inputRate == 0 {
		inputRate = int(binary.LittleEndian.Uint32(wavData[24:28]))
	}

	if outputRate == 0 {
		outputRate = inputRate
	}

	if (len(wavData)-wavHeaderSize)%2 == 1 {
		return nil, errors.New("illegal wav data: payload must be encoded in byte pairs")
	}

	numPCMSamples := (len(wavData) - wavHeaderSize) / 2
	PCMsamples := make([]int16, numPCMSamples)
	for i := 0; i < numPCMSamples; i++ {
		PCMsamples[i] += int16(wavData[wavHeaderSize+i*2])
		PCMsamples[i] += int16(wavData[wavHeaderSize+i*2+1]) << 8
	}

	if inputRate != outputRate {
		PCMsamples = ResamplePCM(PCMsamples, inputRate, outputRate)
	}

	return PCMsamples, nil
}
