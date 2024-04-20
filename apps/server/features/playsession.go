package features

import (
	"github.com/altierawr/vstreamer/ent"
	"github.com/altierawr/vstreamer/ffmpeg"
	"golang.org/x/net/context"
)

type Streamable struct {
	streams map[int]ffmpeg.Transcoder
}

var Streamables map[int]ffmpeg.Transcoder

func InitPlaySession(ctx context.Context, client *ent.Client, id int) error {
	return nil
}
