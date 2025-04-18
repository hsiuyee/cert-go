package certgo

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"

	"github.com/Alonza0314/cert-go/model"
	"github.com/Alonza0314/cert-go/util"
	logger "github.com/Alonza0314/logger-go"
)

func CreateCsr(cfg model.Certificate) (*x509.CertificateRequest, error) {
	logger.Info("CreateCsr", "creating csr")

	// check csr exists
	if util.FileExists(cfg.CsrFilePath) {
		logger.Warn("CreateCsr", "csr already exists")
		return nil, errors.New("csr already exists")
	}

	var privateKey *ecdsa.PrivateKey
	var err error

	// check private key exists
	if !util.FileExists(cfg.KeyFilePath) {
		logger.Warn("CreateCsr", "private key does not exist")
		privateKey, err = CreatePrivateKey(cfg.KeyFilePath)
		if err != nil {
			return nil, err
		}
	}

	if privateKey == nil {
		privateKey, err = util.ReadPrivateKey(cfg.KeyFilePath)
		if err != nil {
			return nil, err
		}
	}

	template := &x509.CertificateRequest{
		Subject: pkix.Name{
			Organization: []string{cfg.Organization},
			CommonName:   cfg.CommonName,
		},
	}

	// create csr
	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, template, privateKey)
	if err != nil {
		logger.Error("CreateCsr", err.Error())
		return nil, err
	}

	// encode csr
	csrPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csrBytes,
	})

	// create directory exists
	if !util.FileDirExists(cfg.CsrFilePath) {
		logger.Warn("CreateCsr", util.FileDir(cfg.CsrFilePath)+" directory not exists, creating...")
		if err := util.FileDirCreate(cfg.CsrFilePath); err != nil {
			return nil, err
		}
		logger.Info("CreateCsr", util.FileDir(cfg.CsrFilePath)+" directory created")
	}

	// save csr
	if err := util.FileWrite(cfg.CsrFilePath, csrPEM, 0644); err != nil {
		return nil, err
	}

	logger.Info("CreateCsr", "csr created")

	csr, err := x509.ParseCertificateRequest(csrBytes)
	if err != nil {
		logger.Error("CreateCsr", err.Error())
		return nil, err
	}
	return csr, nil
}
