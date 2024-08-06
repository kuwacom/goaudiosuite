package goaudiosuite

// WAVToOpus WAVデータをOpusにエンコードします
//
// 引数:
//
//	wavData: 入力WAVデータのバイトスライス
//	inputRate: 入力サンプリングレート 0の場合はWAVヘッダーから読み取ります
//	outputRate: 出力サンプリングレート 0の場合は入力サンプリングレートと同じに設定されます
//	frameSizeMs: 各フレームの時間（ミリ秒単位）
//	opusOutput: エンコードされたOpusデータを送信するためのチャネル
//
// 戻り値:
//
//	エンコード中に発生したエラー（ある場合）
func WAVToOpus(wavData []byte, inputRate int, channels int, outputRate int, frameSizeMs int, opusOutput chan []byte) error {
	pcmData, err := WAVToPCM(wavData, inputRate, outputRate)
	if err != nil {
		return err
	}
	return PCMToOpus(pcmData, outputRate, channels, frameSizeMs, opusOutput)
}
