package main

import (
	"go.uber.org/fx"

	"github.com/athulanilthomas/www/api/internal/config"
	"github.com/athulanilthomas/www/api/internal/handler"
	"github.com/athulanilthomas/www/api/internal/server"
	"github.com/athulanilthomas/www/api/internal/service"
	"github.com/athulanilthomas/www/api/internal/spotify"
)

func main() {
	fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(spotify.NewAuth),
		fx.Provide(spotify.NewSpotifyClient),
		fx.Provide(service.NewService),
		fx.Provide(handler.NewSpotifyHandler),
		fx.Provide(handler.NewAuthHandler),
		fx.Provide(server.NewRouter),
		fx.Provide(server.NewServer),

		fx.Invoke(server.RegisterRoutes),
		fx.Invoke(server.Start),
	).Run()
}
