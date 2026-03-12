package service

import (
	"context"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/sync/singleflight"
)

type Cache struct {
	result    *spotify.CurrentlyPlaying
	fetchedAt time.Time
	mu        sync.RWMutex
}

type Service struct {
	client *spotify.Client
	pool   *ants.Pool
	group  singleflight.Group
	cache  Cache
}

type Job struct {
	ctx    context.Context
	result chan *spotify.CurrentlyPlaying
	err    chan error
}

func NewService(client *spotify.Client, pool *ants.Pool) *Service {
	// TODO: Implementation
	return &Service{
		client: client,
	}
}
