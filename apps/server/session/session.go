package session

import (
	"context"
	"fmt"
	"slices"

	"github.com/altierawr/vstreamer/ent"
	"github.com/altierawr/vstreamer/ent/video"
	"github.com/altierawr/vstreamer/ffmpeg"
)

var MediaStreams map[int][]struct {
	stream     *ent.Stream
	transcoder *ffmpeg.Transcoder
}

var videoHeights = []int{4320, 2160, 1440, 1080, 720, 480, 360, 240}

func CreatePlaySession(ctx context.Context, client *ent.Client, videoID int) (*ent.PlaySession, error) {
	playSession, err := client.PlaySession.Create().
		Save(ctx)
	if err != nil {
		return nil, err
	}

	CreatePlaySessionMedia(ctx, client, playSession.ID, videoID)

	return playSession, nil
}

func CreatePlaySessionMedia(ctx context.Context, client *ent.Client, playSessionID int, videoID int) error {
	video, err := client.Video.Query().
		Where(video.ID(videoID)).
		Only(ctx)

	probeData, err := ffmpeg.Probe(video.Path)
	if err != nil {
		return err
	}

	var videoStreams []ffmpeg.ProbeStream = []ffmpeg.ProbeStream{}
	var audioStreams []ffmpeg.ProbeStream = []ffmpeg.ProbeStream{}
	for _, stream := range probeData.Streams {
		if stream.CodecType == "video" {
			videoStreams = append(videoStreams, stream)
		}

		if stream.CodecType == "audio" {
			audioStreams = append(audioStreams, stream)
		}
	}

	if len(videoStreams) == 0 {
		return fmt.Errorf("no video streams in file\n")
	}

	// Use first video stream by default
	// Todo: use all streams as sometimes there are lower resolution versions available so we wouldn't have to transcode as much
	videoStream := videoStreams[0]
	aspectRatio := float32(videoStream.Width) / float32(videoStream.Height)

	videoResolutions := []string{fmt.Sprintf("%dx%d", videoStream.Width, videoStream.Height)}
	for _, videoHeight := range videoHeights {
		if videoHeight >= videoStream.Height {
			continue
		}

		videoResolutions = append(videoResolutions, fmt.Sprintf("%dx%d", int(float32(videoHeight)*aspectRatio), videoHeight))
	}

	videoCodecs := []string{videoStream.CodecName}
	for _, codec := range ffmpeg.TranscodingVideoCodecs {
		if slices.Contains(videoCodecs, codec.Codec) {
			continue
		}

		videoCodecs = append(videoCodecs, codec.Codec)
	}

	media, err := client.PlaySessionMedia.Create().
		SetResolutions(videoResolutions).
		SetVideoCodecs(videoCodecs).
		SetSessionID(playSessionID).
		SetVideoID(videoID).
		Save(ctx)
	if err != nil {
		return err
	}

	// Create audio tracks
	for _, audioStream := range audioStreams {
		codecs := []string{audioStream.CodecName}
		for _, codec := range ffmpeg.TranscodingAudioCodecs {
			if slices.Contains(codecs, codec.Codec) {
				continue
			}

			codecs = append(codecs, codec.Codec)
		}

		track := client.AudioTrack.Create().
			SetMediaID(media.ID).
			SetName(audioStream.Profile).
			SetNrChannels(audioStream.Channels).
			SetChannelLayout(audioStream.ChannelLayout).
			SetCodecs(codecs)

		if audioStream.Tags["language"] != "und" {
			track.SetLanguage(audioStream.Tags["language"])
		}

		track.Save(ctx)
	}

	return nil
}
