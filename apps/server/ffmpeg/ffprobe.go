package ffmpeg

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type ProbeContainer struct {
	Streams []ProbeStream `json:"streams"`
	Format  ProbeFormat   `json:"format"`
}

type ProbeStream struct {
	Index          int               `json:"index"`
	CodecName      string            `json:"codec_name"`
	CodecLongName  string            `json:"codec_long_name"`
	CodecType      string            `json:"codec_type"`
	CodecTag       string            `json:"codec_tag"`
	ClosedCaptions int               `json:"closed_captions"`
	PixFmt         string            `json:"pix_fmt"`
	ColorRange     string            `json:"color_range"`
	ColorSpace     string            `json:"color_space"`
	ColorTransfer  string            `json:"color_transfer"`
	ColorPrimaries string            `json:"color_primaries"`
	Profile        string            `json:"profile"`
	Level          int               `json:"level"`
	Channels       int               `json:"channels"`
	ChannelLayout  string            `json:"channel_layout"`
	BitRate        string            `json:"bit_rate"`
	Width          int               `json:"width"`
	Height         int               `json:"height"`
	Extradata      string            `json:"extradata"`
	Tags           map[string]string `json:"tags"`
	Disposition    map[string]int    `json:"disposition"`
	TimeBase       string            `json:"time_base"`
	DurationTs     int               `json:"duration_ts"`
	RFrameRate     string            `json:"r_frame_rate"`
}

type ProbeFormat struct {
	Filename         string            `json:"filename"`
	NBStreams        int               `json:"nb_streams"`
	NBPrograms       int               `json:"nb_programs"`
	FormatName       string            `json:"format_name"`
	FormatLongName   string            `json:"format_long_name"`
	StartTimeSeconds float64           `json:"start_time,string"`
	DurationSeconds  float64           `json:"duration,string"`
	Size             uint64            `json:"size,string"`
	BitRate          uint64            `json:"bit_rate,string"`
	ProbeScore       float64           `json:"probe_score"`
	Tags             map[string]string `json:"tags"`
}

func Probe(path string) (*ProbeContainer, error) {
	fmt.Printf("probing file %s\n", path)

	cmd := exec.Command(
		"ffprobe",
		"-show_data",
		"-show_format",
		"-show_streams", path, "-print_format", "json", "-v", "quiet")
	cmd.Stderr = os.Stderr

	fmt.Printf("starting probe %s with args %s\n", cmd.Path, cmd.Args)

	r, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	cmdOut, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	var v ProbeContainer
	err = json.Unmarshal(cmdOut, &v)
	if err != nil {
		return nil, err
	}

	if len(v.Streams) == 0 {
		return nil, fmt.Errorf("no streams found for %s, is this an actual media file?", path)
	}

	return &v, nil
}
