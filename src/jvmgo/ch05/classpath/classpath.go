package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClassPath Entry
	extClassPath Entry
	userClassPath Entry
}

func Parse(jreOption string, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) ReadClass(classname string) ([]byte, Entry, error) {
	classname = classname + ".class"
	if data, entry, err := self.bootClassPath.readClass(classname); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClassPath.readClass(classname); err == nil {
		return data, entry, err
	}
	return self.userClassPath.readClass(classname)
}

func (self *Classpath) String() string  {
	return self.userClassPath.String()
}
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	//jre/lib/*
	jerLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClassPath = newEntry(jerLibPath)
	//jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClassPath = newEntry(jreExtPath)
}
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return  jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}

	self.userClassPath = newEntry(cpOption)
}
