// Package controller provides the controller for the api
package controller

import (
	"errors"
	"net/http"

	"github.com/hay-kot/safeserve/errtrace"
	"github.com/hay-kot/safeserve/server"
	"github.com/rs/zerolog"
)

type Controller struct {
	l     zerolog.Logger
	build string
}

func New(l zerolog.Logger, build string) *Controller {
	return &Controller{
		l:     l,
		build: build,
	}
}

type StatusResponse struct {
	Status string `json:"status"`
	Build  string `json:"build"`
}

func (c *Controller) Status(w http.ResponseWriter, r *http.Request) error {
	return server.JSON(w, http.StatusOK, StatusResponse{
		Status: "ok",
		Build:  c.build,
	})
}

var ErrMakeError = errors.New("errorMaker error")

func (c *Controller) ErrorMaker(w http.ResponseWriter, r *http.Request) error {
	return errtrace.TraceWrap(ErrMakeError, "errorMaker endpoint error")
}
