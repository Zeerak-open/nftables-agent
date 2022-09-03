package utils

import (
	"crypto/tls"
	"log"
)

func LoadCertificates(tlsCert, tlsKey string) *tls.Certificate {
	if (tlsKey != "") != (tlsCert != "") {
		log.Fatal("TLS server key file and cert file should both be present")
	}
	if tlsKey != "" && tlsCert != "" {
		cert, err := tls.LoadX509KeyPair(tlsCert, tlsKey)
		if err != nil {
			log.Fatalf("Couldn't load TLS server key pair, err: %s", err)
		}
		return &cert
	}

	return nil
}
