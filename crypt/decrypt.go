package crypt

import (
	"encoding/hex"
	"io/ioutil"
	"os"

	"github.com/mimoo/disco/libdisco"
)

// Decrypt a file
func Decrypt(path, filename string, password []byte) (string, error) {
	key := libdisco.Hash(password, 1024)

	// Read the file
	inputFile, err := ioutil.ReadFile(path + filename)
	if err != nil {
		return "", err
	}

	decrypted, err := libdisco.Decrypt(key, inputFile)
	if err != nil {
		return "", err
	}

	// Put the name in the 256 byte sized slice
	unpackedName, err := hex.DecodeString(string(decrypted[:256]))

	file := decrypted[256:]

	output := path + string(unpackedName)
	ioutil.WriteFile(output, file, os.FileMode(0777))

	return string(unpackedName), nil
}
