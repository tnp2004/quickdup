package main

import (
	"github.com/tnp2004/quickdup/configs"
	"github.com/tnp2004/quickdup/modules/auth/authMiddleware"
	"github.com/tnp2004/quickdup/modules/servers"
	"github.com/tnp2004/quickdup/pkg/databases"
)

func main() {
	config := configs.NewConfig()
	postgresDB := databases.NewPostgresDB(config.Database)
	authMiddleware := authMiddleware.NewAuthMiddleware(postgresDB)
	server := servers.NewServer(config, postgresDB, authMiddleware)
	server.Start()
}
