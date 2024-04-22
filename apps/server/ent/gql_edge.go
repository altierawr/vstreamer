// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (at *AudioTrack) Media(ctx context.Context) (*PlaySessionMedia, error) {
	result, err := at.Edges.MediaOrErr()
	if IsNotLoaded(err) {
		result, err = at.QueryMedia().Only(ctx)
	}
	return result, err
}

func (l *Library) Videos(ctx context.Context) (result []*Video, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = l.NamedVideos(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = l.Edges.VideosOrErr()
	}
	if IsNotLoaded(err) {
		result, err = l.QueryVideos().All(ctx)
	}
	return result, err
}

func (ps *PlaySession) Clients(ctx context.Context) (result []*PlaybackClient, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ps.NamedClients(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ps.Edges.ClientsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ps.QueryClients().All(ctx)
	}
	return result, err
}

func (ps *PlaySession) Media(ctx context.Context) (*PlaySessionMedia, error) {
	result, err := ps.Edges.MediaOrErr()
	if IsNotLoaded(err) {
		result, err = ps.QueryMedia().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (psm *PlaySessionMedia) AudioTracks(ctx context.Context) (result []*AudioTrack, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = psm.NamedAudioTracks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = psm.Edges.AudioTracksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = psm.QueryAudioTracks().All(ctx)
	}
	return result, err
}

func (psm *PlaySessionMedia) Video(ctx context.Context) (*Video, error) {
	result, err := psm.Edges.VideoOrErr()
	if IsNotLoaded(err) {
		result, err = psm.QueryVideo().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (psm *PlaySessionMedia) Session(ctx context.Context) (*PlaySession, error) {
	result, err := psm.Edges.SessionOrErr()
	if IsNotLoaded(err) {
		result, err = psm.QuerySession().Only(ctx)
	}
	return result, err
}

func (psm *PlaySessionMedia) VideoCodecs(ctx context.Context) (result []*VideoCodec, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = psm.NamedVideoCodecs(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = psm.Edges.VideoCodecsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = psm.QueryVideoCodecs().All(ctx)
	}
	return result, err
}

func (pc *PlaybackClient) Session(ctx context.Context) (*PlaySession, error) {
	result, err := pc.Edges.SessionOrErr()
	if IsNotLoaded(err) {
		result, err = pc.QuerySession().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (v *Video) PlaySessionMedias(ctx context.Context) (result []*PlaySessionMedia, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = v.NamedPlaySessionMedias(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = v.Edges.PlaySessionMediasOrErr()
	}
	if IsNotLoaded(err) {
		result, err = v.QueryPlaySessionMedias().All(ctx)
	}
	return result, err
}

func (v *Video) Library(ctx context.Context) (*Library, error) {
	result, err := v.Edges.LibraryOrErr()
	if IsNotLoaded(err) {
		result, err = v.QueryLibrary().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (vc *VideoCodec) Media(ctx context.Context) (*PlaySessionMedia, error) {
	result, err := vc.Edges.MediaOrErr()
	if IsNotLoaded(err) {
		result, err = vc.QueryMedia().Only(ctx)
	}
	return result, MaskNotFound(err)
}
