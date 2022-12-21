package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/ganiszulfa/concise/backend/config"
	"github.com/ganiszulfa/concise/backend/internal/http"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var env string
	flag.StringVar(&env,
		"env",
		"example",
		"env of deployment, will load the respective yml conf file.",
	)
	flag.Parse()

	config.Initialize(env)

	http.Serve()
}
