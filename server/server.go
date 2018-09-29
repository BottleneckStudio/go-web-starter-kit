package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// Server represents
// the server for the app
type Server struct {
	Logger        *log.Logger
	Router        *chi.Mux
	TLSConfig     *tls.Config
	Srv           *http.Server
	ServerAddress string
}

// Config relates to
// the different server configuration
type Config struct {
	Server *Server
}

// New returns the
// instance of pointer
// to Server
//
// Sample:
// srv := server.New(&config)
func New(conf *Config) *Server {
	return &Server{
		Logger:        conf.Server.Logger,
		Router:        conf.Server.Router,
		ServerAddress: conf.Server.ServerAddress,
		Srv: &http.Server{
			Addr:         conf.Server.ServerAddress,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
			TLSConfig: &tls.Config{
				// Causes servers to use Go's default ciphersuite preferences,
				// which are tuned to avoid attacks. Does nothing on clients.
				PreferServerCipherSuites: true,
				// Only use curves which have assembly implementations
				CurvePreferences: []tls.CurveID{
					tls.CurveP256,
					tls.X25519, // Go 1.8 only
				},
				MinVersion: tls.VersionTLS12,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
					tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,

					// Best disabled, as they don't provide Forward Secrecy,
					// but might be necessary for some clients
					// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
					// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
				},
			},
			Handler: conf.Server.Router,
		},
	}
}

// Start boots-up a server
// that runs on plain HTTP
func (s *Server) Start() {
	s.Logger.Printf("Server is running on http://localhost%v", s.ServerAddress)
	err := s.Srv.ListenAndServe()
	if err != nil {
		s.Logger.Fatalf("Server is not starting due: %v", err)
	}
}

// StartTLS boots-up a server
// that runs on HTTPS
func (s *Server) StartTLS(certKey, privKey string) {
	s.Logger.Printf("Server is running on https://localhost%v", s.ServerAddress)
	err := s.Srv.ListenAndServeTLS(certKey, privKey)
	if err != nil {
		s.Logger.Fatalf("Server is not starting due: %v", err)
	}
}
