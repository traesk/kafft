package crypt

import (
	"fmt"
	"io/ioutil"

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
		fmt.Println("error decrypting")
		return "", err
	}

	// Create temp file
	file, err := ioutil.TempFile(path, "temp_*")
	if err != nil {
		return "", err
	}
	// Write to it
	file.Write(decrypted)
	defer file.Close()

	return file.Name(), nil
}
