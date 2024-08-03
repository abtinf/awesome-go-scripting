package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	const name = "app"

	var cmd *exec.Cmd
	switch os.Args[1] {
	case "build":
		cmd = exec.Command("go", "build", "-o", name)
	case "run":
		if len(os.Args) < 3 {
			cmd = exec.Command("./" + name)
		} else {
			cmd = exec.Command("./"+name, os.Args[2:]...)
		}
	case "test":
		cmd = exec.Command("go", "test", "./...")
	case "clean":
		cmd = exec.Command("rm", name)
	default:
		log.Fatalln("invalid command")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
