// Code generated by jrpc. DO NOT EDIT.

package handler

import (
	context "context"
	http "net/http"

	network "github.com/jakewright/home-automation/libraries/go/network"
	oops "github.com/jakewright/home-automation/libraries/go/oops"
	router "github.com/jakewright/home-automation/libraries/go/router"
	slog "github.com/jakewright/home-automation/libraries/go/slog"
	def "github.com/jakewright/home-automation/service.scene/def"
)

type controller interface {
	CreateScene(*request, *def.CreateSceneRequest) (*def.CreateSceneResponse, error)
	ReadScene(*request, *def.ReadSceneRequest) (*def.ReadSceneResponse, error)
	ListScenes(*request, *def.ListScenesRequest) (*def.ListScenesResponse, error)
	DeleteScene(*request, *def.DeleteSceneRequest) (*def.DeleteSceneResponse, error)
	SetScene(*request, *def.SetSceneRequest) (*def.SetSceneResponse, error)
}

type request struct {
	context.Context
	*http.Request
}

// NewRouter returns a router with appropriate handlers set
func NewRouter(c controller) *router.Router {
	r := router.New()

	r.Handle("POST", "/scenes", func(w http.ResponseWriter, r *http.Request) {
		body := &def.CreateSceneRequest{}
		if err := network.DecodeRequest(r, body); err != nil {
			err = oops.Wrap(err, oops.ErrBadRequest, "failed to decode request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		if err := body.Validate(); err != nil {
			err = oops.Wrap(err, oops.ErrBadRequest, "failed to validate request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		req := &request{
			Context: r.Context(),
			Request: r,
		}

		rsp, err := c.CreateScene(req, body)
		if err != nil {
			err = oops.WithMessage(err, "failed to handle request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		network.WriteJSONResponse(w, rsp)
	})

	r.Handle("GET", "/scene", func(w http.ResponseWriter, r *http.Request) {
		body := &def.ReadSceneRequest{}
		if err := network.DecodeRequest(r, body); err != nil {
			err = oops.Wrap(err, oops.ErrBadRequest, "failed to decode request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		if err := body.Validate(); err != nil {
			err = oops.Wrap(err, oops.ErrBadRequest, "failed to validate request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		req := &request{
			Context: r.Context(),
			Request: r,
		}

		rsp, err := c.ReadScene(req, body)
		if err != nil {
			err = oops.WithMessage(err, "failed to handle request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		network.WriteJSONResponse(w, rsp)
	})

	r.Handle("GET", "/scenes", func(w http.ResponseWriter, r *http.Request) {
		body := &def.ListScenesRequest{}
		if err := network.DecodeRequest(r, body); err != nil {
			err = oops.Wrap(err, oops.ErrBadRequest, "failed to decode request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		if err := body.Validate(); err != nil {
			err = oops.Wrap(err, oops.ErrBadRequest, "failed to validate request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		req := &request{
			Context: r.Context(),
			Request: r,
		}

		rsp, err := c.ListScenes(req, body)
		if err != nil {
			err = oops.WithMessage(err, "failed to handle request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		network.WriteJSONResponse(w, rsp)
	})

	r.Handle("DELETE", "/scene", func(w http.ResponseWriter, r *http.Request) {
		body := &def.DeleteSceneRequest{}
		if err := network.DecodeRequest(r, body); err != nil {
			err = oops.Wrap(err, oops.ErrBadRequest, "failed to decode request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		if err := body.Validate(); err != nil {
			err = oops.Wrap(err, oops.ErrBadRequest, "failed to validate request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		req := &request{
			Context: r.Context(),
			Request: r,
		}

		rsp, err := c.DeleteScene(req, body)
		if err != nil {
			err = oops.WithMessage(err, "failed to handle request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		network.WriteJSONResponse(w, rsp)
	})

	r.Handle("POST", "/scene/set", func(w http.ResponseWriter, r *http.Request) {
		body := &def.SetSceneRequest{}
		if err := network.DecodeRequest(r, body); err != nil {
			err = oops.Wrap(err, oops.ErrBadRequest, "failed to decode request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		if err := body.Validate(); err != nil {
			err = oops.Wrap(err, oops.ErrBadRequest, "failed to validate request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		req := &request{
			Context: r.Context(),
			Request: r,
		}

		rsp, err := c.SetScene(req, body)
		if err != nil {
			err = oops.WithMessage(err, "failed to handle request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		network.WriteJSONResponse(w, rsp)
	})

	return r
}
