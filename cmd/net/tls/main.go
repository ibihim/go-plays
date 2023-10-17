package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"

	playtls "github.com/ibihim/go-plays/pkg/tls"
)

func main() {
	if err := app(); err != nil {
		panic(err)
	}
}

func app() error {
	caCert, caKey, err := playtls.MakeSelfSignedCA("my-rootCA")
	if err != nil {
		return err
	}

	ip := net.ParseIP("127.0.0.1")
	if ip == nil {
		return fmt.Errorf("invalid ip address: %s", ip)
	}
	serverCert, serverKey, err := playtls.MakeSignedServerCert(caCert, caKey, []string{"localhost"}, []net.IP{ip})
	if err != nil {
		return err
	}

	clientCert, clientKey, err := playtls.MakeSignedClientCert(caCert, caKey, "my-client")
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errChan := make(chan error, 2)
	go func() {
		errChan <- runClient([]tls.Certificate{{
			Certificate: [][]byte{clientCert.Raw, caCert.Raw},
			PrivateKey:  clientKey,
		}}, caCert)

		cancel()
	}()

	go func() {
		if err := runServer(ctx, []tls.Certificate{{
			Certificate: [][]byte{serverCert.Raw, caCert.Raw},
			PrivateKey:  serverKey,
		}}, caCert); err != nil {
			errChan <- err
		}
	}()

	if err := <-errChan; err != nil {
		return err
	}

	return nil
}

func runClient(certs []tls.Certificate, caCert *x509.Certificate) error {
	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(caCert)

	tlsConfig := &tls.Config{
		Certificates:       certs,
		RootCAs:            caCertPool,
		InsecureSkipVerify: false,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	for i := 0; i < 10; i++ {
		res, err := client.Get("https://localhost:8443")
		if err != nil {
			fmt.Println(err)
			time.Sleep(200 * time.Millisecond)
			if i == 9 {
				return err
			}

			continue
		}

		defer res.Body.Close()

		if res.TLS != nil {
			fmt.Printf("Response from server '%s':\n\n", res.TLS.PeerCertificates[0].DNSNames[0])
		} else {
			fmt.Println("Response from server (TLS disabled):")
		}
		_, _ = io.Copy(os.Stdout, res.Body)

		return nil
	}

	return nil
}

func runServer(ctx context.Context, certs []tls.Certificate, clientCA *x509.Certificate) error {
	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(clientCA)

	tlsConfig := &tls.Config{
		Certificates: certs,
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert, // Makes the server require a known client certificate
	}

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
		Handler:   http.HandlerFunc(echoHandler),
	}

	errChan := make(chan error, 1)
	go func() {
		if err := server.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
		close(errChan)
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		return server.Shutdown(shutdownCtx)
	case err := <-errChan:
		return err
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		fmt.Fprintf(
			w,
			"TLS enabled with %s\n========================\n",
			r.TLS.PeerCertificates[0].Subject.CommonName,
		)
		defer fmt.Fprint(w, "========================\n")
	} else {
		fmt.Fprint(w, "!!!!! TLS disabled !!!!!\n")
	}

	fmt.Fprintf(w, "%s %s %s\r\n", r.Method, r.RequestURI, r.Proto)

	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\r\n", name, h)
		}
	}

	fmt.Fprint(w, "\r\n")

	_, err := io.Copy(w, r.Body)
	if err != nil {
		http.Error(w, "Failed to write body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
}
