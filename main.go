package main

import (
	"log"
	"net/http"
	"crypto/tls"
	"os"

	"github.com/sudocat/nova-proxy/config"
	"github.com/gookit/color"
)

func init() {
	config.LoadEnv()
	config.ReadConfigFile()
}

func main() {
	skipTlsVerify := os.Getenv("SKIP_TLS_VERIFY")
	insecureSkipVerify := skipTlsVerify == "true"
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecureSkipVerify}
	config.SetUpLocations()

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	color.Info.Printf("Nova proxy running on http://0.0.0.0:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
