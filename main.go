package main

import (
	_ "demon/components/mysql"
	_ "demon/module/orm"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("demon main")
}
