package classpath

import (
	zip2 "archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if(err != nil) {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error)  {
	zip, err := zip2.OpenReader(self.absPath)
	if err != nil {
		panic(nil)
	}
	defer zip.Close()

	for _, f := range zip.File {
		if f.Name == className {
			file, err :=  f.Open()
			if(err != nil) {
				return nil, nil, err
			}
			defer file.Close()

			data, err := ioutil.ReadAll(file)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
 	}

	return nil, nil, errors.New(fmt.Sprintf("class not found: %s", className))
}

func (self *ZipEntry) String() string {
	return self.absPath
}