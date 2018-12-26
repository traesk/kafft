package util

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
)

/*

 */

// ZipFiles to make them ready for encryption
func ZipFiles(dir string) ([]byte, error) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		//fmt.Println("Reading file: ", f.Name())
		file, err := ioutil.ReadFile(dir + "/" + f.Name())
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

// Zip one or more files, return the zip
func Zip(files []File) ([]byte, error) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	for _, f := range files {

		io, err := w.Create(f.Name)
		if err != nil {
			return nil, err
		}
		_, err = io.Write(f.Body)
		if err != nil {
			return nil, err
		}
	}
	err := w.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ReadFiles either a lite of files, or a directory
func ReadFiles(dir string, filenames []string) ([]File, error) {
	var files []File

	// We got a folder instead of a list of files
	if len(filenames) < 1 {
		fileInfos, err := ioutil.ReadDir(dir)
		if err != nil {
			return nil, err
		}
		// Fill the list of files
		for _, f := range fileInfos {
			filenames = append(filenames, f.Name())
		}
	}
	// Add them to the slice, ready to zip
	for _, name := range filenames {
		data, err := ioutil.ReadFile(dir + name)
		if err != nil {
			return nil, err
		}
		files = append(files, File{name, data})
	}
	return files, nil
}

// UnzipFiles after unencrypting it
func UnzipFiles() {

}

// File used by ZipFiles
type File struct {
	Name string
	Body []byte
}
