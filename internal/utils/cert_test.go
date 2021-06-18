package utils

import (
	"os"
	"testing"
)

func TestCertificates(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cert := &Cert{
			CertFilename:    "cert.pem",
			PrivKeyFilename: "priv.pem",
			Permission:      0600,
		}

		cert.Certificates()
		_, err := cert.GenerateSelfSignedCertificates()

		if err != nil {
			t.Fail()
		}

		t.Cleanup(func() {
			os.Remove("cert.pem")
			os.Remove("priv.pem")
		})
	})
}

func TestSaveCertToFile(t *testing.T) {
	t.Run("error-encode-to-pem", func(t *testing.T) {
		cert := &Cert{
			CertFilename: "cert.pem",
			CertHeader: map[string]string{
				":": ":",
			},
		}
		err := cert.SaveCertToFile([]byte{0xff})
		if err == nil {
			t.Fail()
		}
	})
	t.Run("error-write-file", func(t *testing.T) {
		cert := &Cert{}
		err := cert.SaveCertToFile([]byte{0xff})
		if err == nil {
			t.Fail()
		}
	})
}

func TestSavePrivKeyToFile(t *testing.T) {
	t.Run("error-marshall-private-key", func(t *testing.T) {
		cert := &Cert{}
		err := cert.SavePrivKeyToFile([]byte{0xff})
		if err == nil {
			t.Fail()
		}
	})
	t.Run("error-encode-to-pem", func(t *testing.T) {
		cert := &Cert{
			PrivHeader: map[string]string{
				":": ":",
			},
		}
		priv, pub, err := cert.GeneratePairKey()
		if err != nil || pub == nil || priv == nil {
			t.Error(err)
			t.Fail()
		}
		err = cert.SavePrivKeyToFile(priv)
		if err == nil {
			t.Fail()
		}
	})
	t.Run("error-write-file", func(t *testing.T) {
		cert := &Cert{}
		priv, pub, err := cert.GeneratePairKey()
		if err != nil || pub == nil || priv == nil {
			t.Error(err)
			t.Fail()
		}
		err = cert.SavePrivKeyToFile(priv)
		if err == nil {
			t.Fail()
		}
	})
}

func TestSavePubKeyToFile(t *testing.T) {
	cert := Cert{}
	cert.SavePubKeyToFile([]byte("sss"))
}