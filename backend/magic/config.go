package config

import (
	"fmt"

	"github.com/Liphium/magic/mconfig"
	backend_starter "github.com/Liphium/project-wizard/backend/starter"
)

// This is the function called once you run the project
func Run(ctx *mconfig.Context) {

	ctx.WithEnvironment(&mconfig.Environment{
		"JWT_SECRET": mconfig.ValueStatic("hello"),
		"LISTEN":     mconfig.ValueStatic("127.0.0.1"),
		"PORT":       ctx.ValuePort(3001),
	})

	fmt.Println("Generating config..")
}

func Start() {
	backend_starter.Start()
}
