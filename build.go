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

func main() {
	log.Println("start building...")

	flag.Parse()
	if flag.NArg() < 1 {
		log.Println("please use go run build.go to buld.")
		return
	}
	for _, cmd := range flag.Args() {
		switch cmd {
		case "build":
			build()
		}
	}
}

func build() {
	args := []string{"build", "-ldflags", "-w -s"}

	if runtime.GOOS == "windows" {
		BINARY += ".exe"
	}

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
