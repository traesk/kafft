package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

// SaveInfo of the encrypted file(s)
func SaveInfo(name, content string) error {

	err := ioutil.WriteFile(name+".txt", []byte(content), os.FileMode(0777))
	if err != nil {
		return err
	}
	fmt.Println("Info saved to: ", name+".txt")
	return nil
}

// LoadInfo to unencrypt file(s)
func LoadInfo() {

}
