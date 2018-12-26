package crypt

import (
	"io/ioutil"
	"os"

	"github.com/mimoo/disco/libdisco"
	"github.com/traesk/kafft/util"
)

// Encrypt a file
func Encrypt(path, filename string, password []byte, delete bool) (string, error) {
	key := libdisco.Hash(password, 1024)

	files, err := util.ReadFiles(path, []string{filename})
	if err != nil {
		return "", err
	}
	zip, err := util.Zip(files)
	if err != nil {
		return "", err
	}

	// Encrypt
	encrypted := libdisco.Encrypt(key, zip)

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
