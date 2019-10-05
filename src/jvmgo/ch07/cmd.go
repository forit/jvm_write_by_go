package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag bool
	versionFlag bool
	verboseClassFlag bool
	verboseInstFlag bool
	cpOption string
	XjreOption string
	class string
	args []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version message and exit")
	flag.BoolVar(&cmd.verboseClassFlag, "vClass", false, "print class load message")
	flag.BoolVar(&cmd.verboseInstFlag, "vInst", false, "print bytecode instructions message")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath option")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath option")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "jre path option")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
