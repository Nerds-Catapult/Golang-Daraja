package darajaAuth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"os"
)

type certificationError struct {
	Context string
	Err     error
}

// Error implements the error interface for certificationError.
func (e *certificationError) Error() string {
	return fmt.Sprintf("%s: %v", e.Context, e.Err)
}

func openSSlEncrypt(data, certPath string) (string, error) {
	cert, err := loadCertificate(certPath)
	if err != nil {
		return "", &certificationError{Context: "failed to load certificate", Err: err}
	}

	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, cert.PublicKey.(*rsa.PublicKey), []byte(data))
	if err != nil {
		return "", &certificationError{Context: "encryption failed", Err: err}
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// loadCertificate loads and parses the X.509 certificate.
func loadCertificate(certPath string) (*x509.Certificate, error) {
	certFile, err := os.Open(certPath)
	if err != nil {
		return nil, &certificationError{Context: "failed to open certificate file", Err: err}
	}
	defer certFile.Close()

	certBytes, err := io.ReadAll(certFile)
	if err != nil {
		return nil, &certificationError{Context: "failed to read certificate file", Err: err}
	}

	block, _ := pem.Decode(certBytes)
	if block == nil {
		return nil, &certificationError{Context: "failed to parse certificate PEM", Err: errors.New("no PEM block found")}
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, &certificationError{Context: "failed to parse certificate", Err: err}
	}

	return cert, nil
}
