package dto

import "github.com/zmb3/spotify/v2"

type SpotifyDTO struct {
	Playing     bool               `json:"is_playing"`
	Item        *spotify.FullTrack `json:"item"`
	Progress    spotify.Numeric    `json:"progress_ms"`
	WorkerCount int                `json:"worker_count"`
}

func ToSpotifyDTO(resp *spotify.CurrentlyPlaying, workerCount int) *SpotifyDTO {
	if resp == nil {
		return nil
	}

	return &SpotifyDTO{
		Playing:     resp.Playing,
		Item:        resp.Item,
		Progress:    resp.Progress,
		WorkerCount: workerCount,
	}
}
