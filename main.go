package main

import (
	_ "demon/common"
	_ "demon/module/socket"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("demon main")
}
