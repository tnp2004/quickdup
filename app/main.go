package main

import (
	"github.com/tnp2004/quickdup/configs"
	"github.com/tnp2004/quickdup/modules/servers"
)

func main() {
	config := configs.NewConfig()

	server := servers.NewServer(config)
	server.Start()
}
