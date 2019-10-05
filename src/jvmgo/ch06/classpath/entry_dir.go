package classpath

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if !fileInfo.IsDir() {
		panic(fmt.Sprintf("path:%s is not dir", path))
	}

	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error)  {
	filename := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(filename)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
