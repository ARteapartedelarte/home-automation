package main

import (
	"context"

	"github.com/jakewright/home-automation/libraries/go/bootstrap"
	"github.com/jakewright/home-automation/libraries/go/slog"
	"github.com/jakewright/home-automation/service.dmx/dmx"
	"github.com/jakewright/home-automation/service.dmx/handler"
	"github.com/jakewright/home-automation/service.dmx/universe"
)

//go:generate jrpc dmx.def

func main() {
	conf := struct{ UniverseNumber int }{}

	svc, err := bootstrap.Init(&bootstrap.Opts{
		ServiceName: "service.dmx",
		Config:      &conf,
		Firehose:    true,
	})

	if err != nil {
		slog.Panicf("Failed to initialise service: %v", err)
	}

	u := universe.New(conf.UniverseNumber)

	l := universe.Loader{
		ServiceName: "service.dmx",
		Universe:    u,
	}

	if err := l.FetchDevices(context.Background()); err != nil {
		slog.Panicf("Failed to load devices: %v", err)
	}

	h := handler.DMXController{
		Universe: u,
		Setter:   &dmx.OLA{},
	}

	r := handler.NewRouter().
		GetDevice(h.Read).
		UpdateDevice(h.Update)

	svc.Run(r)
}
