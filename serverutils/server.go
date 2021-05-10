package serverutils

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/haroldcampbell/go_utils/utils"
)

// CreateServer creates a server on the specified port
func CreateServer(serverPort string) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%s", serverPort),
		WriteTimeout: 15 * time.Second, // Good practice: enforce timeouts for servers you create!
		ReadTimeout:  15 * time.Second,
	}
}

// ServerHTTP starts a server with a HTTP listening on the specified port.
func ServerHTTP(src string, handler http.Handler, serverPort string) {
	stem := fmt.Sprintf("%s>ServerHTTP", src)
	httpServer := CreateServer(serverPort)
	httpServer.Handler = Logger(src, handler)

	utils.Log(stem, "[ServerHTTP] Starting server")
	utils.Log(stem, "[ServerHTTP] http listening on: %s", serverPort)
	log.Fatal(httpServer.ListenAndServe())
}

// ServerHTTPMuted starts a server with a HTTP listening on the specified port but without the Logger.
func ServerHTTPMuted(src string, handler http.Handler, serverPort string) {
	stem := fmt.Sprintf("%s>ServerHTTPMuted", src)
	httpServer := CreateServer(serverPort)
	httpServer.Handler = handler

	utils.Log(stem, "[ServerHTTPMuted] Starting server")
	utils.Log(stem, "[ServerHTTPMuted] listening on: %s", serverPort)
	log.Fatal(httpServer.ListenAndServe())
}

// ServerHTTPS starts a server with a HTTPS connection listening on port 443.
func ServerHTTPS(src string, handler http.Handler, allowedRegexPath string, cert string, key string) {
	stem := fmt.Sprintf("%s>ServerHTTPS", src)
	serverPort := "443"
	httpServer := CreateServer(serverPort)
	httpServer.Handler = BlockIPCalls(src, allowedRegexPath, Logger(src, handler))

	utils.Log(stem, "[ServerHTTPS] Starting server")
	utils.Log(stem, "[ServerHTTPS] http listening on: %s", serverPort)
	log.Fatal(httpServer.ListenAndServeTLS(
		cert,
		key,
	))
}

// Logger handler
func Logger(src string, h http.Handler) http.Handler {
	stem := fmt.Sprintf("%s>Logger", src)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.Log(stem, "Request: %s%v RemoteAddr: %v", r.Host, r.URL.Path, r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}

// RedirectTLS ...
func RedirectTLS(w http.ResponseWriter, r *http.Request) {
	stem := "RedirectTLS"
	newTargetURL := fmt.Sprintf("https://%s%s", r.Host, r.RequestURI)

	utils.Log(stem, "Redirecting http request %s -> %s\n", r.RequestURI, newTargetURL)
	utils.Log(stem, "-> RemoteAddr: %v\n", r.RemoteAddr)
	http.Redirect(w, r, newTargetURL, http.StatusMovedPermanently)
}

// BlockIPCalls is used when someone tries to access the ip address directly.
// Let's encrypt doesn't provide certs for ip address
// Example of the allowedRegexPath is `^(www.|api.)?tealoverflow.com$`
func BlockIPCalls(src string, allowedRegexPath string, h http.Handler) http.Handler {
	stem := fmt.Sprintf("%s>BlockIPCalls", src)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.Host, "localhost:") {
			utils.Log(stem, "Serving localhost: %v -> %v", r.RemoteAddr, r.RequestURI)
			return
		}

		rx, _ := regexp.Compile(allowedRegexPath)
		isIPAddress := !rx.MatchString(r.Host)

		if isIPAddress {
			// Simply ignore the request
			utils.Log(stem, "Ignoring direct ip address call from: %v", r.RemoteAddr)
			utils.Log(stem, "-> Request %s -> %s", r.Host, r.RequestURI)
			return
		}
		h.ServeHTTP(w, r)
	})
}
