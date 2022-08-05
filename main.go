package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("demon main")
	add(1, 2)
}

func add(a, b int) int {
	return a + b
}
