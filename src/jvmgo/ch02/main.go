package main

import (
	"fmt"
	"jvmgo/ch02/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1 fake msg")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}
}
func startJvm(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("calsspath:%s class:%s args:%s. jvm start!", cp, cmd.class, cmd.args)

	classname := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(classname)
	if err != nil {
		fmt.Printf("\ncould not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("\nclass data: %v\n", classData)
}