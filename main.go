package main

import (
	log "github.com/sirupsen/logrus"
	"os/exec"
)

func main() {
	//find, _ := exec.LookPath("php")
	//log.Info(find)
	temp := exec.Command("php", "-v")
	out, _ := temp.Output()
	log.Info(string(out))
	log.Info("demon main")

}
