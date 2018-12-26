package util

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
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
}*/

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

// Unzip and write to outputDir
func Unzip(src, outputDir string) ([]string, error) {
	var filenames []string
	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer rc.Close()
		fp := filepath.Join(outputDir + f.Name)

		filenames = append(filenames, fp)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fp, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fp), os.ModePerm); err != nil {
				return nil, err
			}

			outputFile, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return nil, err
			}

			_, err = io.Copy(outputFile, rc)
			outputFile.Close()
			if err != nil {
				return nil, err
			}
		}
	}
	// Remove temp file
	err = os.Remove(src)
	if err != nil {
		return nil, err
	}
	return filenames, nil
}

// File used by ZipFiles
type File struct {
	Name string
	Body []byte
}
