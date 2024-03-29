package http

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"doublequote"
	"doublequote/utils"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"golang.org/x/crypto/acme/autocert"
)

const ShutdownTimeout = 5

type Server struct {
	ln     net.Listener
	server *http.Server
	router *chi.Mux
	now    func() time.Time

	// Services used by the various HTTP routes.
	UserService       dq.UserService
	CryptoService     dq.CryptoService
	SessionService    dq.SessionService
	CollectionService dq.CollectionService
	FeedService       dq.FeedService

	Config dq.Config
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{},
		router: chi.NewRouter(),
		now:    time.Now,
	}

	s.router.Use(chimw.Logger)

	// TODO change this?
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	s.router.Route("/api/pub", func(r chi.Router) {
		r.Use(setContentType)

		s.registerPublicUserRoutes(r)
	})

	s.router.Route("/api", func(r chi.Router) {
		r.Use(s.requireAuth)
		r.Use(setContentType)

		s.registerUserRoutes(r)
		s.registerCollectionRoutes(r)
	})

	s.router.Route("/", func(r chi.Router) {
		staticFs, err := fs.Sub(dq.Assets, "assets/frontend")
		if err != nil {
			log.Fatalf("http: %s", err)
		}

		fserver := http.FileServer(http.FS(staticFs))

		r.Get("/*", func(w http.ResponseWriter, req *http.Request) {
			// TODO currently we open index.html for all frontend requests
			indexFile, err := staticFs.Open("index.html")
			if err != nil {
				Error(w, req, err)
				return
			}

			_, err = staticFs.Open(strings.TrimPrefix(req.URL.Path, "/"))
			if errors.Is(err, fs.ErrNotExist) {
				io.Copy(w, indexFile)
				return
			}

			fserver.ServeHTTP(w, req)
		})
	})

	s.server.Handler = s.router

	return s
}

func (s *Server) Open() (err error) {
	// Open a listener on our bind address.
	if s.Config.HTTP.Domain != "" {
		s.ln = autocert.NewListener(s.Config.HTTP.Domain)
	} else {
		if s.ln, err = net.Listen("tcp", ":"+s.Config.HTTP.Port); err != nil {
			return err
		}
	}

	go s.server.Serve(s.ln)

	return nil
}

func (s *Server) Close() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}

// requireAuth is a middleware for requiring authentication
func (s *Server) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get session, make sure it's not empty.
		sess, err := s.SessionService.Get(r)
		if err != nil {
			Error(w, r, err)
			return
		}
		if sess == nil {
			Error(w, r, dq.Errorf(dq.EUNAUTHORIZED, dq.ErrUnauthorized))
			return
		}

		// Find user associated with session, then add it to the request context.
		u, err := s.UserService.FindUser(r.Context(), dq.UserFilter{ID: utils.IntPtr(sess.UserID())}, dq.UserInclude{})
		if err != nil {
			Error(w, r, err)
			return
		}

		r = r.WithContext(dq.NewContextWithUser(r.Context(), u))
		next.ServeHTTP(w, r)
	})
}

// setContentType is a middleware that sets the response Content-type to application/json.
func setContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		next.ServeHTTP(w, r)
	})
}
