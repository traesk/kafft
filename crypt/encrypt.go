package crypt

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mimoo/disco/libdisco"
	"github.com/traesk/kafft/util"
)

// Encrypt a file
func Encrypt(path, filename string, password []byte, delete bool) (string, error) {
	key := libdisco.Hash(password, 1024)

	// Read the file
	inputFile, err := ioutil.ReadFile(path + filename)
	if err != nil {
		return "", err
	}

	// Arbitrary assumption that 256 bytes is enough for name + type
	fileName := make([]byte, 256)

	// Put the name in the 256 byte sized slice
	n := hex.Encode(fileName, []byte(filename))
	if n == 0 {
		return "", err
	}
	fmt.Println(string(fileName))

	// Append the file to filename
	plain := append(fileName, inputFile...)

	// Encrypt
	encrypted := libdisco.Encrypt(key, plain)

	// Save the file
	name, err := util.GenerateName(16)
	if err != nil {
		return "", err
	}
	output := path + name
	ioutil.WriteFile(output, encrypted, os.FileMode(0777))

	// Optionally remove the file
	if delete {
		if err := os.Remove(filename); err != nil {
			return "", err
		}
	}
	return name, nil
}
