package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	BINARY     = "Demo"
	SOURCE_DIR = "./DemoSrc"
)

func init() {
	if runtime.GOOS == "windows" {
		BINARY += ".exe"
	}
}

func main() {

	flag.Parse()
	if flag.NArg() != 1 {
		log.Println("please use go run builder.go [ build | clean ] to build.")
		return
	}

	cmd := flag.Arg(0)
	switch cmd {
	case "build":
		clean()
		build()
	case "clean":
		clean()
	default:
		log.Printf("Unknow command: %s.\n", cmd)
	}
}

func clean() {
	log.Println("start cleaning...")
	runCommand("rm", "-f", BINARY)
}

func build() {
	log.Println("start building...")
	args := []string{"build", "-ldflags", "-w -s"}

	args = append(args, "-o", BINARY)
	args = append(args, SOURCE_DIR)
	runCommand("go", args...)
}

func runCommand(cmd string, args ...string) {
	log.Println(cmd, strings.Join(args, " "))
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		log.Println(err)
	}
}
