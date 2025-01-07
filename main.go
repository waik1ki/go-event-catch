package main

import (
	"flag"
	"go-smart-contract-event-catch/app"
	"go-smart-contract-event-catch/env"
)

var envFlag = flag.String("env", "./env.toml", "env not found")

func main() {
	flag.Parse()
	app.NewApp(env.NewEnv(*envFlag))
}
