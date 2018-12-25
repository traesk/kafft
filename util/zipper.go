package util

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
)

// ZipFiles to make them ready for encryption
func ZipFiles(source string) ([]byte, error) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	files, err := ioutil.ReadDir(source)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		fmt.Println("Reading file: ", f.Name())
		file, err := ioutil.ReadFile(source + "/" + f.Name())
		if err != nil {
			return nil, err
		}
		io, err := w.Create(f.Name())
		if err != nil {
			return nil, err
		}
		_, err = io.Write(file)
		if err != nil {
			return nil, err
		}
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnzipFiles after unencrypting it
func UnzipFiles() {

}
