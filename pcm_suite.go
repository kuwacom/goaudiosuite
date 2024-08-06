package goaudiosuite

import "math"

const (
	amplitude = 32767 // 16bitPCMの最大振幅
)

// PCMResample PCMデータを指定されたサンプリングレートにリサンプルします
// 線形補間を利用して、新しいサンプル間のデータを計算します
//
// 16bit深度のPCM音声データのみに対応
//
// 引数:
//   originalData: 入力PCMデータのint16スライス
//   originalRate: 入力データのサンプリングレート
//   targetRate: 出力データのサンプリングレート
//
// 戻り値:
//   リサンプルされたPCMデータのint16スライス
func ResamplePCM(originalData []int16, originalRate, targetRate int) []int16 {
	rateRatio := float64(targetRate) / float64(originalRate)
	// 新しいサンプル数用のスライスを作成
	resampledData := make([]int16, int(float64(len(originalData))*rateRatio))

	for i := 0; i < len(resampledData); i++ {
		// 新しいサンプルの位置が元のデータのどの位置に対応するか
		srcIndex := float64(i) / rateRatio
		// 元のデータ内の隣接する2つのサンプルのインデックス
		leftIndex := int(srcIndex)
		rightIndex := leftIndex + 1

		if rightIndex >= len(originalData) {
			resampledData[i] = originalData[leftIndex]
		} else {
			leftValue := float64(originalData[leftIndex])
			rightValue := float64(originalData[rightIndex])
			interpolation := srcIndex - float64(leftIndex)
			resampledData[i] = int16(leftValue*(1-interpolation) + rightValue*interpolation)
		}
	}

	return resampledData
}

// MonoToStereoPCM PCMデータをモノラルからステレオ(1チャンネルから2チャンネル)へ変換します
//
// 16bit深度のPCM音声データのみに対応
//
// 引数:
//   monoData: 入力PCMデータのint16スライス
//
// 戻り値:
//   ステレオ化された2チャンネルのPCMデータのint16スライス
func MonoToStereoPCM(monoData []int16) []int16 {
	// モノラルデータの長さを取得し、ステレオデータの長さを計算
	monoLen := len(monoData)
	stereoLen := monoLen * 2
	stereoData := make([]int16, stereoLen)

	// モノラルデータの各サンプルをステレオデータにコピー
	for i := 0; i < monoLen; i++ {
		sample := monoData[i]
		stereoData[i*2] = sample   // 左チャンネル
		stereoData[i*2+1] = sample // 右チャンネル
	}

	return stereoData
}

// MonoToStereoPCM テスト用の正弦波PCMを16bitで生成
// 引数:
//   sampleRate: 出力PCMのサンプリングレート
//   durationSec: 音の長さ
//   frequency: 周波数
//
// 戻り値:
//   ステレオ化された2チャンネルのPCMデータのint16スライス
func GenerateSineWavePCM(sampleRate float64, durationSec float64, frequency float64) []int16 {
	// PCMデータのサンプル数を計算
	numSamples := int(sampleRate * durationSec)
	// スライスを作成してPCMデータを格納
	data := make([]int16, numSamples)

	// 正弦波のサンプルを計算
	for i := 0; i < numSamples; i++ {
		// サンプルの時刻
		t := float64(i) / sampleRate
		// 正弦波のサンプル値を計算
		value := amplitude * math.Sin(2*math.Pi*frequency*t)
		// 16ビットPCMとしてデータに格納
		data[i] = int16(value)
	}

	return data
}
