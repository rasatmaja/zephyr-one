package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/fs"
	"math/big"
	"os"
	"time"
)

// Certificate ..
type Certificate struct {
	CertPath string
	KeyPath  string
}

// ICertificates ...
type ICertificates interface {
	GeneratePairKey() (pub, priv interface{}, err error)
	GenerateSerialNumber() (sernum *big.Int, err error)
	GenerateCertificate(pub, priv interface{}, sn *big.Int) (cert []byte, err error)
	SaveCertToFile(cert []byte) error
	SavePubKeyToFile(pub interface{}) error
	SavePrivKeyToFile(priv interface{}) error
	GenerateSelfSignedCertificates() (*Certificate, error)
}

// Cert ...
type Cert struct {
	BasePath        string
	CertFilename    string
	PrivKeyFilename string
	PubKeyFilename  string
	Permission      fs.FileMode
}

// Certificates ...
func (cert *Cert) Certificates() *Cert {
	return cert
}

// GeneratePairKey ...
func (cert *Cert) GeneratePairKey() (pub, priv interface{}, err error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, err
}

// GenerateSerialNumber ...
func (cert *Cert) GenerateSerialNumber() (sernum *big.Int, err error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	return serialNumber, err
}

// GenerateCertificate ...
func (cert *Cert) GenerateCertificate(pub, priv interface{}, sn *big.Int) (certs []byte, err error) {
	certTmpl := x509.Certificate{
		SerialNumber: sn,
		Subject: pkix.Name{
			Organization: []string{"Localhost Corp."},
		},
		DNSNames:  []string{"localhost"},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(3 * time.Hour),

		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Generate certificate file
	derBytes, err := x509.CreateCertificate(rand.Reader, &certTmpl, &certTmpl, pub, priv)
	return derBytes, err
}

// SaveCertToFile ...
func (cert *Cert) SaveCertToFile(certs []byte) error {
	pemCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certs})
	if pemCert == nil {
		return fmt.Errorf("Failed to encode certificate to PEM")
	}

	if err := os.WriteFile(cert.CertFilename, pemCert, cert.Permission); err != nil {
		return fmt.Errorf("Failed to write cert file: %v", err)
	}
	return nil
}

// SavePubKeyToFile ...
func (cert *Cert) SavePubKeyToFile(pub interface{}) error { return nil }

// SavePrivKeyToFile ...
func (cert *Cert) SavePrivKeyToFile(priv interface{}) error {

	// Generate key file
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return fmt.Errorf("Unable to marshal private key: %v", err)
	}
	pemKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})
	if pemKey == nil {
		return fmt.Errorf("Failed to encode key to PEM")
	}
	if err := os.WriteFile(cert.PrivKeyFilename, pemKey, cert.Permission); err != nil {
		return fmt.Errorf("Failed to write private key file: %v", err)
	}
	return nil

}

// GenerateSelfSignedCertificates ...
func (cert *Cert) GenerateSelfSignedCertificates() (*Certificate, error) {
	priv, pub, err := cert.GeneratePairKey()
	if err != nil {
		return nil, fmt.Errorf("Failed to generate private key: %v", err)
	}

	sn, err := cert.GenerateSerialNumber()
	if err != nil {
		return nil, fmt.Errorf("Failed to generate serial number: %v", err)
	}

	certByte, err := cert.GenerateCertificate(pub, priv, sn)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate certificate: %v", err)
	}

	err = cert.SaveCertToFile(certByte)
	if err != nil {
		return nil, fmt.Errorf("Failed to create certificate files: %v", err)
	}

	err = cert.SavePrivKeyToFile(priv)
	if err != nil {
		return nil, fmt.Errorf("Failed to create private key files: %v", err)
	}
	return &Certificate{
		CertPath: cert.CertFilename,
		KeyPath:  cert.PrivKeyFilename,
	}, nil
}
