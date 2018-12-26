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
	/*
		// Get the name from the first 256 bytes
		unpackedName, err := hex.DecodeString(string(decrypted[:256]))

		file := decrypted[256:]
	*/

	//output := path + filename + ".hei"
	//fmt.Println("Writing: ", output)
	file, err := ioutil.TempFile(path, "temp_*")
	if err != nil {
		return "", err
	}
	file.Write(decrypted)
	defer file.Close()
	//ioutil.WriteFile(output, decrypted, os.FileMode(0777))

	return file.Name(), nil
}
