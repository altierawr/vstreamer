package ffmpeg

type TranscodeProfile struct {
	VideoPreset VideoPreset
	AudioPreset AudioPreset
}

type VideoPreset struct {
	Codec      string
	FfmpegArgs string
}

type AudioPreset struct {
	Codec      string
	FfmpegArgs string
}

var TranscodingVideoCodecs = []VideoPreset{
	{
		Codec:      "av1",
		FfmpegArgs: "",
	},
	{
		Codec:      "vp9",
		FfmpegArgs: "",
	},
}

var TranscodingAudioCodecs = []AudioPreset{
	{
		Codec:      "flac",
		FfmpegArgs: "",
	},
}
