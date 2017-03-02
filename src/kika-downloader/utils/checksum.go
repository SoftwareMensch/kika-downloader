package utils

import (
	"crypto/sha256"
	"io/ioutil"
)

// SHA256FromFile calculate and return sha256 checksum of a file
func SHA256FromFile(filePath string) ([]byte, error) {
	sha256Builder := sha256.New()

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if _, err = sha256Builder.Write(data); err != nil {
		return nil, err
	}

	return sha256Builder.Sum(nil), nil
}
