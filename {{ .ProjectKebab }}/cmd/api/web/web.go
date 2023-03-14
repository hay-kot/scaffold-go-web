// Package web contains the web server for the api
package web

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/hay-kot/safeserve/errchain"
	"github.com/hay-kot/safeserve/errtrace"
	"github.com/hay-kot/safeserve/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"{{ .Scaffold.gomod }}/cmd/api/web/controller"
	"{{ .Scaffold.gomod }}/internal/web/mid"
)

// Web struct is a wrapper for the server and router that provides
// a abstraction for the server Start and Stop commands. This exists
// mostly to make testing the API easier. In most cases you will only
// every call start and wait for it to return.
type Web struct {
	svr  *server.Server
	mux  *chi.Mux
	conf *Config
	l    zerolog.Logger
}

type WebArgs struct {
	Conf  *Config
	Build string
}

func (web *WebArgs) Validate() error {
	if web.Conf == nil {
		return errors.New("config is required")
	}

	if web.Build == "" {
		return errors.New("build is required")
	}

	return nil
}

func New(args *WebArgs) *Web {
	if err := args.Validate(); err != nil {
		panic(err)
	}

	conf := args.Conf

	if conf.Mode == ModeDevelopment {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// =========================================================================
	// Mux
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(mid.Logger(log.Logger))
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: strings.Split(conf.Web.AllowedOrigins, ","),
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// =========================================================================
	// Web Server and Routes

	svr := server.NewServer(
		server.WithHost(conf.Web.Host),
		server.WithPort(conf.Web.Port),
		server.WithIdleTimeout(conf.Web.IdleTimeout),
		server.WithReadTimeout(conf.Web.ReadTimeout),
		server.WithWriteTimeout(conf.Web.WriteTimeout),
	)

	web := &Web{
		mux:  r,
		svr:  svr,
		conf: args.Conf,
		l:    log.Logger,
	}

	chain := errchain.New(web.errHandler)

	ctrl := controller.New(log.Logger, args.Build)

	web.mux.Get("/status", chain.ToHandlerFunc(ctrl.Status))
	web.mux.Get("/error", chain.ToHandlerFunc(ctrl.ErrorMaker))

	log.Info().Str("host", conf.Web.Host).Str("port", conf.Web.Port).Msgf("starting server on")
	return web
}

func (web *Web) errHandler(h errchain.Handler) http.Handler {
	type errMsg struct {
		Details string `json:"details"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.ServeHTTP(w, r)
		if err != nil {
			if web.conf.Mode == ModeDevelopment {
				fmt.Println(errtrace.TraceString(err))
			}

			web.l.Err(err).Msg("error in handler")

			status := http.StatusInternalServerError
			resp := errMsg{Details: "internal server error"}

			// Check for know request error types here.
			switch {
			case errors.Is(err, controller.ErrMakeError):
				resp.Details = "error maker error"
			}

			err := server.JSON(w, status, resp)
			if err != nil {
				log.Err(err).Msg("error while writing response")
			}
		}
	})
}

func (web *Web) Start() error {
	return web.svr.Start(web.mux)
}

func (web *Web) Shutdown(msg string) error {
	return web.svr.Shutdown(msg)
}
