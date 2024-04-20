package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func CreateTranscoder(outputDir string, filepath string) (*Transcoder, error) {
	dir, err := os.MkdirTemp(outputDir, "transcode-")
	if err != nil {
		return nil, err
	}

	args := []string{}

	args = append(args, []string{
		"-i", filepath,
		"-copyts",
		"-start_at_zero",
		"-c:0", "libsvtav1", // av1
		"-crf", "20",
		"-preset", "9",
		"-force_key_frames", fmt.Sprintf("expr:gte(t,n_forced*%d)", 2),
		"-f", "hls",
		"-start_number", fmt.Sprintf("%d", 0),
		"-hls_time", fmt.Sprintf("%d", 2),
		"-hls_segment_type", "1", // fmp4
		"-hls_segment_filename", "segment%d.m4s",
	}...)

	args = append(args, path.Join(dir, "ffmpeg_manifest.m3u"))

	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Dir = outputDir

	transcode := &Transcoder{
		cmd:       cmd,
		Filepath:  filepath,
		OutputDir: dir,
	}

	return transcode, nil
}
