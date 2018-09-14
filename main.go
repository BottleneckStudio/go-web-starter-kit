package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"github.com/BottleneckStudio/go-web-starter-kit/server"
)

func main() {
	logger := log.New(os.Stdout, "DEBUG: ", log.LstdFlags|log.Lshortfile)
	r := chi.NewRouter()
	serverConfig := server.Config{
		Server: &server.Server{
			Logger: logger,
			Router: r,
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
			ServerAddress: ":1437",
		},
	}

	srv := server.New(&serverConfig)

	srv.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	srv.Start()
	// fmt.Println("vim-go")
}
