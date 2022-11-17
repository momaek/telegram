package main

import (
	"log"

	"github.com/fuckproxy/frcdn"
	"github.com/urfave/cli/v2"
)

func PutFile(ctx *cli.Context) (err error) {
	conf, err := GetConfig()
	if err != nil {
		return
	}
	frcdn.Init(conf)
	path := ctx.Args().First()

	u, err := frcdn.PutFile(path)
	log.Println(u)

	return
}
