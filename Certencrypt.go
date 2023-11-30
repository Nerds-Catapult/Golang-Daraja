package darajaAuth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"
	"os"
)

func openSSlEncrypt(data, certPath string) (string, error) {
	cert, err := loadCertificate(certPath)
	if err != nil {
		return "", err
	}

	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, cert.PublicKey.(*rsa.PublicKey), []byte(data))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func loadCertificate(certPath string) (*x509.Certificate, error) {
	certFile, err := os.Open(certPath)
	if err != nil {
		return nil, err
	}
	defer certFile.Close()

	certBytes, err := io.ReadAll(certFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(certBytes)
	if block == nil {
		return nil, errors.New("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return cert, nil
}