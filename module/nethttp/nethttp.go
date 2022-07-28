package nethttp

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	log.Info("net/http module")
	http.HandleFunc("/", setup)
	_ = http.ListenAndServe(":8003", nil)
}

func setup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.String())
}
