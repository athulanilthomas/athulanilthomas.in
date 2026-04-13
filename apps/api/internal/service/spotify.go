package service

import (
	"context"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/sync/singleflight"
)

type cache struct {
	result    *spotify.CurrentlyPlaying
	fetchedAt time.Time
	mu        sync.RWMutex
}

type Service struct {
	client *spotify.Client
	pool   *ants.PoolWithFunc
	group  singleflight.Group
	cache  cache
	mux    sync.RWMutex
}

type job struct {
	ctx    context.Context
	result chan *spotify.CurrentlyPlaying
	err    chan error
}

func NewService(client *spotify.Client) (*Service, error) {
	svc := &Service{
		client: client,
		group:  singleflight.Group{},
		cache:  cache{},
	}

	pool, err := ants.NewPoolWithFunc(10, func(payload any) {
		job := payload.(*job)

		svc.mux.RLock()
		c := svc.client
		svc.mux.RUnlock()

		if result, err := c.PlayerCurrentlyPlaying(job.ctx); err != nil {
			job.err <- err
		} else {
			job.result <- result
		}
	})

	if err != nil {
		return nil, err
	}

	svc.pool = pool

	return svc, nil
}

func (s *Service) GetCurrentlyPlaying(ctx context.Context) (*spotify.CurrentlyPlaying, error) {
	s.cache.mu.RLock()

	if s.cache.result != nil && time.Since(s.cache.fetchedAt) < 2*time.Second {
		s.cache.mu.RUnlock()
		return s.cache.result, nil
	}

	s.cache.mu.RUnlock()

	res, e, _ := s.group.Do("currently-playing", func() (any, error) {
		j := &job{
			ctx:    ctx,
			result: make(chan *spotify.CurrentlyPlaying, 1),
			err:    make(chan error, 1),
		}

		s.pool.Invoke(j)

		select {
		case r := <-j.result:
			s.cache.mu.Lock()
			s.cache.result = r
			s.cache.fetchedAt = time.Now()
			s.cache.mu.Unlock()
			return r, nil
		case e := <-j.err:
			return nil, e
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	})

	if e != nil {
		return nil, e
	}

	return res.(*spotify.CurrentlyPlaying), nil
}

func (s *Service) GetWorkerCount(ctx context.Context) int {
	count := s.pool.Running()
	return count
}

func (s *Service) SetClient(client *spotify.Client) {
	s.mux.Lock()
	s.client = client
	s.mux.Unlock()
}
