package tls

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"net"
	"time"
)

var (
	year              = 365 * 24 * time.Hour
	minimalRSAKeySize = 2048
)

func MakeSignedClientCert(cacert *x509.Certificate, caPrivateKey *rsa.PrivateKey, name string) (*x509.Certificate, *rsa.PrivateKey, error) {
	// Generate a private key.
	privateKey, err := rsa.GenerateKey(rand.Reader, minimalRSAKeySize)
	if err != nil {
		return nil, nil, err
	}

	// Generate subject key id.
	subjectKeyID := sha1.Sum(privateKey.PublicKey.N.Bytes())
	authorityKeyID := cacert.SubjectKeyId

	// Generate serial number with at least 20 bits of entropy.
	serialNumber, err := generateSerialNumber()
	if err != nil {
		return nil, nil, err
	}

	// Create certificate template.
	template := &x509.Certificate{
		Subject: pkix.Name{CommonName: name},

		NotBefore: time.Now().Add(-1 * time.Second),
		NotAfter:  time.Now().Add(year),

		SerialNumber:   serialNumber,
		SubjectKeyId:   subjectKeyID[:],
		AuthorityKeyId: authorityKeyID,

		SignatureAlgorithm: x509.SHA256WithRSA,

		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},

		BasicConstraintsValid: true,
	}

	// Sign Certificate
	derBytes, err := x509.CreateCertificate(
		rand.Reader,
		template,
		cacert,
		privateKey.Public(),
		caPrivateKey,
	)
	if err != nil {
		return nil, nil, err
	}

	// Parse Certificate into x509.Certificate.
	certs, err := x509.ParseCertificates(derBytes)
	if err != nil {
		return nil, nil, err
	}
	if len(certs) != 1 {
		return nil, nil, fmt.Errorf("expected 1 certificate, got %d", len(certs))
	}

	return certs[0], privateKey, nil
}

func MakeSignedServerCert(caCert *x509.Certificate, caPrivateKey *rsa.PrivateKey, dnsNames []string, ipAddresses []net.IP) (*x509.Certificate, *rsa.PrivateKey, error) {
	if len(dnsNames) == 0 && len(ipAddresses) == 0 {
		return nil, nil, fmt.Errorf("at least one DNS name or IP address must be provided")
	}

	// Generate a private key.
	privateKey, err := rsa.GenerateKey(rand.Reader, minimalRSAKeySize)
	if err != nil {
		return nil, nil, err
	}

	// Generate subject key id.
	subjectKeyID := sha1.Sum(privateKey.PublicKey.N.Bytes())
	authorityKeyID := caCert.SubjectKeyId

	// Generate serial number with at least 20 bits of entropy.
	serialNumber, err := generateSerialNumber()
	if err != nil {
		return nil, nil, err
	}

	// Create certificate template.
	template := &x509.Certificate{
		Subject: pkix.Name{CommonName: dnsNames[0]},

		NotBefore: time.Now().Add(-1 * time.Second),
		NotAfter:  time.Now().Add(year),

		SerialNumber:   serialNumber,
		SubjectKeyId:   subjectKeyID[:],
		AuthorityKeyId: authorityKeyID,

		SignatureAlgorithm: x509.SHA256WithRSA,

		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},

		DNSNames:    dnsNames,
		IPAddresses: ipAddresses,

		BasicConstraintsValid: true,
	}

	// Sign Certificate
	derBytes, err := x509.CreateCertificate(
		rand.Reader,
		template,
		caCert,
		privateKey.Public(),
		caPrivateKey,
	)
	if err != nil {
		return nil, nil, err
	}

	// Parse Certificate into x509.Certificate.
	certs, err := x509.ParseCertificates(derBytes)
	if err != nil {
		return nil, nil, err
	}
	if len(certs) != 1 {
		return nil, nil, fmt.Errorf("expected 1 certificate, got %d", len(certs))
	}

	return certs[0], privateKey, nil
}

func MakeSelfSignedCA(name string) (*x509.Certificate, *rsa.PrivateKey, error) {
	// Generate a private key.
	privateKey, err := rsa.GenerateKey(rand.Reader, minimalRSAKeySize)
	if err != nil {
		return nil, nil, err
	}

	// Generate authority key id and subject key id.
	keyID := sha1.Sum(privateKey.PublicKey.N.Bytes())

	// Generate serial number with at least 20 bits of entropy.
	serialNumber, err := generateSerialNumber()
	if err != nil {
		return nil, nil, err
	}

	// Create certificate template.
	template := &x509.Certificate{
		Subject: pkix.Name{CommonName: name},

		NotBefore: time.Now().Add(-1 * time.Second),
		NotAfter:  time.Now().Add(year),

		SerialNumber:   serialNumber,
		AuthorityKeyId: keyID[:],
		SubjectKeyId:   keyID[:],

		SignatureAlgorithm: x509.SHA256WithRSA,

		KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,

		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// Sign Certificate
	derBytes, err := x509.CreateCertificate(
		rand.Reader,
		template,
		template,
		privateKey.Public(),
		privateKey,
	)
	if err != nil {
		return nil, nil, err
	}

	// Parse Certificate into x509.Certificate.
	certs, err := x509.ParseCertificates(derBytes)
	if err != nil {
		return nil, nil, err
	}
	if len(certs) != 1 {
		return nil, nil, fmt.Errorf("expected 1 certificate, got %d", len(certs))
	}

	return certs[0], privateKey, nil
}

func generateSerialNumber() (*big.Int, error) {
	max := new(big.Int).Lsh(big.NewInt(1), 63)
	serialNumber, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil, err
	}

	return serialNumber, nil
}
