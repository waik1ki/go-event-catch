package env

import (
	"os"

	"github.com/naoina/toml"
)

type Env struct {
	MongoDB struct {
		Uri      string
		Database string

		Collection struct {
			NFTCollection string
			NFT           string
			Tx            string
		}
	}
}

func NewEnv(path string) *Env {
	c := new(Env)

	if file, err := os.Open(path); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
