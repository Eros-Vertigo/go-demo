package main

import (
	_ "demon/components/redis"
	_ "demon/module/orm"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("demon main")
}
