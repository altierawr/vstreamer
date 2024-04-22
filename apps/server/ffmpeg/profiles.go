package ffmpeg

type TranscodeProfile struct {
	VideoPreset VideoPreset
	AudioPreset AudioPreset
}

type VideoPreset struct {
	Codec      string
	Mime       string
	FfmpegArgs []string
}

type AudioPreset struct {
	Codec      string
	FfmpegArgs string
}

var Av1HDRPreset = VideoPreset{
	Codec:      "av1",
	Mime:       "av01.0.19M.10",
	FfmpegArgs: []string{},
}

var Av1SDRPreset = VideoPreset{
	Codec:      "av1",
	Mime:       "av01.0.19M.10",
	FfmpegArgs: []string{},
}

var Vp9HDRPreset = VideoPreset{
	Codec:      "vp9",
	Mime:       "vp9",
	FfmpegArgs: []string{},
}

var Vp9SDRPreset = VideoPreset{
	Codec:      "vp9",
	Mime:       "vp9",
	FfmpegArgs: []string{},
}

var TranscodingAudioCodecs = []AudioPreset{
	{
		Codec:      "flac",
		FfmpegArgs: "",
	},
}
