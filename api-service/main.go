package main

import (
	"api-service/config"
	"github.com/strr0/go-micro-api/cmd"
)

func main() {
	cmd.DefaultResolvers["vpath2"] = config.NewResolver
	cmd.Run()
}
