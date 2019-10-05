package main

import "fmt"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1 fake msg")
	} else if cmd.helpFlag || cmd.cpOption == "" || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}
}
func startJvm(cmd *Cmd) {
	fmt.Printf("calsspath:%s class:%s args:%s. jvm start!", cmd.cpOption, cmd.class, cmd.args)
}