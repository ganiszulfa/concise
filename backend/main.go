package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/ganiszulfa/concise/backend/config"
	"github.com/ganiszulfa/concise/backend/internal/http"
	log "github.com/sirupsen/logrus"
)

func main() {

	var env string
	flag.StringVar(&env,
		"env",
		"example",
		"env of deployment, will load the respective yml conf file.",
	)
	flag.Parse()

	config.Initialize(env)

	go func() {
		http.Serve()
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		log.Warn(sig)
		done <- true
	}()
	<-done
}
