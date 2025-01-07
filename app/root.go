package app

import (
	"go-smart-contract-event-catch/env"
	"go-smart-contract-event-catch/repository"
)

type App struct {
	env *env.Env

	repository *repository.Repository
}

func NewApp(env *env.Env) {
	a := &App{env: env}

	var err error

	if a.repository, err = repository.NewRepository(env); err != nil {
		panic(err)
	}
}
