package server

import (
	"fmt"
	putfile "main/src/Minio/MinioActions/PutFile"
	JsonSuccessResponse "main/src/Utils/JsonSuccessResponse"
	"net/http"

	"github.com/rs/cors"
	"github.com/unrolled/secure"
)

func RunAPI() error {
	secureMiddleware := secure.New(secure.Options{
		ContentSecurityPolicy: "default-src 'none'; connect-src 'self'; object-src 'none'; frame-ancestors 'none';",
	})

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		JsonSuccessResponse.JsonSuccessResponse(w)
	})

	mux.HandleFunc("/api/uploads/tony-hawk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		putfile.FPutFile("media-app-one", "images/tony-hawk.webp", "images/old-tony-hawk.webp")

		JsonSuccessResponse.JsonSuccessResponse(w)
	})

	mux.HandleFunc("/api/uploads/table", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		putfile.FPutFile("media-app-one", "images/table.png", "images/table.png")

		JsonSuccessResponse.JsonSuccessResponse(w)
	})

	c := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(mux)

	secureHandler := secureMiddleware.Handler(handler)

	fmt.Println("brian_media_app_one latest version on a bridge network is mapping from https://172.17.0.3:8080/TCP (in app) to https://192.168.1.195:5000 (on host) with no volumes -- running with reduced size via 2 step setup... fixed docker image... maybe...")

	err := http.ListenAndServeTLS(":8080", ".certs/barcs_media_app_one_ca.crt", ".certs/barcs_media_app_one_ca.key", secureHandler)

	if err != nil {
		return err
	}

	return nil
}
