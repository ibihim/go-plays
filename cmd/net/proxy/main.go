package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func proxy(upstreamURL *url.URL, ca *x509.CertPool) http.Handler {
	transport := (http.DefaultTransport.(*http.Transport)).Clone()
	transport.TLSClientConfig = &tls.Config{RootCAs: ca}

	proxy := httputil.NewSingleHostReverseProxy(upstreamURL)
	proxy.Transport = &http.Transport{}

	return proxy
}

func loadCA(filePath string) (*x509.CertPool, error) {
	caFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	pool := x509.NewCertPool()
	if ok := pool.AppendCertsFromPEM([]byte(caFile)); !ok {
		return nil, errors.New("parsing upstream CA certificate")
	}

	return pool, nil
}

func start(address string, handler http.Handler) error {
	server := http.Server{
		Addr: address, Handler: handler,
	}

	return server.ListenAndServe()
}

func startTLS(address string, certFilePath, keyFilePath string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(w, "ok")
	})

	server := http.Server{
		Addr:    address,
		Handler: mux,
	}

	return server.ListenAndServeTLS(certFilePath, keyFilePath)
}

func main() {
	certFile := "./certs/server.crt"
	keyFile := "./certs/server.key"
	caFile := "./certs/ca.crt"
	listenAddress := ":8443"
	proxyAddress := ":4433"
	listenURL, err := url.Parse(fmt.Sprintf("https://localhost%s", listenAddress))
	if err != nil {
		panic(err)
	}

	ca, err := loadCA(caFile)
	if err != nil {
		panic(err)
	}

	go func() {
		fmt.Printf("Proxy listens to %s and proxies to %s\n", proxyAddress, listenURL)
		if err := start(proxyAddress, proxy(listenURL, ca)); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("Server listens to %s\n", listenAddress)
	if err := startTLS(listenAddress, certFile, keyFile); err != nil {
		panic(err)
	}

}
