package http

import (
	"communication-server/configs"
	"communication-server/internal/apps/apis"
	"communication-server/internal/adapters/server/gingonic"
	"communication-server/internal/adapters/databases/sqlc"
)

func Cmd(conf *configs.Configs) {
	db := sqlc.NewDatabase(conf)
	application := apis.NewApplication(db)
	server := gingonic.NewAdapter(conf, application)
	
	server.Run()
}