package main

import (
	"fmt"
//	"os/exec"
	"os"
)

func main() {
	env := os.Environ()
	procAttr := &os.ProcAttr{Env: env,Files: []*os.File{os.Stdin,os.Stdout,os.Stderr}}
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"},procAttr)
	if err != nil {
		fmt.Println("Error :",err)
		os.Exit(1)
	}
	fmt.Println(pid)
}
